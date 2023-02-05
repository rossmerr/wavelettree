package wavelettree

import (
	"reflect"
	"testing"

	"github.com/rossmerr/bitvector"
)

func TestBinaryTree_Prefix(t *testing.T) {

	tests := []struct {
		name  string
		value string
		want  map[rune]*bitvector.BitVector
	}{
		{
			name:  "Binary Prefix",
			value: "mississippi",
			want: map[rune]*bitvector.BitVector{
				'i': bitvector.NewBitVectorFromBool([]bool{false, false}),
				'm': bitvector.NewBitVectorFromBool([]bool{false, true}),
				'p': bitvector.NewBitVectorFromBool([]bool{true, false}),
				's': bitvector.NewBitVectorFromBool([]bool{true, true}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewBinaryTree(tt.value)
			if got := s.Prefix(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryTree.Prefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
