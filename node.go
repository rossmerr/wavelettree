package wavelettree

import (
	"fmt"

	"github.com/rossmerr/bitvector"
	"github.com/rossmerr/wavelettree/prefixtree"
)

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	value  *byte
	vector *bitvector.BitVector
}

func newBitVectorFromBytes(data []byte, prefix prefixtree.Prefix, depth int) (vector *bitvector.BitVector, left, right []byte, ok bool) {
	return newBitVectorFromString(string(data), prefix, depth)
}

func newBitVectorFromString(s string, prefix prefixtree.Prefix, depth int) (vector *bitvector.BitVector, left, right []byte, ok bool) {
	ok = true
	vector = bitvector.NewBitVector(len(s))
	for i, entry := range s {

		partitions := prefix[entry]

		if depth >= partitions.Length() {
			ok = false
			return
		}

		c := partitions.Get(depth)
		vector.Set(i, c)
		if c {
			right = append(right, byte(entry))
		} else {
			left = append(left, byte(entry))
		}
	}
	return
}

func newNode(data []byte, prefix prefixtree.Prefix, parent *Node, depth int) *Node {

	vector, left, right, ok := newBitVectorFromBytes(data, prefix, depth)

	if !ok {
		return nil
	}

	t := &Node{
		vector: vector,
		parent: parent,
	}

	if len(left) > 0 {
		n := newNode(left, prefix, t, depth+1)

		if n != nil {
			t.left = n
		} else {
			t.left = &Node{
				value:  &left[0],
				parent: t,
			}
		}

	}
	if len(right) > 0 {
		n := newNode(right, prefix, t, depth+1)

		if n != nil {
			t.right = n
		} else {
			t.right = &Node{
				value:  &right[0],
				parent: t,
			}
		}
	}
	return t
}

func (t *Node) Length() int {
	return t.vector.Length()
}

func (t *Node) isLeaf() bool {
	return t.vector == nil
}

func (t *Node) Access(i int) rune {
	if t.isLeaf() {
		return rune(*t.value)
	}

	c := t.vector.Get(i)

	rank := t.vector.Rank(c, i)

	if c {
		return t.right.Access(rank)
	} else {
		return t.left.Access(rank)
	}
}

func (t *Node) Rank(prefix *bitvector.BitVector, offset int) int {

	c := prefix.Get(0)

	rank := t.vector.Rank(c, offset)

	vector := bitvector.NewBitVector(prefix.Length() - 1)
	if prefix.Length() > 1 {
		prefix.Copy(vector, 1, prefix.Length())
	} else {
		return rank
	}

	if c {
		return t.right.Rank(vector, rank)
	} else {
		return t.left.Rank(vector, rank)
	}
}

func (t *Node) Walk(prefix *bitvector.BitVector) *Node {

	if t.isLeaf() {
		return t
	}

	c := prefix.Get(0)

	vector := bitvector.NewBitVector(prefix.Length() - 1)
	if prefix.Length() > 1 {
		prefix.Copy(vector, 1, prefix.Length())
	}

	if c {
		return t.right.Walk(vector)
	} else {
		return t.left.Walk(vector)
	}
}

func (t *Node) Select(prefix *bitvector.BitVector, rank int) int {

	if t.isLeaf() {
		return t.parent.Select(prefix, rank)
	}
	i := prefix.Get(prefix.Length() - 1)

	r := t.vector.Select(i, rank)

	if t.parent != nil {

		vector := bitvector.NewBitVector(prefix.Length() - 1)
		prefix.Copy(vector, 0, prefix.Length()-1)

		return t.parent.Select(vector, r)
	}
	return r

}

func (t Node) String() string {
	str := ""
	if t.left != nil {
		str += fmt.Sprintf(" left: %s", t.left)
	}
	if t.right != nil {
		str += fmt.Sprintf(" right: %s", t.right)
	}

	if t.value != nil {
		str += fmt.Sprintf(" value: %s", string(*t.value))
	}

	return fmt.Sprintf("{%s }", str)
}
