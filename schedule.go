package splatoon

import (
	"path"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Schedule struct {
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
	Stages []*Stage  `json:"stages"`
}

type Stage struct {
	Type string `json:"type"`
	Rule string `json:"rule"`
	Maps []*Map `json:"maps"`
}

type Map struct {
	Name   string `json:"name"`
	ImgSrc string `json:"img_src"`
}

const (
	ScheduleURL    = "https://splatoon.nintendo.net/schedule"
	ScheduleFormat = "1/02 15:04"
)

var (
	timeLocation *time.Location
)

func (c *Client) Schedules() ([]*Schedule, error) {
	if timeLocation == nil {
		loc, _ := time.LoadLocation("Asia/Tokyo")
		timeLocation = loc
	}

	r, err := c.hc.Get(ScheduleURL)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	doc, err := newDocument(r)
	if err != nil {
		return nil, err
	}

	schedules := doc.Find(".stage-schedule")
	ss := make([]*Schedule, schedules.Size())
	schedules.Each(func(i int, d *goquery.Selection) {
		text := d.Text()
		times := strings.Split(text, " ~ ")
		start, _ := time.ParseInLocation(ScheduleFormat, times[0], timeLocation)
		end, _ := time.ParseInLocation(ScheduleFormat, times[1], timeLocation)
		s := &Schedule{
			Start: start,
			End:   end,
		}
		ss[i] = s
	})

	stages := doc.Find(".stage-list")
	count := 0
	for i, _ := range ss {
		ss[i].Stages = make([]*Stage, 2)
		for _, j := range []int{0, 1} {
			stage := stages.Eq(count)
			s := &Stage{
				Type: "regular",
				Rule: "ナワバリバトル",
			}
			rule := stage.Find(".match-rule")
			if rule.Size() > 0 {
				s.Type = "earnest"
				s.Rule = rule.Find(".rule-description").Eq(0).Text()
			}
			maps := stage.Find(".map-box")
			s.Maps = make([]*Map, maps.Size())
			maps.Each(func(i int, d *goquery.Selection) {
				m := &Map{
					Name:   d.Find(".map-name").Eq(0).Text(),
					ImgSrc: path.Join(TopURL, d.Find(".map-image").Eq(0).AttrOr("data-retina-image", "")),
				}
				s.Maps[i] = m
			})
			ss[i].Stages[j] = s
			count += 1
		}
	}

	return ss, nil
}