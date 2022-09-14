package wavelettree

type Vector []byte

func NewVector(data []byte, prefix map[rune]Vector, depth int) (vector Vector, left, right []byte, ok bool) {
	return NewVectorFromString(string(data), prefix, depth)
}

func NewVectorFromString(s string, prefix map[rune]Vector, depth int) (vector Vector, left, right []byte, ok bool) {
	ok = true
	for _, entry := range s {

		partitions := prefix[entry]

		if depth >= len(partitions) {
			ok = false
			return
		}

		c := partitions[depth]
		vector = append(vector, c)
		if c == Left {
			left = append(left, byte(entry))
		} else {
			right = append(right, byte(entry))
		}
	}
	return
}

func (v Vector) Rank(i byte, offset int) int {
	rank := 0

	for _, e := range v[:offset] {
		if e == i {
			rank++
		}

	}

	return rank
}

func (v Vector) Select(i byte, rank int) int {
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
