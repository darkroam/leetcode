package maxArea

//11. 盛最多水的容器
//给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//说明：你不能倾斜容器，且 n 的值至少为 2。
//图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
//示例：
//	输入：[1,8,6,2,5,4,8,3,7]
//	输出：49
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/string-to-integer-atoi
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//import "fmt"

//func MaxArea(height []int) int {
//var maxArea int
//for i := 0; i < len(height)-1; i++ {
//for j := i + 1; j < len(height); j++ {
//tmpHight := 0
//if height[j] > height[i] {
//tmpHight = height[i]
//} else {
//tmpHight = height[j]
//}
//if maxArea < (j-i)*tmpHight {
//maxArea = (j - i) * tmpHight
//}
//}
//}
//return maxArea
//}

func MaxArea(height []int) int {
	var maxArea int
	var l, r, ptr = 0, len(height) - 1, 0
	//fmt.Println(" l = ", l, " ptr = ", ptr, " r = ", r)
	if height[l] > height[r] {
		maxArea = (r - l) * height[r]
	} else {
		maxArea = (r - l) * height[l]
	}
	for l != r {
		if height[l] > height[r] {
			ptr = r - 1
			for height[ptr] <= height[r] && ptr != l {
				ptr--
			}
			r = ptr
			tmpHight := 0
			//fmt.Println(" l = ", l, " ptr = ", ptr, " r = ", r)
			if height[l] > height[r] {
				tmpHight = height[r]
			} else {
				tmpHight = height[l]
			}
			if maxArea < (r-l)*tmpHight {
				maxArea = (r - l) * tmpHight
			}
		} else {
			ptr = l + 1
			for height[ptr] <= height[l] && ptr != r {
				ptr++
			}
			l = ptr
			tmpHight := 0
			//fmt.Println(" l = ", l, " ptr = ", ptr, " r = ", r)
			if height[l] > height[r] {
				tmpHight = height[r]
			} else {
				tmpHight = height[l]
			}
			if maxArea < (r-l)*tmpHight {
				maxArea = (r - l) * tmpHight
			}
		}
	}
	return maxArea
}

//space: O(?)
//time : O(?)
