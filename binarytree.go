package wavelettree

import (
	"sort"

	"github.com/rossmerr/bitvector"
)

type BinaryTree struct {
	Left  *BinaryTree
	Right *BinaryTree
	Value *rune
}

func NewBinaryTree(value string) *BinaryTree {
	runeFrequencies := binaryCount(value)
	binaryList := rankByBinaryCount(runeFrequencies)
	return buildBinaryTree(value, binaryList)
}

func binaryCount(value string) map[rune]int {
	runeFrequencies := make(map[rune]int)

	for _, entry := range value {
		if _, ok := runeFrequencies[entry]; !ok {
			runeFrequencies[entry] = len(runeFrequencies)
		}
	}

	return runeFrequencies
}

type binaryList []*BinaryTree

func rankByBinaryCount(runeFrequencies map[rune]int) binaryList {

	list := make(binaryList, len(runeFrequencies))
	i := 0
	keys := make([]string, 0)
	for k, _ := range runeFrequencies {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)

	for _, k := range keys {
		r := []rune(k)[0]
		list[i] = &BinaryTree{
			Value: &r,
		}
		i++
	}

	return list
}

func buildBinaryTree(value string, list binaryList) *BinaryTree {

	for {
		first := list[0]
		list = list[1:]

		if len(list) == 0 {
			return first
		}

		second := list[0]
		list = list[1:]

		t := &BinaryTree{
			Left:  first,
			Right: second,
		}

		if len(list) == 0 {
			return t
		}

		list = append(list, t)
	}

}

func (s *BinaryTree) isLeaf() bool {
	return s.Value != nil
}

func (s *BinaryTree) Prefix() map[rune]*bitvector.BitVector {
	prefix := map[rune]*bitvector.BitVector{}
	left := s.Left
	if left.isLeaf() {
		vector := bitvector.NewBitVectorFromBool([]bool{false})
		prefix[rune(*left.Value)] = vector
	} else {
		m := left.Prefix()

		for r, v := range m {
			vector := bitvector.NewBitVectorFromVectorPadStart(v, 1)
			vector.Set(0, false)
			prefix[r] = vector
		}
	}

	right := s.Right
	if right.isLeaf() {
		vector := bitvector.NewBitVectorFromBool([]bool{true})
		prefix[rune(*right.Value)] = vector
	} else {
		m := right.Prefix()
		for r, v := range m {
			vector := bitvector.NewBitVectorFromVectorPadStart(v, 1)
			vector.Set(0, true)
			prefix[r] = vector
		}
	}

	return prefix
}
