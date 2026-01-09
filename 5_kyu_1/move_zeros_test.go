package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeros(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "",
			arr:  []int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1},
			want: []int{1, 2, 1, 1, 3, 1, 0, 0, 0, 0},
		},
		{
			name: "",
			arr:  []int{0, 0, 0, 0, 0, 0},
			want: []int{0, 0, 0, 0, 0, 0},
		},

		{
			name: "",
			arr:  []int{1, 0, 2, 0, 3, 0, 4, 0, 5},
			want: []int{1, 2, 3, 4, 5, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, MoveZeros(tt.arr))
		})
	}
}
