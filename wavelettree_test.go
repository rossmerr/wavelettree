package wavelettree

import (
	"fmt"
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
			wt:   NewBalancedWaveletTree([]rune("mississippi")),
		},
		{
			name: "binarytree mississippi",
			i:    1,
			want: rune('i'),
			wt:   NewBalancedWaveletTree([]rune("mississippi")),
		},
		{
			name: "binarytree mississippi",
			i:    2,
			want: rune('s'),
			wt:   NewBalancedWaveletTree([]rune("mississippi")),
		}, {
			name: "binarytree mississippi",
			i:    3,
			want: rune('s'),
			wt:   NewBalancedWaveletTree([]rune("mississippi")),
		}, {
			name: "binarytree mississippi",
			i:    4,
			want: rune('i'),
			wt:   NewBalancedWaveletTree([]rune("mississippi")),
		}, {
			name: "binarytree mississippi",
			i:    5,
			want: rune('s'),
			wt:   NewBalancedWaveletTree([]rune("mississippi")),
		}, {
			name: "binarytree mississippi",
			i:    6,
			want: rune('s'),
			wt:   NewBalancedWaveletTree([]rune("mississippi")),
		},
		{
			name: "binarytree mississippi",
			i:    7,
			want: rune('i'),
			wt:   NewBalancedWaveletTree([]rune("mississippi")),
		},
		{
			name: "binarytree mississippi",
			i:    8,
			want: rune('p'),
			wt:   NewBalancedWaveletTree([]rune("mississippi")),
		},
		{
			name: "binarytree mississippi",
			i:    9,
			want: rune('p'),
			wt:   NewBalancedWaveletTree([]rune("mississippi")),
		},
		{
			name: "binarytree mississippi",
			i:    10,
			want: rune('i'),
			wt:   NewBalancedWaveletTree([]rune("mississippi")),
		},
		{
			name: "huffman mississippi",
			i:    0,
			want: rune('m'),
			wt:   NewHuffmanCodeWaveletTree([]rune("mississippi")),
		},
		{
			name: "huffman mississippi",
			i:    1,
			want: rune('i'),
			wt:   NewHuffmanCodeWaveletTree([]rune("mississippi")),
		},
		{
			name: "huffman mississippi",
			i:    2,
			want: rune('s'),
			wt:   NewHuffmanCodeWaveletTree([]rune("mississippi")),
		},
		{
			name: "huffman mississippi",
			i:    3,
			want: rune('s'),
			wt:   NewHuffmanCodeWaveletTree([]rune("mississippi")),
		},
		{
			name: "huffman mississippi",
			i:    4,
			want: rune('i'),
			wt:   NewHuffmanCodeWaveletTree([]rune("mississippi")),
		},
		{
			name: "huffman mississippi",
			i:    5,
			want: rune('s'),
			wt:   NewHuffmanCodeWaveletTree([]rune("mississippi")),
		},
		{
			name: "huffman mississippi",
			i:    6,
			want: rune('s'),
			wt:   NewHuffmanCodeWaveletTree([]rune("mississippi")),
		},
		{
			name: "huffman mississippi",
			i:    7,
			want: rune('i'),
			wt:   NewHuffmanCodeWaveletTree([]rune("mississippi")),
		},

		{
			name: "huffman mississippi",
			i:    8,
			want: rune('p'),
			wt:   NewHuffmanCodeWaveletTree([]rune("mississippi")),
		},
		{
			name: "huffman mississippi",
			i:    9,
			want: rune('p'),
			wt:   NewHuffmanCodeWaveletTree([]rune("mississippi")),
		},
		{
			name: "huffman mississippi",
			i:    10,
			want: rune('i'),
			wt:   NewHuffmanCodeWaveletTree([]rune("mississippi")),
		},
		{
			name: "huffman quick fox",
			i:    0,
			want: rune('T'),
			wt:   NewHuffmanCodeWaveletTree([]rune("The quick brown fox jumps over the lazy dog")),
		},
		{
			name: "huffman quick fox",
			i:    0,
			want: rune('T'),
			wt:   NewHuffmanCodeWaveletTree([]rune("The quick brown fox jumps over the lazy dog")),
		},
		{
			name: "huffman quick fox",
			i:    20,
			want: rune('j'),
			wt:   NewHuffmanCodeWaveletTree([]rune("The quick brown fox jumps over the lazy dog")),
		},
		{
			name: "huffman quick fox",
			i:    42,
			want: rune('g'),
			wt:   NewHuffmanCodeWaveletTree([]rune("The quick brown fox jumps over the lazy dog")),
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
		name    string
		wt      *WaveletTree
		args    args
		want    int
		wantErr error
	}{

		{
			name: "binarytree mississippi",
			wt:   NewBalancedWaveletTree([]rune("mississippi")),
			args: args{
				c:      'm',
				offset: 0,
			},
			want: 0,
		},
		{
			name: "binarytree mississippi",
			wt:   NewBalancedWaveletTree([]rune("mississippi")),
			args: args{
				c:      'm',
				offset: 1,
			},
			want: 1,
		},
		{
			name: "binarytree mississippi",
			wt:   NewBalancedWaveletTree([]rune("mississippi")),
			args: args{
				c:      'i',
				offset: 6,
			},
			want: 2,
		},
		{
			name: "binarytree mississippi",
			wt:   NewBalancedWaveletTree([]rune("mississippi")),
			args: args{
				c:      'i',
				offset: 7,
			},
			want: 2,
		},
		{
			name: "binarytree mississippi",
			wt:   NewBalancedWaveletTree([]rune("mississippi")),
			args: args{
				c:      'i',
				offset: 8,
			},
			want: 3,
		},
		{
			name: "huffman mississippi",
			wt:   NewHuffmanCodeWaveletTree([]rune("mississippi")),
			args: args{
				c:      'i',
				offset: 6,
			},
			want: 2,
		},
		{
			name: "huffman mississippi",
			wt:   NewHuffmanCodeWaveletTree([]rune("mississippi")),
			args: args{
				c:      'i',
				offset: 7,
			},
			want: 2,
		},
		{
			name: "huffman mississippi",
			wt:   NewHuffmanCodeWaveletTree([]rune("mississippi")),
			args: args{
				c:      'i',
				offset: 8,
			},
			want: 3,
		},
		{
			name: "huffman mississippi",
			wt:   NewHuffmanCodeWaveletTree([]rune("mississippi")),
			args: args{
				c:      'i',
				offset: 11,
			},
			want: 4,
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
		{
			name: "huffman quick fox",
			wt:   NewHuffmanCodeWaveletTree([]rune("The quick brown fox jumps over the lazy dog")),
			args: args{
				c:      'u',
				offset: 42,
			},
			want: 2,
		},
		{
			name: "huffman quick fox",
			wt:   NewHuffmanCodeWaveletTree([]rune("The quick brown fox jumps over the lazy dog")),
			args: args{
				c:      ' ',
				offset: 42,
			},
			want: 8,
		},
		{
			name: "huffman quick fox",
			wt:   NewHuffmanCodeWaveletTree([]rune("The quick brown fox jumps over the lazy dog")),
			args: args{
				c:      'u',
				offset: 15,
			},
			want: 1,
		},

		{
			name: "huffman quick fox",
			wt:   NewHuffmanCodeWaveletTree([]rune("The quick brown fox jumps over the lazy dog")),
			args: args{
				c:      '@',
				offset: 42,
			},
			want:    0,
			wantErr: fmt.Errorf("rune '@' code 64 not found in prefix"),
		},
		{
			name: "huffman quick fox",
			wt:   NewHuffmanCodeWaveletTree([]rune("The quick brown fox jumps over the lazy dog")),
			args: args{
				c:      's',
				offset: 42,
			},
			want: 1,
		},
		{
			name: "huffman quick fox",
			wt:   NewHuffmanCodeWaveletTree([]rune("The quick brown fox jumps over the lazy dog")),
			args: args{
				c:      's',
				offset: 0,
			},
			want: 0,
		},
		{
			name: "BWT encoded huffman quick fox",
			wt:   NewHuffmanCodeWaveletTree([]rune("\x03        Tabcdeeefghhijklmnoooopqrrstuuvwxyz")),
			args: args{
				c:      's',
				offset: 0,
			},
			want: 0,
		},
		{
			name: "BWT encoded huffman quick fox",
			wt:   NewHuffmanCodeWaveletTree([]rune("\x03        Tabcdeeefghhijklmnoooopqrrstuuvwxyz")),
			args: args{
				c:      's',
				offset: 42,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.wt.Rank(tt.args.c, tt.args.offset)
			if tt.wantErr != nil && err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("WaveletTree.Rank() Error = %v, want %v", got, tt.want)
			} else if got != tt.want {
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
			name: "BWT encoded huffman quick fox",
			wt:   NewHuffmanCodeWaveletTree([]rune("\x03        Tabcdeeefghhijklmnoooopqrrstuuvwxyz")),
			args: args{
				c:    'o',
				rank: 0,
			},
			want: 27,
		},
		{
			name: "BWT encoded huffman quick fox",
			wt:   NewHuffmanCodeWaveletTree([]rune("\x03        Tabcdeeefghhijklmnoooopqrrstuuvwxyz")),
			args: args{
				c:    'o',
				rank: 4,
			},
			want: 30,
		},
		{
			name: "huffman quick fox",
			wt:   NewHuffmanCodeWaveletTree([]rune("The quick brown fox jumps over the lazy dog")),
			args: args{
				c:    ' ',
				rank: 2,
			},
			want: 15,
		},
		{
			name: "huffman quick fox",
			wt:   NewHuffmanCodeWaveletTree([]rune("The quick brown fox jumps over the lazy dog")),
			args: args{
				c:    'u',
				rank: 0,
			},
			want: 5,
		},
		{
			name: "huffman quick fox",
			wt:   NewHuffmanCodeWaveletTree([]rune("The quick brown fox jumps over the lazy dog")),
			args: args{
				c:    'u',
				rank: 1,
			},
			want: 21,
		},
		{
			name: "binarytree mississippi",
			wt:   NewBalancedWaveletTree([]rune("mississippi")),
			args: args{
				c:    's',
				rank: 0,
			},
			want: 2,
		},
		{
			name: "binarytree mississippi",
			wt:   NewBalancedWaveletTree([]rune("mississippi")),
			args: args{
				c:    's',
				rank: 1,
			},
			want: 3,
		},
		{
			name: "binarytree mississippi",
			wt:   NewBalancedWaveletTree([]rune("mississippi")),
			args: args{
				c:    's',
				rank: 3,
			},
			want: 6,
		},
		{
			name: "huffman mississippi",
			wt:   NewHuffmanCodeWaveletTree([]rune("mississippi")),
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
			got := tt.wt.Select(tt.args.c, tt.args.rank)
			if got != tt.want {
				t.Errorf("WaveletTree.Select(%v) = %v, want %v", i, got, tt.want)
			}
		})
	}
}
