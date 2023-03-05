package wavelettree

import (
	"fmt"

	"github.com/rossmerr/bitvector"
	"github.com/rossmerr/wavelettree/prefixtree"
)

type node struct {
	parent *node
	left   *node
	right  *node
	value  *rune
	vector *bitvector.BitVector
}

func buildNode(data []rune, prefix *prefixtree.Prefix) *node {
	return buildChildNode(data, prefix, nil, 0)
}
func buildChildNode(data []rune, prefix *prefixtree.Prefix, parent *node, depth int) *node {
	vector := bitvector.NewBitVector(len(data))
	left, right := []rune{}, []rune{}

	for i, entry := range data {

		partitions := prefix.Get(entry)

		if depth >= partitions.Length() {
			return nil
		}

		c := partitions.Get(depth)
		vector.Set(i, c)
		if c {
			right = append(right, entry)
		} else {
			left = append(left, entry)
		}
	}

	t := &node{
		vector: vector,
		parent: parent,
	}

	if len(left) > 0 {
		n := buildChildNode(left, prefix, t, depth+1)

		if n != nil {
			t.left = n
		} else {
			t.left = &node{
				value:  &left[0],
				parent: t,
			}
		}

	}
	if len(right) > 0 {
		n := buildChildNode(right, prefix, t, depth+1)

		if n != nil {
			t.right = n
		} else {
			t.right = &node{
				value:  &right[0],
				parent: t,
			}
		}
	}
	return t
}

func (t *node) Length() int {
	return t.vector.Length()
}

func (t *node) isLeaf() bool {
	return t.vector == nil
}

func (t *node) Access(i int) rune {
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

func (t *node) Rank(prefix *bitvector.BitVector, offset int) int {

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

func (t *node) Walk(prefix *bitvector.BitVector) *node {

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

func (t *node) Select(prefix *bitvector.BitVector, rank int) int {

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

func (t node) String() string {
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
