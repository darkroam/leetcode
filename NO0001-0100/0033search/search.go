package search

//33. 搜索旋转排序数组
//假设按照升序排序的数组在预先未知的某个点上进行了旋转。
//( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
//搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
//你可以假设数组中不存在重复的元素。
//你的算法时间复杂度必须是 O(log n) 级别。
//示例 1:
//	输入: nums = [4,5,6,7,0,1,2], target = 0
//	输出: 4
//示例 2:
//	输入: nums = [4,5,6,7,0,1,2], target = 3
//	输出: -1
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/string-to-integer-atoi
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//import "fmt"

func Search(nums []int, target int) int {
	for l, r := 0, len(nums)-1; l <= r; {
		half := (l + r) / 2
		if nums[half] == target {
			return half
		} else if nums[l] < nums[half] {
			if nums[l] < target && target < nums[half] {
				r = half - 1
			} else if nums[l] == target {
				return l
			} else {
				l = half + 1
			}
		} else {
			if nums[half] < target && target < nums[r] {
				l = half + 1
			} else if nums[r] == target {
				return r
			} else {
				r = half - 1
			}
		}
	}
	return -1
}

//space: O(?)
//time : O(?)
