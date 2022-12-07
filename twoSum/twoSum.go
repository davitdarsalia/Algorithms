package twoSum

func SumV1(data []int, target int) []int {
	var res []int
	for i, left := range data {
		for j, right := range data {
			if left+right == target && i != j {
				res = []int{j, i}
			}
		}
	}
	return res
}

func SumV2(data []int, target int) []int {
	hm := make(map[int]int)

	for i, v := range data {
		_, ok := hm[v]

		if ok {
			return []int{hm[v], i}
		}
		hm[target-v] = i
	}

	return nil
}
