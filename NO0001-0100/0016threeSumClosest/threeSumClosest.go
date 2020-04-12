package threeSumClosest

//16. 最接近的三数之和
//给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
//例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
//与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/string-to-integer-atoi
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//import "fmt"
import "sort"

func ThreeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var res = nums[0] + nums[1] + nums[2]
	//fmt.Println(nums)

	//fmt.Println("i : ", nums[0], " l : ", nums[1], " r : ", nums[2])
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for l, r := i+1, len(nums)-1; l < r; {
			sum := nums[i] + nums[l] + nums[r]
			//fmt.Println(res)
			//fmt.Println("i : ", nums[i], " l : ", nums[l], " r : ", nums[r])
			//fmt.Println("i : ", i, " l : ", l, " r : ", r)

			if sum == target {
				return target
			} else if res > target && sum > target {
				if res > sum {
					res = sum
				}
				r--
			} else if res > target && sum < target {
				if res-target > target-sum {
					res = sum
				}
				l++
			} else if res < target && sum > target {
				if target-res > sum-target {
					res = sum
				}
				r--
			} else if res < target && sum < target {
				if sum > res {
					res = sum
				}
				l++
			} else {
				return res
			}
		}
	}
	return res
}

//space: O(?)
//time : O(?)
