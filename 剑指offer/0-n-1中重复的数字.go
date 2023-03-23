func missingNumber(nums []int) int {

	l, r := 0, len(nums)-1

	// 这里使用的是另外一种二分的写法，acwing的模板没调出来
	for l <= r {
		mid := (l + r) >> 1
		if mid == nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l
}