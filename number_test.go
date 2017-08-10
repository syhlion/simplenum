package simplenum

import "testing"

func TestRound(t *testing.T) {
	a := 1.575456
	b := Round(a, 2)
	if b != 1.58 {
		t.Error(a, " round 2 error ", b)
	}
	c := Round(a, 1)
	if c != 1.6 {
		t.Error(a, " round 2 error ", b)
	}
}
func TestCeil(t *testing.T) {
	a := 1.575456
	b := Ceil(a, 2)
	if b != 1.58 {
		t.Error(a, " ceil 2 error ", b)
	}
	c := Ceil(a, 3)
	if c != 1.576 {
		t.Error(a, " ceil 3 error ", c)
	}
}
func TestFloor(t *testing.T) {
	a := 1.575456
	b := Floor(a, 2)
	if b != 1.57 {
		t.Error(a, " floor 2 error ", b)
	}
	c := Floor(a, 3)
	if c != 1.575 {
		t.Error(a, " florr 3 error ", c)
	}
}
