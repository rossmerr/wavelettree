package wavelettree

import (
	"sort"

	"github.com/rossmerr/bitvector"
)

type HuffmanCode struct {
	Left      *HuffmanCode
	Right     *HuffmanCode
	Value     *rune
	Frequency int
}

func NewHuffmanCode(value string) *HuffmanCode {
	runeFrequencies := frequencyCount(value)
	huffmanList := rankByRuneCount(runeFrequencies)
	return buildHuffmanTree(huffmanList)
}

func (s *HuffmanCode) isLeaf() bool {
	return s.Value != nil
}

func (s *HuffmanCode) Prefix() map[rune]*bitvector.BitVector {
	prefix := map[rune]*bitvector.BitVector{}
	left := s.Left
	if left.isLeaf() {
		vector := bitvector.NewBitVectorFromBool([]bool{false})
		prefix[rune(*left.Value)] = vector
	} else {
		m := left.Prefix()
		for k, v := range m {
			vector := bitvector.NewBitVectorFromVectorPadStart(v, 1)
			vector.Set(0, false)
			prefix[k] = vector
		}
	}

	right := s.Right
	if right.isLeaf() {
		vector := bitvector.NewBitVectorFromBool([]bool{true})

		prefix[rune(*right.Value)] = vector
	} else {
		m := right.Prefix()
		for k, v := range m {
			vector := bitvector.NewBitVectorFromVectorPadStart(v, 1)
			vector.Set(0, true)
			prefix[k] = vector
		}
	}

	return prefix
}

func buildHuffmanTree(list huffmanList) *HuffmanCode {
	for {
		first := list[0]
		list = list[1:]

		if len(list) == 0 {
			return first
		}

		second := list[0]
		list = list[1:]

		sum := first.Frequency + second.Frequency

		t := &HuffmanCode{
			Frequency: sum,
			Left:      first,
			Right:     second,
		}

		if len(list) == 0 {
			return t
		}

		for _, pair := range list {
			if pair.Frequency >= sum {
				list = append([]*HuffmanCode{t}, list...)
			} else {
				list = append(list, t)

			}
			break
		}

	}
}

func frequencyCount(value string) map[rune]int {
	runeFrequencies := make(map[rune]int)
	for i := range value {
		r := rune(value[i])
		if _, ok := runeFrequencies[r]; ok {
			runeFrequencies[r] = runeFrequencies[r] + 1
		} else {
			runeFrequencies[r] = 1
		}
	}

	return runeFrequencies
}

func rankByRuneCount(runeFrequencies map[rune]int) huffmanList {
	list := make(huffmanList, len(runeFrequencies))
	i := 0

	keys := make([]string, 0)
	for k, _ := range runeFrequencies {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)

	for _, k := range keys {
		r := []rune(k)[0]

		list[i] = &HuffmanCode{Value: &r, Frequency: runeFrequencies[r]}
		i++
	}
	sort.Sort(list)
	return list
}

type huffmanList []*HuffmanCode

func (p huffmanList) Len() int           { return len(p) }
func (p huffmanList) Less(i, j int) bool { return p[i].Frequency < p[j].Frequency }
func (p huffmanList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
