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

func TestStageString(t *testing.T) {
	m := &Map{
		Name:   "ハコフグ倉庫",
		ImgSrc: "https:/splatoon.nintendo.net/assets/img/svg/stage/@2x/75113647a9572411e417d9adaaf243fa09aaf895e416e5af258281d8471a3cae-6a5ba387f31fb7660ffffa1f30da6d1709bfeb696cdf5c1739fdeb142321ec62.jpg",
	}
	s := &Stage{
		Type: "regular",
		Rule: "ナワバリバトル",
	}

	s.Maps = make([]*Map, 2)
	for i, _ := range s.Maps {
		s.Maps[i] = m
	}

	if len(s.String()) < 1 {
		t.Errorf("ERROR: Map String is Empty")
	}

	// if s.String() != id {
	// 	t.Errorf("ERROR: Assets Fail %s != %s", m.String(), id)
	// }

	t.Logf("Map String: %s", s.String())
}

func TestMapString(t *testing.T) {
	m := &Map{
		Name:   "ハコフグ倉庫",
		ImgSrc: "https:/splatoon.nintendo.net/assets/img/svg/stage/@2x/75113647a9572411e417d9adaaf243fa09aaf895e416e5af258281d8471a3cae-6a5ba387f31fb7660ffffa1f30da6d1709bfeb696cdf5c1739fdeb142321ec62.jpg",
	}

	id := "75113647a9572411e417d9adaaf243fa09aaf895e416e5af258281d8471a3cae-6a5ba387f31fb7660ffffa1f30da6d1709bfeb696cdf5c1739fdeb142321ec62"

	if len(m.String()) < 1 {
		t.Errorf("ERROR: Map String is Empty")
	}

	if m.String() != id {
		t.Errorf("ERROR: Assets Fail %s != %s", m.String(), id)
	}

	t.Logf("Map String: %s", m.String())
}
