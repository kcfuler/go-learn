package main

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
	rv = l

	res := lv - rv
	return res
}
