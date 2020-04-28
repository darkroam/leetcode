package searchRange

//34. 在排序数组中查找元素的第一个和最后一个位置
//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//你的算法时间复杂度必须是 O(log n) 级别。
//如果数组中不存在目标值，返回 [-1, -1]。
//示例 1:
//	输入: nums = [5,7,7,8,8,10], target = 8
//	输出: [3,4]
//示例 2:
//	输入: nums = [5,7,7,8,8,10], target = 6
//	输出: [-1,-1]
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/string-to-integer-atoi
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//import "fmt"

func SearchRange(nums []int, target int) []int {
	length := len(nums) - 1
	for l, r := 0, length; l <= r; {
		half := (l + r) / 2
		if nums[half] == target {
			tmp := half
			for l <= half {
				mid := (l + half) / 2
				if nums[mid] == target {
					if mid == 0 || nums[mid-1] != target {
						l = mid
						break
					} else {
						half = mid - 1
					}
				} else if target < nums[mid] {
					half = mid - 1
				} else {
					l = mid + 1
				}
			}
			half = tmp
			for half <= r {
				mid := (half + r) / 2
				if nums[mid] == target {
					if mid == length || nums[mid+1] != target {
						r = mid
						break
					} else {
						half = mid + 1
					}
				} else if nums[mid] < target {
					half = mid + 1
				} else {
					r = mid - 1
				}
			}
			return []int{l, r}
		} else if target < nums[half] {
			r = half - 1
		} else {
			l = half + 1
		}
	}
	return []int{-1, -1}
}

//space: O(?)
//time : O(?)
