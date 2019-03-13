package gofasta

import (
	"fmt"
	"testing"
)

func TestAlignment_UngappedCoords(t *testing.T) {
	seq1 := "TTT---TTCTTATTG"
	seq2 := "TTT---TTCTTTTTG"
	seq3 := "TTTTTCTTC---TTG"
	a := Alignment{
		NewCharSequence("test", "", seq1),
		NewCharSequence("test", "", seq2),
		NewCharSequence("test", "", seq3),
	}
	expR := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	expC := []int{0, 1, 2, 6, 7, 8, 9, 10, 11, 12, 13, 14, 0, 1, 2, 6, 7, 8, 9, 10, 11, 12, 13, 14, 0, 1, 2, 3, 4, 5, 6, 7, 8, 12, 13, 14}

	r, c := a.UngappedCoords("-")

	for i, expValue := range expR {
		if r[i] != expValue {
			t.Errorf("UngappedCoords(\"-\"): expected row value at (%d) %d, actual %d",
				i, expValue, r[i],
			)
		}
	}
	for i, expValue := range expC {
		if c[i] != expValue {
			t.Errorf("UngappedCoords(\"-\"): expected column value at (%d) %d, actual %d",
				i, expValue, c[i],
			)
		}
	}
}

func TestAlignment_UngappedPositionMatrix(t *testing.T) {
	seq1 := "TTT---TTCTTATTG"
	seq2 := "TTT---TTCTTTTTG"
	seq3 := "TTTTTCTTC---TTG"
	a := Alignment{
		NewCharSequence("test", "", seq1),
		NewCharSequence("test", "", seq2),
		NewCharSequence("test", "", seq3),
	}
	exp := [][]int{
		[]int{0, 1, 2, -1, -1, -1, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		[]int{0, 1, 2, -1, -1, -1, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, -1, -1, -1, 9, 10, 11},
	}

	m := a.UngappedPositionMatrix("-")

	for i, expRow := range exp {
		for j, expValue := range expRow {
			if m[i][j] != expValue {
				t.Errorf("UngappedPositionMatrix(\"-\"): expected value at (%d,%d) %d, actual %d",
					i, j, expValue, m[i][j],
				)
			}
		}
	}
}

func TestAlignment_ToUpper(t *testing.T) {
	names := []string{"test1", "test2", "test3"}
	descs := []string{"test", "", "abc xyz"}
	seqs := []string{"ttt---ttcttattg", "ttt---ttctttttg", "tttttcttc---ttg"}
	a := Alignment{
		NewCharSequence(names[0], descs[0], seqs[0]),
		NewCharSequence(names[1], descs[1], seqs[1]),
		NewCharSequence(names[2], descs[2], seqs[2]),
	}
	exps := []string{"TTT---TTCTTATTG", "TTT---TTCTTTTTG", "TTTTTCTTC---TTG"}
	a.ToUpper()
	for i, seq := range a {
		if exp := exps[i]; exp != seq.Sequence() {
			t.Errorf("ToUpper: expected %#v, actual %#v", exp, seq.Sequence())
		}
	}
}

func TestAlignment_ToLower(t *testing.T) {
	names := []string{"test1", "test2", "test3"}
	descs := []string{"test", "", "abc xyz"}
	seqs := []string{"TTT---TTCTTATTG", "TTT---TTCTTTTTG", "TTTTTCTTC---TTG"}
	a := Alignment{
		NewCharSequence(names[0], descs[0], seqs[0]),
		NewCharSequence(names[1], descs[1], seqs[1]),
		NewCharSequence(names[2], descs[2], seqs[2]),
	}
	exps := []string{"ttt---ttcttattg", "ttt---ttctttttg", "tttttcttc---ttg"}
	a.ToLower()
	for i, seq := range a {
		if exp := exps[i]; exp != seq.Sequence() {
			t.Errorf("ToLower: expected %#v, actual %#v", exp, seq.Sequence())
		}
	}
}

func TestAlignment_ToFasta(t *testing.T) {
	names := []string{"test1", "test2", "test3"}
	descs := []string{"test", "", "abc xyz"}
	seqs := []string{"TTT---TTCTTATTG", "TTT---TTCTTTTTG", "TTTTTCTTC---TTG"}
	a := Alignment{
		NewCharSequence(names[0], descs[0], seqs[0]),
		NewCharSequence(names[1], descs[1], seqs[1]),
		NewCharSequence(names[2], descs[2], seqs[2]),
	}
	exp := fmt.Sprintf(">%s %s\n%s\n", names[0], descs[0], seqs[0]) +
		fmt.Sprintf(">%s\n%s\n", names[1], seqs[1]) +
		fmt.Sprintf(">%s %s\n%s\n", names[2], descs[2], seqs[2])

	actual := a.ToFasta()

	if len(actual) == 0 {
		t.Errorf("ToFasta: expected non-empty byte slice")
	}
	if exp != actual {
		t.Errorf("ToFasta: expected does not match actual: %v != %v", exp, actual)
	}

}

func TestAlignment_Valid_Empty(t *testing.T) {
	a := Alignment{}
	if a.Valid() != true {
		t.Errorf("Valid: expected true but instead got false")
	}
}

func TestAlignment_Valid_True(t *testing.T) {
	names := []string{"test1", "test2", "test3"}
	descs := []string{"test", "", "abc xyz"}
	seqs := []string{"TTT---TTCTTATTG", "TTT---TTCTTTTTG", "TTTTTCTTC---TTG"}
	a := Alignment{
		NewCharSequence(names[0], descs[0], seqs[0]),
		NewCharSequence(names[1], descs[1], seqs[1]),
		NewCharSequence(names[2], descs[2], seqs[2]),
	}
	if a.Valid() != true {
		t.Errorf("Valid: expected true but instead got false")
	}
}

func TestAlignment_Valid_False(t *testing.T) {
	names := []string{"test1", "test2", "test3"}
	descs := []string{"test", "", "abc xyz"}
	seqs := []string{"TTT---TTCTTATTG", "TTT---TTCTTTTTGA", "TTTTTCTTC---TTG"}
	a := Alignment{
		NewCharSequence(names[0], descs[0], seqs[0]),
		NewCharSequence(names[1], descs[1], seqs[1]),
		NewCharSequence(names[2], descs[2], seqs[2]),
	}
	if a.Valid() != false {
		t.Errorf("Valid: expected false but instead got true")
	}
}
