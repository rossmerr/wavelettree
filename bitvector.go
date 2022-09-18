package wavelettree

type BitVector []bool

func NewVector(data []byte, prefix map[rune]BitVector, depth int) (vector BitVector, left, right []byte, ok bool) {
	return NewVectorFromString(string(data), prefix, depth)
}

func NewVectorFromString(s string, prefix map[rune]BitVector, depth int) (vector BitVector, left, right []byte, ok bool) {
	ok = true
	for _, entry := range s {

		partitions := prefix[entry]

		if depth >= len(partitions) {
			ok = false
			return
		}

		c := partitions[depth]
		vector = append(vector, c)
		if c == false {
			left = append(left, byte(entry))
		} else {
			right = append(right, byte(entry))
		}
	}
	return
}

func (v BitVector) Rank(i bool, offset int) int {
	rank := 0

	for _, e := range v[:offset] {
		if e == i {
			rank++
		}

	}

	return rank
}

func (v BitVector) Select(i bool, rank int) int {
	offset := 0
	for c, e := range v {
		if offset == rank {
			return c
		}
		if e == i {
			offset++
		}

	}

	return 0
}

func (v BitVector) Concat(vectors []BitVector) BitVector {
	vector := []bool{}
	vector = append(vector, v...)
	for _, v := range vectors {
		vector = append(vector, v...)
	}

	return vector
}
