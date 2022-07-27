package wavelettree

type Tree interface {
	Access(i int) rune
	Rank(prefix Vector, offset int) int
	Prefix() map[rune]Vector
	Walk(prefix Vector) Tree
	Select(prefix Vector, rank int) int
}
