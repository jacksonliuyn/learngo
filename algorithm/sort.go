package algorithm

func partition(nums []int, l, h int) int {
	i := l
	j := h

	v := nums[l]
	for {
		for nums[i] < v && i != h {
			i++
		}
		for v < nums[j] && j != l {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[j], nums[l] = nums[l], nums[j]
	return j

}
