package firstMissingPositive

//41. 缺失的第一个正数
//给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。
//示例 1:
//	输入: [1,2,0]
//	输出: 3
//示例 2:
//	输入: [3,4,-1,1]
//	输出: 2
//示例 3:
//	输入: [7,8,9,11,12]
//	输出: 1
//提示：
//	你的算法的时间复杂度应为O(n)，并且只能使用常数级别的额外空间。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/string-to-integer-atoi
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//import "fmt"

func FirstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for 0 < nums[i] && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

//space: O(?)
//time : O(?)
