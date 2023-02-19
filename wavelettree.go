package wavelettree

import (
	"fmt"
)

type WaveletTree struct {
	root   *Node
	prefix Prefix
	n      int // Length of the root bitvector
}

func NewWaveletTree(value string, prefix Prefix) *WaveletTree {
	root := newNode([]byte(value), prefix, nil, 0)
	tree := &WaveletTree{
		root:   root,
		prefix: prefix,
		n:      root.Length(),
	}

	return tree
}

func NewBalancedWaveletTree(value string) *WaveletTree {
	prefix := NewBinaryTree(value)
	root := newNode([]byte(value), prefix, nil, 0)
	tree := &WaveletTree{
		root:   root,
		prefix: prefix,
		n:      root.Length(),
	}

	return tree
}

func NewHuffmanCodeWaveletTree(value string) *WaveletTree {
	prefix := NewHuffmanCodeTree(value)

	root := newNode([]byte(value), prefix, nil, 0)

	tree := &WaveletTree{
		root:   root,
		prefix: prefix,
		n:      root.Length(),
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

func (wt *WaveletTree) Length() int {
	return wt.n
}

func (wt WaveletTree) String() string {
	str := ""
	str += fmt.Sprintf(" length: %v", wt.n)

	if wt.root != nil {
		str += fmt.Sprintf(", root: %s", wt.root)
	}

	if wt.prefix != nil {
		str += fmt.Sprintf(", prefix: %+v", wt.prefix)
	}

	return fmt.Sprintf("{%s }", str)
}
