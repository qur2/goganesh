package main

import (
	"testing"
)

func TestItemify(t *testing.T) {
	s1 := "23:haxxor:prog"
	it1, _ := Itemify(s1)
	if it1.Score != 23 {
		t.Errorf("Wrong item score: %d", it1.Score)
	}
	if it1.Title != "haxxor:prog" {
		t.Errorf("Wrong item title: %s", it1.Title)
	}

	s2 := "haxxor prog"
	it2, err := Itemify(s2)
	if it2.Title != "" {
		t.Errorf("Wrong item title: %s", it2.Title)
	}
	if err == nil {
		t.Errorf("Invalid input with no error")
	}

	s3 := ":haxxor prog"
	it3, err := Itemify(s3)
	if it3.Score != 0 {
		t.Errorf("Wrong item score: %d", it3.Score)
	}
	if it3.Title != "haxxor prog" {
		t.Errorf("Wrong item title: %s", it3.Title)
	}
	if err == nil {
		t.Errorf("Invalid input with no error")
	}
}

func TestSearchByTitle(t *testing.T) {
	items := []*Item{
		&Item{9, "Banana"},
		&Item{8, "Lemon"},
		&Item{5, "Kiwi"},
		&Item{2, "Peach"},
		&Item{2, "Apple"},
	}
	found := searchByTitle(items, "Apple")
	if found != 4 {
		t.Errorf("Found something at the wrong index: %d", found)
	}
	if items[found].Title != "Apple" {
		t.Errorf("Found item has wrong title: %s", items[found].Title)
	}
}

func TestRankUp(t *testing.T) {
	scores := []float64{9, 5, 3, 1, .1}
	it1 := &Item{9, "Banana"}
	it2 := &Item{5, "Lemon"}
	it3 := &Item{3, "Kiwi"}
	it4 := &Item{1, "Peach"}
	it5 := &Item{0.1, "Apple"}
	items := []*Item{it1, it2, it3, it4, it5}
	for i := 0; i < 5; i++ {
		rankUp(items, i)
		if items[i].Score <= scores[i] {
			t.Errorf("Score of item #%d decreased: %.6f", i, items[i].Score)
		}
	}
}

func TestDecay(t *testing.T) {
	scores := []float64{9, 5, 3, 1, .1}
	it1 := &Item{9, "Banana"}
	it2 := &Item{5, "Lemon"}
	it3 := &Item{3, "Kiwi"}
	it4 := &Item{1, "Peach"}
	it5 := &Item{0.1, "Apple"}
	items := []*Item{it1, it2, it3, it4, it5}
	decay(items)
	for i := 0; i < 5; i++ {
		if items[i].Score >= scores[i] {
			t.Errorf("Score of item #%d increased: %.6f", i, items[i].Score)
		}
	}
}
