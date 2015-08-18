package splatoon

import (
	"os"
	"testing"
)

var testClient *Client

func TestMain(m *testing.M) {
	testClient = NewClient(nil)
	if err := tryAuth(testClient); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestAuth(t *testing.T) {
	setT(t)
	defer setT(nil)
	c := NewClient(nil)

	t.Log("Nintendo Login")
	err := tryAuth(c)
	if err != nil {
		t.Error(err)
	}
}

func TestConfirmLogin(t *testing.T) {
	setT(t)
	defer setT(nil)
	c := NewClient(nil)

	err := c.ConfirmLogin()
	if err == ErrNotLogin {
		t.Log("PASS Not Login")
	} else if err != nil {
		t.Error(err)
	}

	err = tryAuth(c)
	if err != nil {
		t.Error(err)
	}

	err = c.ConfirmLogin()
	if err != nil {
		t.Error(err)
	}
}

func tryAuth(c *Client) error {
	return c.Auth(os.Getenv("N_USERNAME"), os.Getenv("N_PASSWORD"))
}

func setT(t *testing.T) {
	test = t
}
