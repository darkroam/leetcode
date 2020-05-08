package nextPermutation

//31. 下一个排列
//实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
//如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
//必须原地修改，只允许使用额外常数空间。
//以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
//	1,2,3 → 1,3,2
//	3,2,1 → 1,2,3
//	1,1,5 → 1,5,1
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/string-to-integer-atoi
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//import "fmt"
//import "sort"

func NextPermutation(nums []int) {
	i := len(nums) - 2
	// [i+1:] 为降序排序
	for ; i >= 0 && nums[i] >= nums[i+1]; i-- {
	}
	if i >= 0 {
		// 从右端寻找第一个比[i]大的[k]，这个[k]肯定是序列里别i大中最小的一个，然后交换
		k := len(nums) - 1
		for ; nums[k] <= nums[i]; k-- {
		}
		nums[i], nums[k] = nums[k], nums[i]
	}

	//[i+1:]为降序排列，所以直接前后交换即可完成升序排列
	start, end := i+1, len(nums)-1
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

//space: O(?)
//time : O(?)
