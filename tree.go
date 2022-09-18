package wavelettree

type Tree interface {
	Access(i int) rune
	Rank(prefix BitVector, offset int) int
	Walk(prefix BitVector) Tree
	Select(prefix BitVector, rank int) int
}
