package wavelettree

import (
	"reflect"
	"testing"

	"github.com/rossmerr/bitvector"
)

func TestNewHuffmanCode(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  *HuffmanCode
	}{
		{
			name:  "Basic",
			value: "BCAADDDCCACACAC",
			want: &HuffmanCode{
				Frequency: 15,
				Left: &HuffmanCode{
					Value: func() *rune {
						r := rune('C')
						return &r
					}(),
					Frequency: 6,
				},
				Right: &HuffmanCode{
					Frequency: 9,
					Left: &HuffmanCode{
						Frequency: 4,
						Left: &HuffmanCode{
							Frequency: 1,
							Value: func() *rune {
								r := rune('B')
								return &r
							}(),
						},
						Right: &HuffmanCode{
							Frequency: 3,
							Value: func() *rune {
								r := rune('D')
								return &r
							}(),
						},
					},
					Right: &HuffmanCode{
						Frequency: 5,
						Value: func() *rune {
							r := rune('A')
							return &r
						}(),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHuffmanCode(tt.value); !reflect.DeepEqual(got, tt.want) {

				t.Errorf("NewHuffmanCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHuffmanCode_Prefix(t *testing.T) {

	tests := []struct {
		name  string
		value string
		want  map[rune]*bitvector.BitVector
	}{
		{
			name:  "HuffmanCode Prefix",
			value: "mississippi",
			want: map[rune]*bitvector.BitVector{
				'i': bitvector.NewBitVectorFromBool([]bool{true, true}),
				'm': bitvector.NewBitVectorFromBool([]bool{true, false, false}),
				'p': bitvector.NewBitVectorFromBool([]bool{true, false, true}),
				's': bitvector.NewBitVectorFromBool([]bool{false}),
			},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewHuffmanCode(tt.value)
			if got := s.Prefix(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryTree.Prefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
