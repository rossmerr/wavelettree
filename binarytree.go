package wavelettree

type BinaryTree struct {
	Left  *BinaryTree
	Right *BinaryTree
	Value *byte
}

func NewBinaryTree(value string) *BinaryTree {
	runeFrequencies := binaryCount(value)
	binaryList := rankByBinaryCount(runeFrequencies)
	return buildBinaryTree(value, binaryList)
}

func binaryCount(value string) map[byte]int {
	runeFrequencies := make(map[byte]int)

	for _, entry := range value {
		if _, ok := runeFrequencies[byte(entry)]; !ok {
			runeFrequencies[byte(entry)] = len(runeFrequencies)
		}
	}

	return runeFrequencies
}

type binaryList []*BinaryTree

func rankByBinaryCount(runeFrequencies map[byte]int) binaryList {

	list := make(binaryList, len(runeFrequencies))
	i := 0
	for k := range runeFrequencies {
		v := k
		list[i] = &BinaryTree{
			Value: &v,
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

func (s *BinaryTree) Prefix() map[rune][]bool {
	prefix := map[rune][]bool{}
	left := s.Left
	if left.isLeaf() {
		prefix[rune(*left.Value)] = []bool{false}
	} else {
		m := left.Prefix()
		for k, v := range m {
			prefix[k] = append([]bool{false}, v...)
		}
	}

	right := s.Right
	if right.isLeaf() {
		prefix[rune(*right.Value)] = []bool{true}
	} else {
		m := right.Prefix()
		for k, v := range m {
			prefix[k] = append([]bool{true}, v...)
		}
	}

	return prefix
}
