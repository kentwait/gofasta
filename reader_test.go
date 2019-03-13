package gofasta

import (
	"strings"
	"testing"
)

func TestFastaToAlignment_Char(t *testing.T) {
	r := strings.NewReader("# comment1\n\n" +
		"; comment2\n" +
		">a test\n" +
		"TTT---TTCTTATTG\n" +
		">b\n" +
		"TAT---TTCTTTTTG\n" +
		">c test1 test2\n" +
		"TTTTTCTTC---TTG")
	exp := []Sequence{
		NewCharSequence("a", "test", "TTT---TTCTTATTG"),
		NewCharSequence("b", "", "TAT---TTCTTTTTG"),
		NewCharSequence("c", "test1 test2", "TTTTTCTTC---TTG"),
	}
	actual := FastaToAlignment(r, false)
	for i, s := range actual {
		if s.ID() != exp[i].ID() {
			t.Errorf("FastaToAlignment: name mismatch in sequence %v: %s != %s\n%#v", i, exp[i].ID(), s.ID(), s)
		}
		if s.Description() != exp[i].Description() {
			t.Errorf("FastaToAlignment: description mismatch in sequence %v: %s != %s\n%#v", i, exp[i].Description(), s.Description(), s)
		}
		if s.Sequence() != exp[i].Sequence() {
			t.Errorf("FastaToAlignment: sequence mismatch in sequence %v: %s != %s\n%#v", i, exp[i].Sequence(), s.Sequence(), s)
		}
	}
}

func TestFastaToAlignment_Codon(t *testing.T) {
	r := strings.NewReader("# comment1\n\n" +
		"; comment2\n" +
		">a test\n" +
		"TTT---TTCTTATTG\n" +
		">b\n" +
		"TAT---TTCTTTTTG\n" +
		">c test1 test2\n" +
		"TTTTTCTTC---TTG")
	exp := []Sequence{
		NewCharSequence("a", "test", "TTT---TTCTTATTG"),
		NewCharSequence("b", "", "TAT---TTCTTTTTG"),
		NewCharSequence("c", "test1 test2", "TTTTTCTTC---TTG"),
	}
	actual := FastaToAlignment(r, true)
	for i, s := range actual {
		if s.ID() != exp[i].ID() {
			t.Errorf("FastaToAlignment: name mismatch in sequence %v: %s != %s\n%#v", i, exp[i].ID(), s.ID(), s)
		}
		if s.Description() != exp[i].Description() {
			t.Errorf("FastaToAlignment: description mismatch in sequence %v: %s != %s\n%#v", i, exp[i].Description(), s.Description(), s)
		}
		if s.Sequence() != exp[i].Sequence() {
			t.Errorf("FastaToAlignment: sequence mismatch in sequence %v: %s != %s\n%#v", i, exp[i].Sequence(), s.Sequence(), s)
		}
	}
}
