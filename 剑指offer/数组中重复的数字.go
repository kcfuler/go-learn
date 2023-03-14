
// 通过map简单实现
func findRepeatNumber(nums []int) int {
	hash := make(map[int]int)
	var res int

	for _, num := range nums {

		if _, ok := hash[num]; ok {
			res = num
			break
		}

		hash[num]++
	}

	return res
}

/*
题目说明尚未被充分使用，即 在一个长度为 n 的数组 nums 里的所有数字都在 0 ~ n-1 的范围内 。 此说明含义：数组元素的 索引 和 值 是 一对多 的关系。
因此，可遍历数组并通过交换操作，使元素的 索引 与 值 一一对应。因而，就能通过索引映射对应的值，起到与字典等价的作用。
*/
func findRepeatNumber(nums []int) int {
	for i, _ := range nums {
		for nums[i] != i { // 如果没有对应上， 就可能是重复或者位置不对
			if nums[nums[i]] == nums[i] { // 重复，返回答案
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i] // 位置不对，就把它放到对应的索引位置上
		}
	}
	return 1
}