package wavelettree

import "sort"

type HuffmanCode struct {
	Left      *HuffmanCode
	Right     *HuffmanCode
	Key       byte
	Frequency int
}

func NewHuffmanCode(value string) *HuffmanCode {
	return newHuffmanCode(value)
}

func newHuffmanCode(value string) *HuffmanCode {
	wordFrequencies := frequencyCount(value)
	pairList := rankByRuneCount(wordFrequencies)
	tree := buildTree(pairList)

	return tree
}

func buildTree(pairList PairList) *HuffmanCode {
	for {
		last := pairList[0]
		pairList = pairList[1:]

		if len(pairList) == 0 {
			return last
		}

		penultimate := pairList[0]
		pairList = pairList[1:]

		sum := last.Frequency + penultimate.Frequency

		t := &HuffmanCode{
			Frequency: sum,
			Left:      last,
			Right:     penultimate,
		}

		if len(pairList) == 0 {
			return t
		}

		for _, pair := range pairList {
			if pair.Frequency >= sum {
				pairList = append([]*HuffmanCode{t}, pairList...)
			} else {
				pairList = append(pairList, t)

			}
			break
		}

	}
}

func frequencyCount(value string) map[byte]int {
	wordFrequencies := make(map[byte]int)
	for i := range value {
		r := value[i]
		if _, ok := wordFrequencies[r]; ok {
			wordFrequencies[r] = wordFrequencies[r] + 1
		} else {
			wordFrequencies[r] = 1
		}
	}

	return wordFrequencies
}

func rankByRuneCount(wordFrequencies map[byte]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = &HuffmanCode{Key: k, Frequency: v}
		i++
	}
	sort.Sort(pl)
	return pl
}

type PairList []*HuffmanCode

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Frequency < p[j].Frequency }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
