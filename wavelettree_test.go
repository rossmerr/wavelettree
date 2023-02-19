package wavelettree

import (
	"testing"

	"github.com/rossmerr/bitvector"
)

func TestWaveletTree_Access(t *testing.T) {
	tests := []struct {
		name string
		i    int
		want rune
		wt   *WaveletTree
	}{
		{
			name: "binarytree mississippi",
			i:    0,
			want: rune('m'),
			wt:   NewBalancedWaveletTree("mississippi"),
		},
		{
			name: "binarytree mississippi",
			i:    1,
			want: rune('i'),
			wt:   NewBalancedWaveletTree("mississippi"),
		},
		{
			name: "binarytree mississippi",
			i:    2,
			want: rune('s'),
			wt:   NewBalancedWaveletTree("mississippi"),
		}, {
			name: "binarytree mississippi",
			i:    3,
			want: rune('s'),
			wt:   NewBalancedWaveletTree("mississippi"),
		}, {
			name: "binarytree mississippi",
			i:    4,
			want: rune('i'),
			wt:   NewBalancedWaveletTree("mississippi"),
		}, {
			name: "binarytree mississippi",
			i:    5,
			want: rune('s'),
			wt:   NewBalancedWaveletTree("mississippi"),
		}, {
			name: "binarytree mississippi",
			i:    6,
			want: rune('s'),
			wt:   NewBalancedWaveletTree("mississippi"),
		},
		{
			name: "binarytree mississippi",
			i:    7,
			want: rune('i'),
			wt:   NewBalancedWaveletTree("mississippi"),
		},
		{
			name: "binarytree mississippi",
			i:    8,
			want: rune('p'),
			wt:   NewBalancedWaveletTree("mississippi"),
		},
		{
			name: "binarytree mississippi",
			i:    9,
			want: rune('p'),
			wt:   NewBalancedWaveletTree("mississippi"),
		},
		{
			name: "binarytree mississippi",
			i:    10,
			want: rune('i'),
			wt:   NewBalancedWaveletTree("mississippi"),
		},
		{
			name: "huffman mississippi",
			i:    0,
			want: rune('m'),
			wt:   NewHuffmanCodeWaveletTree("mississippi"),
		},
		{
			name: "huffman mississippi",
			i:    1,
			want: rune('i'),
			wt:   NewHuffmanCodeWaveletTree("mississippi"),
		},
		{
			name: "huffman mississippi",
			i:    2,
			want: rune('s'),
			wt:   NewHuffmanCodeWaveletTree("mississippi"),
		},
		{
			name: "huffman mississippi",
			i:    3,
			want: rune('s'),
			wt:   NewHuffmanCodeWaveletTree("mississippi"),
		},
		{
			name: "huffman mississippi",
			i:    4,
			want: rune('i'),
			wt:   NewHuffmanCodeWaveletTree("mississippi"),
		},
		{
			name: "huffman mississippi",
			i:    5,
			want: rune('s'),
			wt:   NewHuffmanCodeWaveletTree("mississippi"),
		},
		{
			name: "huffman mississippi",
			i:    6,
			want: rune('s'),
			wt:   NewHuffmanCodeWaveletTree("mississippi"),
		},
		{
			name: "huffman mississippi",
			i:    7,
			want: rune('i'),
			wt:   NewHuffmanCodeWaveletTree("mississippi"),
		},

		{
			name: "huffman mississippi",
			i:    8,
			want: rune('p'),
			wt:   NewHuffmanCodeWaveletTree("mississippi"),
		},
		{
			name: "huffman mississippi",
			i:    9,
			want: rune('p'),
			wt:   NewHuffmanCodeWaveletTree("mississippi"),
		},
		{
			name: "huffman mississippi",
			i:    10,
			want: rune('i'),
			wt:   NewHuffmanCodeWaveletTree("mississippi"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.wt.Access(tt.i)
			if got != tt.want {
				t.Errorf("WaveletTree.Access(%v) = %v, want %v", tt.i, string(got), string(tt.want))
			}
		})
	}
}

func TestWaveletTree_Rank(t *testing.T) {

	type args struct {
		c      rune
		offset int
	}
	tests := []struct {
		name string
		wt   *WaveletTree
		args args
		want int
	}{
		{
			name: "binarytree mississippi",
			wt:   NewBalancedWaveletTree("mississippi"),
			args: args{
				c:      'm',
				offset: 0,
			},
			want: 0,
		},
		{
			name: "binarytree mississippi",
			wt:   NewBalancedWaveletTree("mississippi"),
			args: args{
				c:      'm',
				offset: 1,
			},
			want: 1,
		},
		{
			name: "binarytree mississippi",
			wt:   NewBalancedWaveletTree("mississippi"),
			args: args{
				c:      'i',
				offset: 6,
			},
			want: 2,
		},
		{
			name: "huffman mississippi",
			wt:   NewHuffmanCodeWaveletTree("mississippi"),
			args: args{
				c:      'i',
				offset: 6,
			},
			want: 2,
		},
		{
			name: "00110110110",
			wt: func() *WaveletTree {
				root := &Node{
					vector: func() *bitvector.BitVector {
						vector := bitvector.NewBitVectorFromBool([]bool{false, false, true, true, false, true, true, false, true, true, false})
						return vector
					}(),
				}
				level1 := &Node{
					vector: func() *bitvector.BitVector {
						vector := bitvector.NewBitVectorFromBool([]bool{true, false, false, false, false})
						return vector
					}(),
					parent: root,
				}
				level2 := &Node{
					parent: level1,
				}
				level1.left = level2
				root.left = level1
				vector := bitvector.NewBitVectorFromBool([]bool{false, false})
				wt := &WaveletTree{
					root:   root,
					prefix: map[rune]*bitvector.BitVector{'i': vector},
				}
				return wt
			}(),
			args: args{
				c:      'i',
				offset: 6,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.wt.Rank(tt.args.c, tt.args.offset); got != tt.want {
				t.Errorf("WaveletTree.Rank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWaveletTree_Select(t *testing.T) {
	type args struct {
		c    rune
		rank int
	}
	tests := []struct {
		name string
		wt   *WaveletTree
		args args
		want int
	}{
		{
			name: "binarytree mississippi",
			wt:   NewBalancedWaveletTree("mississippi"),
			args: args{
				c:    's',
				rank: 3,
			},
			want: 6,
		},
		{
			name: "huffman mississippi",
			wt:   NewHuffmanCodeWaveletTree("mississippi"),
			args: args{
				c:    's',
				rank: 3,
			},
			want: 6,
		},
		{
			name: "00110110110",
			wt: func() *WaveletTree {
				root := &Node{
					vector: func() *bitvector.BitVector {
						vector := bitvector.NewBitVectorFromBool([]bool{false, false, true, true, false, true, true, false, true, true, false})
						return vector
					}(),
				}
				level1 := &Node{
					vector: func() *bitvector.BitVector {
						vector := bitvector.NewBitVectorFromBool([]bool{true, true, true, true, false, false})
						return vector
					}(),
					parent: root,
				}
				level2 := &Node{
					parent: level1,
				}
				level1.right = level2
				root.right = level1
				vector := bitvector.NewBitVectorFromBool([]bool{true, true})
				wt := &WaveletTree{
					root: root,
					prefix: map[rune]*bitvector.BitVector{
						's': vector,
					},
				}
				return wt
			}(), args: args{
				c:    's',
				rank: 3,
			},
			want: 6,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.wt.Select(tt.args.c, tt.args.rank); got != tt.want {
				t.Errorf("WaveletTree.Select(%v) = %v, want %v", i, got, tt.want)
			}
		})
	}
}
