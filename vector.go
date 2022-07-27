package wavelettree

type Vector []byte

func NewVector(data []byte) (vector Vector, left, right []byte, end bool) {
	keys := make(map[byte]int)

	for _, entry := range data {
		if _, ok := keys[entry]; !ok {
			keys[entry] = len(keys)
		}

		i := keys[entry]

		if int(i)%2 == Left {
			vector = append(vector, Left)
			left = append(left, entry)

		} else {
			vector = append(vector, Right)
			right = append(right, entry)

		}
	}

	return vector, left, right, len(keys) == 1
}

func NewVectorFromString(s string) (vector Vector, left, right []byte, end bool) {
	return NewVector([]byte(s))
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
