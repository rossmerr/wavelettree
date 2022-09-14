package wavelettree

import "sort"

type HuffmanCode struct {
	Left      *HuffmanCode
	Right     *HuffmanCode
	Value     rune
	Frequency int
}

func NewHuffmanCode(value string) *HuffmanCode {
	runeFrequencies := frequencyCount(value)
	huffmanList := rankByRuneCount(runeFrequencies)
	return buildHuffmanTree(huffmanList)
}

func (s *HuffmanCode) isLeaf() bool {
	return s.Value == -1
}

func (s *HuffmanCode) Prefix() map[rune]Vector {
	prefix := map[rune]Vector{}
	left := s.Left
	if left.isLeaf() {
		prefix[left.Value] = []byte{Left}
	} else {
		m := left.Prefix()
		for k, v := range m {
			prefix[k] = append([]byte{Left}, v...)
		}
	}

	right := s.Right
	if right.isLeaf() {
		prefix[right.Value] = []byte{Right}
	} else {
		m := right.Prefix()
		for k, v := range m {
			prefix[k] = append([]byte{Right}, v...)
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
		r := value[i]
		if _, ok := runeFrequencies[rune(r)]; ok {
			runeFrequencies[rune(r)] = runeFrequencies[rune(r)] + 1
		} else {
			runeFrequencies[rune(r)] = 1
		}
	}

	return runeFrequencies
}

func rankByRuneCount(runeFrequencies map[rune]int) huffmanList {
	list := make(huffmanList, len(runeFrequencies))
	i := 0
	for k, v := range runeFrequencies {
		list[i] = &HuffmanCode{Value: k, Frequency: v}
		i++
	}
	sort.Sort(list)
	return list
}

type huffmanList []*HuffmanCode

func (p huffmanList) Len() int           { return len(p) }
func (p huffmanList) Less(i, j int) bool { return p[i].Frequency < p[j].Frequency }
func (p huffmanList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
