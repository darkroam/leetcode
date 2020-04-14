package fourSum

//18. 四数之和
//给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
//注意：
//答案中不可以包含重复的四元组。
//示例：
//	给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
//	满足要求的四元组集合为：
//	[
//	  [-1,  0, 0, 1],
//	  [-2, -1, 1, 2],
//	  [-2,  0, 0, 2]
//	]
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/string-to-integer-atoi
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//import "fmt"
import "sort"

func FourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 4 {
		return res
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	//fmt.Println(nums)

	for i := 0; i < len(nums)-3; i++ {
		tmpTarget1 := target - nums[i]
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			tmpTarget := tmpTarget1 - nums[j]
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for l, r := j+1, len(nums)-1; l < r; {
				//fmt.Println("i : ", nums[i], " j : ", nums[j], " l : ", nums[l], " r : ", nums[r])
				//fmt.Println("i : ", i, " j ", j, " l : ", l, " r : ", r)
				sum := nums[l] + nums[r]
				if sum == tmpTarget {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					//fmt.Println(res)
					//fmt.Println("i : ", nums[i], " l : ", nums[l], " r : ", nums[r])
					//fmt.Println("i : ", i, " l : ", l, " r : ", r)
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l, r = l+1, r-1
				} else if sum < tmpTarget {
					for nums[l] == nums[l+1] && l+1 < r { //此处两个判断条件交换，性能差距很大
						l++
					}
					l++
				} else {
					for nums[r] == nums[r-1] && l < r-1 { //此处两个判断条件交换，性能差距很大
						r--
					}
					r--
				}
			}
		}
	}
	//fmt.Println(len(res))
	return res
}

//space: O(?)
//time : O(?)
