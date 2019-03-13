package gofasta

// Sequence is an interface for single character sequences stored as a string
// and multi-character sequences stored as a slice.
type Sequence interface {
	SequenceMeta
	SequenceGetter
	SequenceSetter
}

// SequenceMeta defines methods to retrieve sequence metadata.
type SequenceMeta interface {
	ID() string
	Description() string
}

// SequenceGetter contains methods to retrieve information about sequence data.
type SequenceGetter interface {
	Sequence() string
	Char(int) string
	UngappedCoords(string) []int
	UngappedPositionSlice(string) []int
}

// SequenceSetter contains methods to set/modify sequence data.
type SequenceSetter interface {
	SetID(string)
	SetDescription(string)
	SetSequence(string)
	ToUpper()
	ToLower()
}
