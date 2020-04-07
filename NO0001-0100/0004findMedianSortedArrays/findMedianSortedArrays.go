package findMedianSortedArrays

//4. 寻找两个有序数组的中位数
//给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
//请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//你可以假设 nums1 和 nums2 不会同时为空。
//示例 1:
//	nums1 = [1, 3]
//	nums2 = [2]
//	则中位数是 2.0
//示例 2:
//	nums1 = [1, 2]
//	nums2 = [3, 4]
//	则中位数是 (2 + 3)/2 = 2.5
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

import "fmt"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var len1 = len(nums1)
	var len2 = len(nums2)
	var midVar1, midVar2 int
	var mid1 int = (len1 + len2 + 1) / 2
	var mid2 int = (len1 + len2 + 2) / 2
	var l1, l2 int

	for i := 0; i <= mid1; i++ {
		midVar1 = midVar2
		if l1 == len1 && l2 == len2 {
			continue
		}
		if l1 == len1 {
			midVar2 = nums2[l2]
			l2++
			continue
		}
		if l2 == len2 {
			midVar2 = nums1[l1]
			l1++
			continue
		}
		if nums1[l1] < nums2[l2] {
			midVar2 = nums1[l1]
			l1++
		} else {
			midVar2 = nums2[l2]
			l2++
		}
	}

	fmt.Println("nums1=", midVar1, "  nums2=", midVar2)
	if mid1 == mid2 {
		return float64(midVar1)
	} else {
		return (float64(midVar1) + float64(midVar2)) / 2
	}
}

//space: O(?)
//time : O(?)

//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//m,n:= len(nums1), len(nums2)
//if (m>n) {
//nums1,nums2,m,n=nums2,nums1,n,m
//}
//i,j,min,max,half,maxl,maxr:=0,0,0,m,(m+n+1)/2,0,0
//for ;min<=max; {
//i=(min+max)/2
//j=half-i
//if (i<max && nums2[j-1]>nums1[i]) {
//min=i+1
//} else if (i>min && nums1[i-1]>nums2[j]) {
//max=i-1
//} else {
//maxl,maxr=0,0
//if (i==0) {
//maxl=nums2[j-1]
//} else if (j==0) {
//maxl=nums1[i-1]
//} else {
//maxl = func_max(nums1[i-1], nums2[j-1])
//}
//if (m+n)%2==1 {
//return float64(maxl)
//}

//if (i==m) {
//maxr = nums2[j];
//} else if (j==n) {
//maxr = nums1[i];
//} else {
//maxr = func_min(nums2[j],nums1[i])
//}

//return float64(maxl+maxr)/2.0
//}
//}
//return 0.0
//}

//func func_max(a,b int) int {
//if (a<b) {
//return b;
//}
//return a;
//}

//func func_min(a,b int) int {
//if (a>b) {
//return b;
//}
//return a;
//}
