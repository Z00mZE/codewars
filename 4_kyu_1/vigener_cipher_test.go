package __kyu_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVigenerCipher_Encode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		key   string
		alpha string
		input string
		want  string
	}{
		{
			name:  "lxfopvefrnhr",
			key:   "lemon",
			alpha: "abcdefghijklmnopqrstuvwxyz",
			input: "attackatdawn",
			want:  "lxfopvefrnhr",
		},
		{
			name:  "rovwsoiv",
			key:   "password",
			alpha: "abcdefghijklmnopqrstuvwxyz",
			input: "codewars",
			want:  "rovwsoiv",
		},
		{
			name:  "CODEWARS",
			key:   "password",
			alpha: "abcdefghijklmnopqrstuvwxyz",
			input: "CODEWARS",
			want:  "CODEWARS",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := VigenèreCipher{
				Key:   tt.key,
				Alpha: tt.alpha,
			}
			assert.Equal(t, tt.want, c.Encode(tt.input))
		})
	}
}

func TestVigenerCipher_Decode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		key   string
		alpha string
		input string
		want  string
	}{
		{
			name:  "lxfopvefrnhr",
			key:   "lemon",
			alpha: "abcdefghijklmnopqrstuvwxyz",
			input: "attackatdawn",
			want:  "lxfopvefrnhr",
		},
		{
			name:  "rovwsoiv",
			key:   "password",
			alpha: "abcdefghijklmnopqrstuvwxyz",
			input: "codewars",
			want:  "rovwsoiv",
		},
		{
			name:  "CODEWARS",
			key:   "password",
			alpha: "abcdefghijklmnopqrstuvwxyz",
			input: "CODEWARS",
			want:  "CODEWARS",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := VigenèreCipher{
				Key:   tt.key,
				Alpha: tt.alpha,
			}
			assert.Equal(t, tt.input, c.Decode(tt.want))
		})
	}
}
