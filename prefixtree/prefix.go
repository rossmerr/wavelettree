package prefixtree

import "github.com/rossmerr/bitvector"

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type Prefix struct {
	arr     []rune
	prefix  map[rune]*bitvector.BitVector
	version int
}

func NewPrefix() *Prefix {
	return &Prefix{
		arr:    []rune{},
		prefix: map[rune]*bitvector.BitVector{},
	}
}

func NewPrefixFromMap(prefixMap map[rune]*bitvector.BitVector) *Prefix {
	prefix := NewPrefix()

	for k, v := range prefixMap {
		prefix.Append(k, v)
	}

	return prefix
}

func (s *Prefix) Append(r rune, vector *bitvector.BitVector) {
	s.prefix[r] = vector
	s.arr = append(s.arr, r)
	s.version++
}

func (s *Prefix) Element(index int) (rune, *bitvector.BitVector) {
	r := s.arr[index]
	return r, s.prefix[r]
}

func (s *Prefix) Get(r rune) *bitvector.BitVector {
	return s.prefix[r]
}

func (s *Prefix) Has(r rune) bool {
	_, ok := s.prefix[r]
	return ok
}

func (s *Prefix) Length() int {
	return len(s.arr)
}

func (s *Prefix) Enumerate() *PrefixIterator {
	return NewPrefixIteratorWithOffset(s, 0, s.Length())
}

type PrefixIterator struct {
	prefix     *Prefix
	version    int
	indexStart int
	indexEnd   int
}

func NewPrefixIteratorWithOffset(prefix *Prefix, indexStart, indexEnd int) *PrefixIterator {
	if indexStart > prefix.Length() {
		panic("indexStart grater or equal to length")
	}
	if indexStart < 0 {
		panic("indexStart must be non negative number")
	}

	if indexStart > indexEnd {
		panic("indexEnd must be greater then indexStart")
	}

	if prefix.Length()-indexStart < 0 {
		panic("invalid indexStart length")
	}

	if indexEnd > prefix.Length() {
		panic("indexEnd must be greater then prefix length")
	}

	return &PrefixIterator{
		prefix:     prefix,
		indexStart: indexStart,
		indexEnd:   indexEnd,
		version:    prefix.version,
	}
}

func (s *PrefixIterator) Reset() {
	if s.version != s.prefix.version {
		panic("version failed")
	}
	s.indexStart = 0
}

func (s *PrefixIterator) HasNext() bool {
	return s.indexStart < s.indexEnd
}

func (s *PrefixIterator) Next() (rune, *bitvector.BitVector, int) {
	if s.version != s.prefix.version {
		panic("version failed")
	}

	if s.indexStart < s.indexEnd {
		index := s.indexStart
		currentElement, currentVector := s.prefix.Element(index)
		s.indexStart++
		return currentElement, currentVector, index
	}

	s.indexStart = s.prefix.Length()

	return rune('\x10'), nil, s.indexStart
}
