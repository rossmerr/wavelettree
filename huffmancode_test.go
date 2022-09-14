package wavelettree

import (
	"reflect"
	"testing"
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
					Value:     'C',
					Frequency: 6,
				},
				Right: &HuffmanCode{
					Frequency: 9,
					Left: &HuffmanCode{
						Frequency: 4,
						Left: &HuffmanCode{
							Frequency: 1,
							Value:     'B',
						},
						Right: &HuffmanCode{
							Frequency: 3,
							Value:     'D',
						},
					},
					Right: &HuffmanCode{
						Frequency: 5,
						Value:     'A',
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
