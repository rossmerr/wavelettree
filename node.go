package wavelettree

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	value  *byte
	vector BitVector
}

func newNode(data []byte, prefix map[rune][]bool, parent *Node, depth int) *Node {

	vector, left, right, ok := NewVector(data, prefix, depth)

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

func (t *Node) isLeaf() bool {
	return t.vector == nil
}

func (t *Node) Access(i int) rune {
	if t.isLeaf() {
		return rune(*t.value)
	}

	c := t.vector[i]
	rank := t.vector.Rank(c, i)

	if c == false {
		return t.left.Access(rank)
	} else {
		return t.right.Access(rank)
	}
}

func (t *Node) Rank(prefix BitVector, offset int) int {
	if t.isLeaf() {
		return offset
	}

	c := prefix[0]

	rank := t.vector.Rank(c, offset)

	if c == false {
		return t.left.Rank(prefix[1:], rank)
	} else {
		return t.right.Rank(prefix[1:], rank)
	}
}

func (t *Node) Walk(prefix BitVector) Tree {

	if t.isLeaf() {
		return t
	}

	c := prefix[0]
	if c == false {
		return t.left.Walk(prefix[1:])
	} else {
		return t.right.Walk(prefix[1:])

	}
}

func (t *Node) Select(prefix BitVector, rank int) int {

	if t.isLeaf() {
		return t.parent.Select(prefix, rank)
	}
	i := prefix[len(prefix)-1]
	r := t.vector.Select(i, rank)

	if t.parent != nil {
		return t.parent.Select(prefix[:len(prefix)-1], r)
	}
	return r

}
