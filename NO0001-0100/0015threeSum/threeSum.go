package threeSum

//15. 三数之和
//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
//注意：答案中不可以包含重复的三元组。
//示例：
//	给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//	满足要求的三元组集合为：
//	[
//	  [-1, 0, 1],
//	  [-1, -1, 2]
//	]
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/string-to-integer-atoi
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//import "fmt"
import "sort"

func ThreeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 3 {
		return res
	}
	//for i := 0; i < len(nums)-1; i++ {
	//for j := 0; j < len(nums)-1-i; j++ {
	//if nums[j] > nums[j+1] {
	//nums[j+1], nums[j] = nums[j], nums[j+1]
	//}
	//}
	//}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	//fmt.Println(nums)

	var target int
	for i := 0; i < len(nums)-2 && nums[i] <= target; i++ {
		//if nums[i] > 0 { //本件测加入到循环判断中
		//return res
		//}
		tmpTarget := target - nums[i]
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for l, r := i+1, len(nums)-1; l < r; {
			sum := nums[l] + nums[r]
			if sum == tmpTarget {
				res = append(res, []int{nums[i], nums[l], nums[r]})
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
	return res

}

//space: O(?)
//time : O(?)
