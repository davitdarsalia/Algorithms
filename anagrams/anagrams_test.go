package anagrams

import (
	"testing"
)

func TestAnagrams(t *testing.T) {
	cases := []struct {
		input [2]string
		want  bool
	}{
		{
			input: [2]string{"brush", "shrub"},
			want:  true,
		},
		{
			input: [2]string{"nameless", "salesmen"},
			want:  true,
		},
		{
			input: [2]string{"god", "dog"},
			want:  true,
		},
	}

	for i, c := range cases {
		got := IsAnagram(c.input)
		if got != c.want {
			t.Errorf("CASE INDEX[%d] Got: %v, Want: %v", i, got, c.want)
		}
	}
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsAnagram([2]string{"nameless", "salesmen"})
	}
}
