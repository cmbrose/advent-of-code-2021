package main

import (
	"testing"
)

func TestVolume(t *testing.T) {
	s := Sector{0, 0, 0, 0, 0, 0, false}
	if s.Volume() != 1 {
		t.Fatalf("Expected volume to be 0 but was %d", s.Volume())
	}

	s = Sector{0, 1, 0, 0, 0, 0, false}
	if s.Volume() != 2 {
		t.Fatalf("Expected volume to be 2 but was %d", s.Volume())
	}

	s = Sector{0, 1, 0, 1, 0, 1, false}
	if s.Volume() != 8 {
		t.Fatalf("Expected volume to be 8 but was %d", s.Volume())
	}

	s = Sector{-1, 0, 0, 0, 0, 0, false}
	if s.Volume() != 2 {
		t.Fatalf("Expected volume to be 2 but was %d", s.Volume())
	}

	s = Sector{-1, 0, -1, 0, -1, 0, false}
	if s.Volume() != 8 {
		t.Fatalf("Expected volume to be 8 but was %d", s.Volume())
	}
}

func TestIntersect(t *testing.T) {
	base := Sector{-1, -1, -1, 1, 1, 1, false}

	test := Sector{0, 0, 0, 0, 0, 0, false}
	intersect, ok := base.Intersect(test)
	if !ok || !intersect.Equals(Sector{0, 0, 0, 0, 0, 0, false}) {
		t.Fatalf("Intersection is wrong. Ok=%v, Expected %v, got %v", ok, Sector{0, 0, 0, 0, 0, 0, false}, intersect)
	}

	test = Sector{0, 0, 0, 1, 1, 1, false}
	intersect, ok = base.Intersect(test)
	if !ok || !intersect.Equals(Sector{0, 0, 0, 1, 1, 1, false}) {
		t.Fatalf("Intersection is wrong. Ok=%v, Expected %v, got %v", ok, Sector{0, 0, 0, 0, 0, 0, false}, intersect)
	}

	test = Sector{-2, 0, 0, 2, 0, 0, false}
	intersect, ok = base.Intersect(test)
	if !ok || !intersect.Equals(Sector{-1, 0, 0, 1, 0, 0, false}) {
		t.Fatalf("Intersection is wrong. Ok=%v, Expected %v, got %v", ok, Sector{0, 0, 0, 0, 0, 0, false}, intersect)
	}

	test = Sector{0, -2, 0, 0, 2, 0, false}
	intersect, ok = base.Intersect(test)
	if !ok || !intersect.Equals(Sector{0, -1, 0, 0, 1, 0, false}) {
		t.Fatalf("Intersection is wrong. Ok=%v, Expected %v, got %v", ok, Sector{0, 0, 0, 0, 0, 0, false}, intersect)
	}

	test = Sector{0, 0, -2, 0, 0, 2, false}
	intersect, ok = base.Intersect(test)
	if !ok || !intersect.Equals(Sector{0, 0, -1, 0, 0, 1, false}) {
		t.Fatalf("Intersection is wrong. Ok=%v, Expected %v, got %v", ok, Sector{0, 0, 0, 0, 0, 0, false}, intersect)
	}

	intersect, ok = base.Intersect(base)
	if !ok || !intersect.Equals(base) {
		t.Fatalf("Intersection is wrong. Ok=%v, Expected %v, got %v", ok, Sector{0, 0, 0, 0, 0, 0, false}, intersect)
	}

	test = Sector{-2, -2, -2, 2, 2, 2, false}
	intersect, ok = base.Intersect(test)
	if !ok || !intersect.Equals(base) {
		t.Fatalf("Intersection is wrong. Ok=%v, Expected %v, got %v", ok, Sector{0, 0, 0, 0, 0, 0, false}, intersect)
	}

	test = Sector{0, -2, -2, 0, 2, 2, false}
	intersect, ok = base.Intersect(test)
	if !ok || !intersect.Equals(Sector{0, -1, -1, 0, 1, 1, false}) {
		t.Fatalf("Intersection is wrong. Ok=%v, Expected %v, got %v", ok, Sector{0, 0, 0, 0, 0, 0, false}, intersect)
	}

	test = Sector{0, 0, -2, 0, 1, 2, false}
	intersect, ok = base.Intersect(test)
	if !ok || !intersect.Equals(Sector{0, 0, -1, 0, 1, 1, false}) {
		t.Fatalf("Intersection is wrong. Ok=%v, Expected %v, got %v", ok, Sector{0, 0, 0, 0, 0, 0, false}, intersect)
	}

	test = Sector{0, 0, 0, 0, 0, 0, false}
	intersect, ok = test.Intersect(base)
	if !ok || !intersect.Equals(test) {
		t.Fatalf("Intersection is wrong. Ok=%v, Expected %v, got %v", ok, Sector{0, 0, 0, 0, 0, 0, false}, intersect)
	}

	test = Sector{-3, -3, -3, 3, 3, 3, false}
	intersect, ok = base.Intersect(test)
	if !ok || !intersect.Equals(base) {
		t.Fatalf("Intersection is wrong. Ok=%v, Expected %v, got %v", ok, Sector{0, 0, 0, 0, 0, 0, false}, intersect)
	}

	test = Sector{2, 1, 1, 2, 1, 1, false}
	intersect, ok = base.Intersect(test)
	if ok {
		t.Fatalf("Intersection is wrong. Ok=%v, Expected %v, got %v", ok, Sector{0, 0, 0, 0, 0, 0, false}, intersect)
	}

	test = Sector{-2, 1, 1, -2, 1, 1, false}
	intersect, ok = base.Intersect(test)
	if ok {
		t.Fatalf("Intersection is wrong. Ok=%v, Expected %v, got %v", ok, Sector{0, 0, 0, 0, 0, 0, false}, intersect)
	}

	test = Sector{1, 2, 1, 1, 2, 1, false}
	intersect, ok = base.Intersect(test)
	if ok {
		t.Fatalf("Intersection is wrong. Ok=%v, Expected %v, got %v", ok, Sector{0, 0, 0, 0, 0, 0, false}, intersect)
	}

	test = Sector{1, -2, 1, 1, -2, 1, false}
	intersect, ok = base.Intersect(test)
	if ok {
		t.Fatalf("Intersection is wrong. Ok=%v, Expected %v, got %v", ok, Sector{0, 0, 0, 0, 0, 0, false}, intersect)
	}

	test = Sector{1, 1, 2, 1, 1, 2, false}
	intersect, ok = base.Intersect(test)
	if ok {
		t.Fatalf("Intersection is wrong. Ok=%v, Expected %v, got %v", ok, Sector{0, 0, 0, 0, 0, 0, false}, intersect)
	}

	test = Sector{1, 1, -2, 1, 1, -2, false}
	intersect, ok = base.Intersect(test)
	if ok {
		t.Fatalf("Intersection is wrong. Ok=%v, Expected %v, got %v", ok, Sector{0, 0, 0, 0, 0, 0, false}, intersect)
	}
}

func TestSplit(t *testing.T) {
	base := Sector{-1, -1, -1, 1, 1, 1, false}

	split := base.Split(base)
	if split != nil {
		t.Fatalf("Spliting with self should result in nil")
	}

	test := Sector{-1, -1, -1, 1, -1, -1, false}
	split = base.Split(test)
	if len(split) != 2 {
		t.Fatalf("Split is wrong. Expected 2 items, got %v", split)
	}
	if !split[0].Equals(Sector{-1, 0, -1, 1, 1, 1, false}) {
		t.Fatalf("Split[0] is wrong. Expected %v, got %v", Sector{-1, 0, -1, 1, 1, 1, false}, split[0])
	}
	if !split[1].Equals(Sector{-1, -1, 0, 1, -1, 1, false}) {
		t.Fatalf("Split[0] is wrong. Expected %v, got %v", Sector{-1, -1, 0, 1, -1, 1, false}, split[1])
	}

	sector := Sector{-1, 0, 0, 1, 0, 0, false}
	test = Sector{-1, 0, 0, -1, 0, 0, false}
	split = sector.Split(test)
	if len(split) != 1 {
		t.Fatalf("Split is wrong. Expected 1 items, got %v", split)
	}

	sector = Sector{-1, 0, 0, 1, 0, 0, false}
	test = Sector{-1, 0, 0, 0, 0, 0, false}
	split = sector.Split(test)
	if len(split) != 1 {
		t.Fatalf("Split is wrong. Expected 1 items, got %v", split)
	}

	sector = Sector{-1, 0, 0, 1, 0, 0, false}
	test = Sector{0, 0, 0, 1, 0, 0, false}
	split = sector.Split(test)
	if len(split) != 1 {
		t.Fatalf("Split is wrong. Expected 1 items, got %v", split)
	}

	sector = Sector{-1, -1, 0, 1, 1, 0, false}
	test = Sector{-1, -1, 0, 1, -1, 0, false}
	split = sector.Split(test)
	if len(split) != 1 {
		t.Fatalf("Split is wrong. Expected 1 items, got %v", split)
	}

	sector = Sector{-1, -1, 0, 1, 1, 0, false}
	test = Sector{-1, -1, 0, 0, 0, 0, false}
	split = sector.Split(test)
	if len(split) != 2 {
		t.Fatalf("Split is wrong. Expected 2 items, got %v", split)
	}

	test = Sector{-1, 0, -1, 1, 0, -1, false}
	split = base.Split(test)
	if len(split) != 3 {
		t.Fatalf("Split is wrong. Expected 3 items, got %v", split)
	}

	test = Sector{-1, 0, 0, 1, 0, 0, false}
	split = base.Split(test)
	if len(split) != 4 {
		t.Fatalf("Split is wrong. Expected 4 items, got %d %v", len(split), split)
	}

	test = Sector{0, 0, 0, 0, 0, 0, false}
	split = base.Split(test)
	if len(split) != 6 {
		t.Fatalf("Split is wrong. Expected 6 items, got %d %v", len(split), split)
	}

	test = Sector{-1, 0, 0, 0, 0, 0, false}
	split = base.Split(test)
	if len(split) != 5 {
		t.Fatalf("Split is wrong. Expected 5 items, got %d %v", len(split), split)
	}

	test = Sector{0, 0, 0, 1, 0, 0, false}
	split = base.Split(test)
	if len(split) != 5 {
		t.Fatalf("Split is wrong. Expected 5 items, got %d %v", len(split), split)
	}

	test = Sector{1, 1, 1, 1, 1, 1, false}
	split = base.Split(test)
	if len(split) != 3 {
		t.Fatalf("Split is wrong. Expected 3 items, got %d %v", len(split), split)
	}
}
