package gofasta

import (
	"fmt"
	"os"
)

// Alignment is a slice of Sequence pointers.
type Alignment []Sequence

// UngappedCoords returns the row and column positions in the sequence alignment
// where the character does not match the gap character.
func (a Alignment) UngappedCoords(gapChar string) (rowCoords, colCoords []int) {
	var currColCoords []int
	for i, s := range a {
		currColCoords = s.UngappedCoords(gapChar)
		for c := 0; c < len(currColCoords); c++ {
			rowCoords = append(rowCoords, i)
		}
		colCoords = append(colCoords, currColCoords...)
	}
	return
}

// UngappedPositionMatrix returns a matrix that counts only over characters
// that does not match the gap character for each sequence in the alignment.
// If a character in a sequence matches the gap character, -1 is inserted
// instead of the ungapped count.
func (a Alignment) UngappedPositionMatrix(gapChar string) (m [][]int) {
	for _, s := range a {
		m = append(m, s.UngappedPositionSlice(gapChar))
	}
	return
}

// ToUpper changes the case of all sequences to all uppercase letters.
func (a Alignment) ToUpper() {
	for _, s := range a {
		s.ToUpper()
	}
}

// ToLower changes the case of of all sequences to all lowercase letters.
func (a Alignment) ToLower() {
	for _, s := range a {
		s.ToLower()
	}
}

// Valid tells whether the sequences in the alignment have the same length.
func (a Alignment) Valid() bool {
	if len(a) == 0 {
		return true
	}
	length := len(a[0].Sequence())
	for _, s := range a {
		if length != len(s.Sequence()) {
			return false
		}
	}
	return true
}

// ToFasta saves the sequence alignment to a FASTA file.
func (a Alignment) ToFasta(path string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var b []byte
	a.Write(b)
	f.Write(b)
	f.Sync()
}

func (a Alignment) Write(p []byte) (n int, err error) {
	// Append each Sequence in Alignment
	var blen int
	for _, s := range a {
		var bstr []byte
		if len(s.Description()) > 0 {
			bstr = []byte(fmt.Sprintf(">%s %s\n", s.ID(), s.Description()))
		} else {
			bstr = []byte(fmt.Sprintf(">%s\n", s.ID()))
		}
		bstr = []byte(fmt.Sprintf("%s\n", s.Sequence()))
		blen += len(bstr)
		p = append(p, bstr...)
	}
	return blen, nil
}
