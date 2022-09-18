package wavelettree

import (
	"testing"
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
			i:    4,
			want: rune('i'),
			wt:   NewWaveletTree("mississippi"),
		},
		{
			name: "binarytree mississippi",
			i:    8,
			want: rune('p'),
			wt:   NewWaveletTree("mississippi"),
		},
		{
			name: "huffman mississippi",
			i:    4,
			want: rune('i'),
			wt:   NewHuffmanCodeWaveletTree("mississippi"),
		},
		{
			name: "huffman mississippi",
			i:    8,
			want: rune('p'),
			wt:   NewHuffmanCodeWaveletTree("mississippi"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.wt.Access(tt.i); got != tt.want {
				t.Errorf("WaveletTree.Access() = %v, want %v", got, tt.want)
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
			wt: func() *WaveletTree {
				return NewWaveletTree("mississippi")

			}(),
			args: args{
				c:      'i',
				offset: 6,
			},
			want: 2,
		},
		{
			name: "huffman mississippi",
			wt: func() *WaveletTree {
				return NewHuffmanCodeWaveletTree("mississippi")

			}(),
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
					vector: []bool{false, false, true, true, false, true, true, false, true, true, false},
				}
				level1 := &Node{
					vector: []bool{true, false, false, false, false},
					parent: root,
				}
				// i := byte('i')
				level2 := &Node{
					// value:  &i,
					parent: level1,
				}
				level1.left = level2
				root.left = level1
				wt := &WaveletTree{
					root: root,
					prefix: map[rune]BitVector{
						'i': {false, false},
					},
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
			wt: func() *WaveletTree {
				return NewWaveletTree("mississippi")
			}(),
			args: args{
				c:    's',
				rank: 3,
			},
			want: 6,
		},
		{
			name: "huffman mississippi",
			wt: func() *WaveletTree {
				return NewHuffmanCodeWaveletTree("mississippi")
			}(),
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
					vector: []bool{false, false, true, true, false, true, true, false, true, true, false},
				}
				level1 := &Node{
					vector: []bool{true, true, true, true, false, false},
					parent: root,
				}
				// s := byte('s')
				level2 := &Node{
					// value:  &s,
					parent: level1,
				}
				level1.right = level2
				root.right = level1
				wt := &WaveletTree{
					root: root,
					prefix: map[rune]BitVector{
						's': {true, true},
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.wt.Select(tt.args.c, tt.args.rank); got != tt.want {
				t.Errorf("WaveletTree.Select() = %v, want %v", got, tt.want)
			}
		})
	}
}
