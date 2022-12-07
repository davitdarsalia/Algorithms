package subsequence

import "testing"

func TestSubsequence(t *testing.T) {
	cases := []struct {
		sequence string
		str      string
		want     bool
	}{
		{
			sequence: "abc",
			str:      "ah1b81c",
			want:     true,
		},
		{
			sequence: "gcp",
			str:      "googlecloudplatform",
			want:     true,
		},
		{
			sequence: "cde",
			str:      "c19d81e",
			want:     true,
		},
		{
			sequence: "ksqucoqp",
			str:      "kadskqcuc1o021q041pp",
			want:     true,
		},
		{
			sequence: "akc",
			str:      "c_k-0",
			want:     false,
		},
		{
			sequence: "abcs",
			str:      "2",
			want:     false,
		},
	}

	for i, c := range cases {
		got := ValidSubsequence(c.sequence, c.str)
		if got != c.want {
			t.Errorf("CASE INDEX[%d] Got: %v, Want: %v", i, got, c.want)
		}
	}
}
