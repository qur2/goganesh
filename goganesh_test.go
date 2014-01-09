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
	it1 := &Item{9, "Banana"}
	it2 := &Item{8, "Lemon"}
	it3 := &Item{4, "Kiwi"}
	it4 := &Item{2, "Peach"}
	it5 := &Item{2, "Apple"}
	items := []*Item{ it1, it2, it3, it4, it5 }
	rankUp(items, 0)
	if items[0] != it1 {
		t.Errorf("Wrong item #0: %s", items[0])
	}
	rankUp(items, 4)
	if items[3] != it5 {
		t.Errorf("Wrong item #3: %s", items[3])
	}
	rankUp(items, 3)
	if items[2] != it5 {
		t.Errorf("Wrong item #2: %s", items[2])
	}
}
