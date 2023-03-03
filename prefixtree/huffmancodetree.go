package prefixtree

import (
	"sort"

	"github.com/rossmerr/bitvector"
)

type HuffmanCodeTree struct {
	Left      *HuffmanCodeTree
	Right     *HuffmanCodeTree
	Value     *rune
	Frequency int
}

func NewHuffmanCodeTree(value []rune) *Prefix {
	runeFrequencies, keys := frequencyCount(value)
	huffmanList := rankByRuneCount(runeFrequencies, keys)
	tree := buildHuffmanTree(huffmanList)
	return tree.prefix()
}

func NewHuffmanCodeTreeFromFrequencies(runeFrequencies map[rune]int, keys []rune) *Prefix {
	huffmanList := rankByRuneCount(runeFrequencies, keys)
	tree := buildHuffmanTree(huffmanList)
	return tree.prefix()
}

func (s *HuffmanCodeTree) isLeaf() bool {
	return s.Value != nil
}

func (s *HuffmanCodeTree) prefix() *Prefix {
	prefix := NewPrefix()
	left := s.Left
	if left.isLeaf() {
		vector := bitvector.NewBitVectorFromBool([]bool{false})
		prefix.Append(*left.Value, vector)
	} else {
		m := left.prefix()
		iterator := m.Enumerate()
		for iterator.HasNext() {
			k, v, _ := iterator.Next()
			vector := bitvector.NewBitVectorFromVectorPadStart(v, 1)
			vector.Set(0, false)
			prefix.Append(k, vector)
		}
	}

	right := s.Right
	if right.isLeaf() {
		vector := bitvector.NewBitVectorFromBool([]bool{true})
		prefix.Append(*right.Value, vector)
	} else {
		m := right.prefix()
		iterator := m.Enumerate()
		for iterator.HasNext() {
			k, v, _ := iterator.Next()
			vector := bitvector.NewBitVectorFromVectorPadStart(v, 1)
			vector.Set(0, true)
			prefix.Append(k, vector)
		}
	}

	return prefix
}

func buildHuffmanTree(list huffmanList) *HuffmanCodeTree {
	for {
		first := list[0]
		list = list[1:]

		if len(list) == 0 {
			return first
		}

		second := list[0]
		list = list[1:]

		sum := first.Frequency + second.Frequency

		t := &HuffmanCodeTree{
			Frequency: sum,
			Left:      first,
			Right:     second,
		}

		if len(list) == 0 {
			return t
		}

		for _, pair := range list {
			if pair.Frequency >= sum {
				list = append([]*HuffmanCodeTree{t}, list...)
			} else {
				list = append(list, t)

			}
			break
		}

	}
}

func frequencyCount(value []rune) (map[rune]int, []rune) {
	runeFrequencies := make(map[rune]int)
	keys := make([]rune, 0)

	for _, r := range value {
		if _, ok := runeFrequencies[r]; ok {
			runeFrequencies[r] = runeFrequencies[r] + 1
		} else {
			runeFrequencies[r] = 1
			keys = append(keys, r)

		}
	}

	return runeFrequencies, keys
}

func rankByRuneCount(runeFrequencies map[rune]int, keys []rune) huffmanList {
	list := make(huffmanList, len(runeFrequencies))

	for i, r := range keys {
		v := r
		list[i] = &HuffmanCodeTree{Value: &v, Frequency: runeFrequencies[r]}
	}
	sort.Sort(list)
	return list
}

type huffmanList []*HuffmanCodeTree

func (p huffmanList) Len() int           { return len(p) }
func (p huffmanList) Less(i, j int) bool { return p[i].Frequency < p[j].Frequency }
func (p huffmanList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
