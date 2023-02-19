package prefixtree

import (
	"reflect"
	"testing"

	"github.com/rossmerr/bitvector"
)

func TestBinaryTree_Prefix(t *testing.T) {

	tests := []struct {
		name  string
		value string
		want  Prefix
	}{
		{
			name:  "BinaryTree Prefix",
			value: "mississippi",
			want: Prefix{
				'i': bitvector.NewBitVectorFromBool([]bool{false, true}),
				'm': bitvector.NewBitVectorFromBool([]bool{false, false}),
				'p': bitvector.NewBitVectorFromBool([]bool{true, true}),
				's': bitvector.NewBitVectorFromBool([]bool{true, false}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBinaryTree(tt.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryTree.Prefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
