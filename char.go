package gofasta

import (
	"sort"
	"strings"
)

// CharSequence struct for storing single-character biological sequences such
// as nucleotides and single-letter amino acids. However, any sequence that
// whose element can be represented as a single string character can be stored
// in CharSequence.
type CharSequence struct {
	name        string
	description string
	sequence    string
}

// NewCharSequence contructs a new CharSequence.
func NewCharSequence(name, description, sequence string) *CharSequence {
	return &CharSequence{name, description, sequence}
}

// ID returns the name of CharSequence.
func (s *CharSequence) ID() string {
	return s.name
}

// Description returns a description of CharSequence.
func (s *CharSequence) Description() string {
	return s.description
}

// Sequence returns the sequence of CharSequence.
func (s *CharSequence) Sequence() string {
	return s.sequence
}

// Char returns a single character by index from the sequence of CharSequence.
func (s *CharSequence) Char(i int) string {
	return string([]rune(s.sequence)[i])
}

// SetID sets the name of CharSequence.
func (s *CharSequence) SetID(id string) {
	s.name = id
}

// SetDescription sets a description for CharSequence.
func (s *CharSequence) SetDescription(title string) {
	s.description = title
}

// SetSequence sets the sequence of CharSequence.
func (s *CharSequence) SetSequence(seq string) {
	s.sequence = seq
}

// UngappedCoords returns the positions in the sequence where the character
// does not match the gap character.
func (s *CharSequence) UngappedCoords(gapChar string) (colCoords []int) {
	set := make(map[int]struct{})
	// Assumes gapChar contains only a "single character"
	// Convert single character string to rune slice, taking the first item
	gapRune := []rune(gapChar)[0]
	// Range over rune slice, j counts by Unicode code points, s is the rune representation of the character
	for j, s := range []rune(s.sequence) {
		// If sequence rune is not a gap character rune, add to rune position to set, 0-indexed
		if s != gapRune {
			set[j] = struct{}{} // Uses empty anonymous struct
		}
	}
	// Range over set of positions
	// Since this is a map, order is scrambled
	for key := range set {
		colCoords = append(colCoords, key)
	}
	sort.Ints(colCoords)
	return
}

// UngappedPositionSlice returns a slice that counts only over characters
// that does not match the gap character in the sequence.
// If a character matches the gap character, -1 is inserted instead of the
// ungapped count.
func (s *CharSequence) UngappedPositionSlice(gapChar string) (arr []int) {
	// Assumes gapChar contains only a "single character"
	// Convert single character string to rune slice, taking the first item
	gapRune := []rune(gapChar)[0]
	cnt := 0
	for _, s := range []rune(s.sequence) {
		// If sequence rune is not a gap character rune, append current count value to array and increment
		if s != gapRune {
			arr = append(arr, cnt)
			cnt++
			// If it is equal to the gap character rune, then append a -1.
			// Do not increment.
		} else {
			arr = append(arr, -1)
		}
	}
	return
}

// ToUpper changes the case of the sequence to all uppercase letters.
func (s *CharSequence) ToUpper() {
	s.sequence = strings.ToUpper(s.sequence)
}

// ToLower changes the case of the sequence to all lowercase letters.
func (s *CharSequence) ToLower() {
	s.sequence = strings.ToLower(s.sequence)
}
