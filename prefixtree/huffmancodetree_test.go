package prefixtree

import (
	"reflect"
	"testing"

	"github.com/rossmerr/bitvector"
)

func TestHuffmanCodeTree_Prefix(t *testing.T) {

	tests := []struct {
		name  string
		value string
		want  Prefix
	}{
		{
			name:  "HuffmanCodeTree Prefix",
			value: "mississippi",
			want: Prefix{
				'i': bitvector.NewBitVectorFromBool([]bool{true, true}),
				'm': bitvector.NewBitVectorFromBool([]bool{true, false, false}),
				'p': bitvector.NewBitVectorFromBool([]bool{true, false, true}),
				's': bitvector.NewBitVectorFromBool([]bool{false}),
			},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHuffmanCodeTree([]rune(tt.value))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryTree.Prefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
