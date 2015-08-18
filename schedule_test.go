package splatoon

import (
	"testing"
)

func TestSchedules(t *testing.T) {
	setT(t)
	defer setT(nil)
	ss, err := testClient.Schedules()
	if err != nil {
		t.Error(err)
	}

	for _, s := range ss {
		t.Logf("Schedules %v", s)
		for _, s := range s.Stages {
			t.Logf("Stage %v", s)
			for _, m := range s.Maps {
				t.Logf("Map %v", m)
			}
		}
	}
}
