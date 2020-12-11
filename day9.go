package aoc2020

func Day9Part1() int {
	nums := MapInt(ReadLines("day9.input.txt"))

	buffer := 25
	for i := buffer; i <= len(nums); i++ {
		found := false
		for j := 0; j < buffer; j++ {
			for k := 0; k < buffer; k++ {
				if i == k {
					continue
				}
				if nums[i] == nums[i-j-1]+nums[i-k-1] {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			return nums[i]
		}
	}

	return -1
}

func Day9Part2() int {
	return -1
}
