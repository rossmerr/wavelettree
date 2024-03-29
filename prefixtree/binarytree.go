package prefixtree

import (
	"github.com/rossmerr/bitvector"
)

type BinaryTree struct {
	Left  *BinaryTree
	Right *BinaryTree
	Value *rune
}

func NewBinaryTree(value []rune) *Prefix {
	runeFrequencies, keys := binaryCount(value)
	binaryList := rankByBinaryCount(runeFrequencies, keys)
	tree := buildBinaryTree(binaryList)
	return tree.prefix()
}

func binaryCount(value []rune) (map[rune]int, []rune) {
	runeFrequencies := make(map[rune]int)
	keys := make([]rune, 0)

	for _, r := range value {
		if _, ok := runeFrequencies[r]; !ok {
			runeFrequencies[r] = len(runeFrequencies)
			keys = append(keys, r)
		}
	}

	return runeFrequencies, keys
}

type binaryList []*BinaryTree

func rankByBinaryCount(runeFrequencies map[rune]int, keys []rune) binaryList {
	list := make(binaryList, len(runeFrequencies))

	for i, r := range keys {
		v := r
		list[i] = &BinaryTree{
			Value: &v,
		}
	}

	return list
}

func buildBinaryTree(list binaryList) *BinaryTree {

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

func (s *BinaryTree) prefix() *Prefix {
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
