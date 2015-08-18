package splatoon

import (
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

const (
	TopURL  = "https://splatoon.nintendo.net/"
	AuthURL = "https://splatoon.nintendo.net/users/auth/nintendo"
)

var (
	ErrNotLogin = errors.New("Not login")
)

var test *testing.T

type Client struct {
	hc *http.Client
}

func NewClient(client *http.Client) *Client {
	c := &Client{
		hc: client,
	}

	if client == nil {
		client = &http.Client{}
		c.hc = client
	}

	if client.Jar == nil {
		jar, _ := cookiejar.New(nil)
		c.hc.Jar = jar
	}

	return c
}

func (c *Client) Auth(username, password string) error {
	r, err := c.hc.Get(AuthURL)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	doc, err := newDocument(r)
	if err != nil {
		return err
	}

	ignores := []string{"username", "password", "lang"}
	val := url.Values{}
	doc.Find("input").Each(func(_ int, s *goquery.Selection) {
		n, _ := s.Attr("name")
		for _, ignore := range ignores {
			if ignore == n {
				return
			}
		}
		v, ok := s.Attr("value")
		// logf("splatoon auth: input %s %s", name, val)
		if ok {
			val.Add(n, v)
		}
	})
	val.Add("lang", "ja-JP")
	val.Add("username", username)
	val.Add("password", password)
	action, _ := doc.Find("form").Eq(0).Attr("action")

	// logf("%v", val)
	r, err = c.hc.PostForm(action, val)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	doc, err = newDocument(r)
	if err != nil {
		return err
	}

	errStr := strings.TrimSpace(doc.Find(".hb-error-wrapper").Text())
	if len(errStr) > 0 {
		return errors.New(errStr)
	}

	return nil
}

func (c *Client) ConfirmLogin() error {
	u, _ := url.Parse(TopURL)
	ws := false
	for _, c := range c.hc.Jar.Cookies(u) {
		if c.Name == "_wag_session" {
			ws = true
		}
	}

	if !ws {
		return ErrNotLogin
	}

	r, err := c.hc.Get(TopURL)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.Request.URL.Path == "/sign_in" {
		return ErrNotLogin
	}

	return nil
}

func newDocument(r *http.Response) (*goquery.Document, error) {
	return goquery.NewDocumentFromResponse(r)
}

func logf(format string, args ...interface{}) {
	if test == nil {
		return
	}
	test.Logf(format, args...)
}
