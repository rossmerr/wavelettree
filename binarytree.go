package wavelettree

type BinaryTree struct {
	parent *BinaryTree
	left   *BinaryTree
	right  *BinaryTree
	value  byte
	vector Vector
}

const (
	Left  = 0
	Right = 1
)

func NewBinaryTree(v Vector, left, right []byte) *BinaryTree {
	return newBinaryTree(v, left, right, nil)
}

func newBinaryTree(v Vector, left, right []byte, parent *BinaryTree) *BinaryTree {
	t := &BinaryTree{
		vector: v,
		parent: parent,
	}

	if len(left) > 0 {
		t.insertLeft(left)
	}
	if len(right) > 0 {
		t.insertRight(right)
	}
	return t
}

func (t *BinaryTree) isLeaf() bool {
	return t.vector == nil
}

func (t *BinaryTree) insertLeft(data []byte) {
	vector, left, right, end := NewVector(data)
	if end {
		t.left = &BinaryTree{
			value:  data[0],
			parent: t,
		}
		return
	}

	t.left = newBinaryTree(vector, left, right, t)

}

func (t *BinaryTree) insertRight(data []byte) {
	vector, left, right, end := NewVector(data)
	if end {
		t.right = &BinaryTree{
			value:  data[0],
			parent: t,
		}
		return
	}

	t.right = newBinaryTree(vector, left, right, t)

}

func (t *BinaryTree) Access(i int) rune {
	if t.isLeaf() {
		return rune(t.value)
	}

	c := t.vector[i]
	if c == Left {
		rank := t.vector.Rank(0, i)
		return t.left.Access(rank)
	} else {
		rank := t.vector.Rank(1, i)
		return t.right.Access(rank)
	}
}

func (t *BinaryTree) Prefix() map[rune]Vector {
	prefix := map[rune]Vector{}
	left := t.left
	if left.isLeaf() {
		prefix[rune(left.value)] = []byte{Left}
	} else {
		m := left.Prefix()
		for k, v := range m {
			prefix[k] = append([]byte{Left}, v...)
		}
	}

	right := t.right
	if right.isLeaf() {
		prefix[rune(right.value)] = []byte{Right}
	} else {
		m := right.Prefix()
		for k, v := range m {
			prefix[k] = append([]byte{Right}, v...)
		}
	}

	return prefix
}

func (t *BinaryTree) Rank(prefix Vector, offset int) int {
	if t.isLeaf() {
		return offset
	}

	c := prefix[0]

	rank := t.vector.Rank(c, offset)

	if c == Left {
		return t.left.Rank(prefix[1:], rank)
	} else {
		return t.right.Rank(prefix[1:], rank)
	}
}

func (t *BinaryTree) Walk(prefix Vector) Tree {

	if t.isLeaf() {
		return t
	}

	c := prefix[0]
	if c == Left {
		return t.left.Walk(prefix[1:])
	} else {
		return t.right.Walk(prefix[1:])

	}
}

func (t *BinaryTree) Select(prefix Vector, rank int) int {

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
