package prefixtree

import "github.com/rossmerr/bitvector"

type Prefix map[rune]*bitvector.BitVector

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (s Prefix) Append(r rune) *bitvector.BitVector {
	vector := bitvector.NewBitVector(0)
	s[r] = vector
	return vector
}
