package gofasta

import "testing"

func TestNewCharSequence(t *testing.T) {
	name := "a"
	desc := "test"
	seq := "ATGGCGTAG"
	exp := CharSequence{name, desc, seq}
	actual := *NewCharSequence(name, desc, seq)

	if exp.name != actual.name && exp.description != actual.description && exp.sequence != actual.sequence {
		t.Errorf("NewCharSequence: expected %#v, actual %#v", exp, actual)
	}
}

func TestCharSequence_Properties(t *testing.T) {
	id := "a"
	desc := "test"
	seq := "ATGGCGTAG"
	actual := CharSequence{id, desc, seq}

	if id != actual.ID() {
		t.Errorf("expected %#v, actual %#v", id, actual.ID())
	}
	if desc != actual.Description() {
		t.Errorf("expected %#v, actual %#v", desc, actual.Description())
	}
	if seq != actual.Sequence() {
		t.Errorf("Sequence: expected %#v, actual %#v", seq, actual.Sequence())
	}
}

func TestCharSequence_SetID(t *testing.T) {
	seq := CharSequence{"a", "test", "ATGGCGTAG"}
	exp := "x"
	seq.SetID(exp)
	if exp != seq.name {
		t.Errorf("Setexpected %#v, actual %#v", exp, seq.name)
	}
}

func TestCharSequence_SetDescription(t *testing.T) {
	seq := CharSequence{"a", "test", "ATGGCGTAG"}
	exp := "test again"
	seq.SetDescription(exp)
	if exp != seq.description {
		t.Errorf("Setexpected %#v, actual %#v", exp, seq.description)
	}
}

func TestCharSequence_SetSequence(t *testing.T) {
	seq := CharSequence{"a", "test", "ATGGCGTAG"}
	exp := "CCCCCCCCC"
	seq.SetSequence(exp)
	if exp != seq.sequence {
		t.Errorf("SetSequence: expected %#v, actual %#v", exp, seq.sequence)
	}
}

func TestCharSequence_ToUpper(t *testing.T) {
	seq := CharSequence{"a", "test", "atggcgtag"}
	exp := "ATGGCGTAG"
	seq.ToUpper()
	if exp != seq.sequence {
		t.Errorf("ToUpper: expected %#v, actual %#v", exp, seq.sequence)
	}
}

func TestCharSequence_ToLower(t *testing.T) {
	seq := CharSequence{"a", "test", "ATGGCGTAG"}
	exp := "atggcgtag"
	seq.ToLower()
	if exp != seq.sequence {
		t.Errorf("ToLower: expected %#v, actual %#v", exp, seq.sequence)
	}
}

func TestCharSequence_Char(t *testing.T) {
	seq := CharSequence{"a", "test", "GTGGCGTAG"}
	exp := "A"
	actual := seq.Char(7)
	if exp != actual {
		t.Errorf("Char: expected %#v, actual %#v", exp, actual)
	}
}

func TestCharSequence_UngappedCoords(t *testing.T) {
	seq := "TTT---TTCTTATTG"
	s := CharSequence{"test", "", seq}
	exp := []int{0, 1, 2, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	res := s.UngappedCoords("-")

	for i, expValue := range exp {
		if expValue != res[i] {
			t.Errorf("UngappedCoords(\"-\"): expected (%d) %d, actual %d",
				i, expValue, res[i],
			)
		}
	}
}

func TestCharSequence_UngappedPositionSlice(t *testing.T) {
	seq := "TTT---TTCTTATTG"
	s := CharSequence{"test", "", seq}
	exp := []int{0, 1, 2, -1, -1, -1, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	res := s.UngappedPositionSlice("-")

	for i, expValue := range exp {
		if expValue != res[i] {
			t.Errorf("UngappedCoords(\"-\"): expected (%d) %d, actual %d",
				i, expValue, res[i],
			)
		}
	}
}
