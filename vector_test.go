package wavelettree

import "testing"

func TestVector_Rank(t *testing.T) {
	type args struct {
		i      byte
		offset int
	}
	tests := []struct {
		name string
		v    Vector
		args args
		want int
	}{
		{
			name: "mississippi",
			v: func() Vector {
				v, _, _, _ := NewVectorFromString("mississippi")
				return v
			}(),
			args: args{
				i:      0,
				offset: 4,
			},
			want: 3,
		},
		{
			name: "00110110110",
			v:    []byte{0, 0, 1, 1, 0, 1, 1, 0, 1, 1, 0},
			args: args{
				i:      0,
				offset: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Rank(tt.args.i, tt.args.offset); got != tt.want {
				t.Errorf("Vector.Rank() = %v, want %v", got, tt.want)
			}
		})
	}
}
