package wavelettree

type WaveletTree struct {
	root   Tree
	prefix map[rune][]bool
}

func NewWaveletTree(value string) *WaveletTree {
	return NewBalancedWaveletTree(value)
}

func NewBalancedWaveletTree(value string) *WaveletTree {
	br := NewBinaryTree(value)
	prefix := br.Prefix()
	root := newNode([]byte(value), prefix, nil, 0)

	tree := &WaveletTree{
		root:   root,
		prefix: prefix,
	}

	return tree
}

func NewHuffmanCodeWaveletTree(value string) *WaveletTree {
	hc := NewHuffmanCode(value)
	prefix := hc.Prefix()
	root := newNode([]byte(value), prefix, nil, 0)
	tree := &WaveletTree{
		root:   root,
		prefix: prefix,
	}

	return tree
}

func (wt *WaveletTree) Access(i int) rune {
	return wt.root.Access(i)
}

func (wt *WaveletTree) Rank(c rune, offset int) int {
	prefix := wt.prefix[c]
	return wt.root.Rank(prefix, offset)
}

func (wt *WaveletTree) Select(c rune, rank int) int {
	prefix := wt.prefix[c]
	start := wt.root.Walk(prefix)

	return start.Select(prefix, rank)
}
