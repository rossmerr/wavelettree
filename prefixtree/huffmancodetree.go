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

func NewHuffmanCodeTree(value string) Prefix {
	runeFrequencies, keys := frequencyCount(value)
	huffmanList := rankByRuneCount(runeFrequencies, keys)
	tree := buildHuffmanTree(huffmanList)
	return tree.prefix()
}

func (s *HuffmanCodeTree) isLeaf() bool {
	return s.Value != nil
}

func (s *HuffmanCodeTree) prefix() Prefix {
	prefix := Prefix{}
	left := s.Left
	if left.isLeaf() {
		vector := bitvector.NewBitVectorFromBool([]bool{false})
		prefix[*left.Value] = vector
	} else {
		m := left.prefix()
		for k, v := range m {
			vector := bitvector.NewBitVectorFromVectorPadStart(v, 1)
			vector.Set(0, false)
			prefix[k] = vector
		}
	}

	right := s.Right
	if right.isLeaf() {
		vector := bitvector.NewBitVectorFromBool([]bool{true})
		prefix[*right.Value] = vector
	} else {
		m := right.prefix()
		for k, v := range m {
			vector := bitvector.NewBitVectorFromVectorPadStart(v, 1)
			vector.Set(0, true)
			prefix[k] = vector
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

func frequencyCount(value string) (map[rune]int, []rune) {
	runeFrequencies := make(map[rune]int)
	keys := make([]rune, 0)

	for i := range value {
		r := rune(value[i])
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
