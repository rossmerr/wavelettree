package wavelettree

type WaveletTree struct {
	root   Tree
	prefix map[rune]Vector
}

func NewWaveletTree(value string) *WaveletTree {
	return NewBalancedWaveletTree(value)
}

func NewBalancedWaveletTree(value string) *WaveletTree {
	vector, left, right, _ := NewVectorFromString(value)
	root := NewBinaryTree(vector, left, right)
	tree := &WaveletTree{
		root:   root,
		prefix: root.Prefix(),
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
