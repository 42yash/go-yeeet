func findDuplicate(nums []int) int {
	prev := map[int]bool{}

	for _, num := range nums {
		val := prev[num]

		if val {
			return num
		}

		prev[num] = true
	}
	return -1
}
