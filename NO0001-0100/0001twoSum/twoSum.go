package twosum

//1. 两数之和
//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
//示例:
//	给定 nums = [2, 7, 11, 15], target = 9
//	因为 nums[0] + nums[1] = 2 + 7 = 9
//	所以返回 [0, 1]
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/two-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func TwoSum(nums []int, target int) []int {
	for i, num1 := range nums[:len(nums)-1] {
		for j, num2 := range nums[i+1:] {
			if num1+num2 == target {
				return []int{i, i + 1 + j}
			}
		}
	}
	return []int{}
}

//space: O(1)
//time : O(n^2)
