package wavelettree

import (
	"testing"
)

func TestWaveletTree_Access(t *testing.T) {
	type fields struct {
		root Tree
	}
	type args struct {
		i int
	}
	tests := []struct {
		name  string
		value string
		i     int
		want  rune
	}{
		{
			name:  "mississippi",
			value: "mississippi",
			i:     4,
			want:  rune('i'),
		},
		{
			name:  "mississippi",
			value: "mississippi",
			i:     8,
			want:  rune('p'),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wt := NewWaveletTree(tt.value)
			if got := wt.Access(tt.i); got != tt.want {
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
			name: "mississippi",
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
			name: "00110110110",
			wt: func() *WaveletTree {
				root := &Node{
					vector: []byte{0, 0, 1, 1, 0, 1, 1, 0, 1, 1, 0},
				}
				level1 := &Node{
					vector: []byte{1, 0, 0, 0, 0},
					parent: root,
				}
				level2 := &Node{
					value:  'i',
					parent: level1,
				}
				level1.left = level2
				root.left = level1
				wt := &WaveletTree{
					root: root,
					prefix: map[rune]Vector{
						'i': []byte{0, 0},
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
			name: "mississippi",
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
			name: "00110110110",
			wt: func() *WaveletTree {
				root := &Node{
					vector: []byte{0, 0, 1, 1, 0, 1, 1, 0, 1, 1, 0},
				}
				level1 := &Node{
					vector: []byte{1, 1, 1, 1, 0, 0},
					parent: root,
				}
				level2 := &Node{
					value:  's',
					parent: level1,
				}
				level1.right = level2
				root.right = level1
				wt := &WaveletTree{
					root: root,
					prefix: map[rune]Vector{
						's': []byte{1, 1},
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
