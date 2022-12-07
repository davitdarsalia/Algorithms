package palindromes

import "testing"

func TestPalindromes(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{
			input: "mom",
			want:  true,
		},
		{
			input: "noon",
			want:  true,
		},
		{
			input: "civic",
			want:  true,
		},
		{
			input: "rogor",
			want:  true,
		},
		{
			input: "roger",
			want:  false,
		},
		{
			input: "racecar",
			want:  true,
		},
		{
			input: "god",
			want:  false,
		},
	}

	for i, c := range cases {
		got := IsPalindrome(c.input)
		if got != c.want {
			t.Errorf("CASE INDEX[%d] Got: %v, Want: %v", i, got, c.want)
		}
	}
}
