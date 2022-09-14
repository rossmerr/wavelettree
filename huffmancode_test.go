package wavelettree

import (
	"reflect"
	"testing"
)

func TestNewHuffmanCode(t *testing.T) {
	type args struct {
		value string
	}
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
					Key:       byte('C'),
					Frequency: 6,
				},
				Right: &HuffmanCode{
					Frequency: 9,
					Left: &HuffmanCode{
						Frequency: 4,
						Left: &HuffmanCode{
							Frequency: 1,
							Key:       byte('B'),
						},
						Right: &HuffmanCode{
							Frequency: 3,
							Key:       byte('D'),
						},
					},
					Right: &HuffmanCode{
						Frequency: 5,
						Key:       byte('A'),
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
