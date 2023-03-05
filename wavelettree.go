package wavelettree

import (
	"fmt"

	"github.com/rossmerr/wavelettree/prefixtree"
)

// WaveletTree is a succinct data structure to store strings in compressed space.
type WaveletTree struct {
	root   *node
	prefix *prefixtree.Prefix
	n      int // Length of the root bitvector
}

func NewWaveletTree(value []rune, prefix *prefixtree.Prefix) *WaveletTree {

	root := buildNode(value, prefix)
	tree := &WaveletTree{
		root:   root,
		prefix: prefix,
		n:      root.Length(),
	}

	return tree
}

func NewBalancedWaveletTree(value []rune) *WaveletTree {
	prefix := prefixtree.NewBinaryTree(value)
	root := buildNode(value, prefix)
	tree := &WaveletTree{
		root:   root,
		prefix: prefix,
		n:      root.Length(),
	}

	return tree
}

func NewHuffmanCodeWaveletTree(value []rune) *WaveletTree {
	prefix := prefixtree.NewHuffmanCodeTree(value)

	root := buildNode(value, prefix)

	tree := &WaveletTree{
		root:   root,
		prefix: prefix,
		n:      root.Length(),
	}

	return tree
}

// Access gets the run at the index.
func (wt *WaveletTree) Access(index int) rune {
	return wt.root.Access(index)
}

// Rank counts the number of times the rune occurs up to but not including the offset.
func (wt *WaveletTree) Rank(c rune, offset int) (int, error) {
	prefix := wt.prefix.Get(c)
	if prefix == nil {
		return 0, fmt.Errorf("rune '%v' code %v not found in prefix", string(c), c)

	}
	return wt.root.Rank(prefix, offset), nil
}

// Select returns the index of the rune with the given rank
func (wt *WaveletTree) Select(c rune, rank int) int {
	prefix := wt.prefix.Get(c)
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
