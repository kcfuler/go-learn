func search(nums []int, target int) int {
	lv, rv := 0, 0

	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) >> 1 // 求下界
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	lv = l

	l, r = 0, len(nums)-1
	for l < r {
		mid := (l + r + 1) >> 1 // 求上界
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	rv = r

	var res int

	if rv == lv && nums[lv] != target { // 当查找到同一个数时，需要判断它是不是我们需要的
		res = 0
	} else {
		res = rv - lv + 1
	}

	return res
}