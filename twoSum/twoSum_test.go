package twoSum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	t.Run("Brute Approach", func(t *testing.T) {
		cases := []struct {
			input  []int
			target int
			want   []int
		}{
			{
				input:  []int{2, 13, 15, 6, 51},
				target: 53,
				want:   []int{0, 4},
			},
			{
				input:  []int{2, 7, 11, 15},
				target: 9,
				want:   []int{0, 1},
			},
			{
				input:  []int{5, 3, 1, 6, 4, 91, 21, 53, 76, 13, 76, 120},
				target: 211,
				want:   []int{5, 11},
			},
		}

		for i, c := range cases {
			got := SumV1(c.input, c.target)

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("CASE INDEX[%d] Got: %v, Want: %v", i, got, c.want)
			}
		}
	})

	t.Run("Optimal Approach", func(t *testing.T) {
		cases := []struct {
			input  []int
			target int
			want   []int
		}{
			{
				input:  []int{2, 13, 15, 6, 51},
				target: 53,
				want:   []int{0, 4},
			},
			{
				input:  []int{2, 7, 11, 15},
				target: 9,
				want:   []int{0, 1},
			},
			{
				input:  []int{5, 3, 1, 6, 4, 91, 21, 53, 76, 13, 76, 120},
				target: 211,
				want:   []int{5, 11},
			},
		}

		for i, c := range cases {
			got := SumV2(c.input, c.target)

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("CASE INDEX[%d] Got: %v, Want: %v", i, got, c.want)
			}
		}
	})
}
