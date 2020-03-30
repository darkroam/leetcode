package lengthOfLongestSubstring

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//示例 1:
//	输入: "abcabcbb"
//	输出: 3
//	解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//	输入: "bbbbb"
//	输出: 1
//	解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//	输入: "pwwkew"
//	输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

import "fmt"

//import "strings"

func LengthOfLongestSubstring(s string) int {
	//tmpCount(当前扫描的字符串长度) 可以通过i和tmpStr(字符串起始)计算出来
	//tmpStr 可以替换为l
	//i 可以替换为r
	//index[c] == 0 为空
	//index[c] 记录字符位置偏移量+1,以简化是否存在的判断

	//三个优势：
	//1.没有用map而是直接用[]int，用字符ascii码字索引
	//2.index值0 为空，记录的值为实际字符编译量+1
	//3.当发现有重复字符时，l向右边移到该字符下一个，然后继续检查
	var maxStr, maxCount int
	var l, r int
	var index [128]int

	for ; r < len(s); r++ {
		c := int(s[r])

		//完全等价于[index[c]==0(第一次记录入表格)||index[c]-1< l(该字符在当前字符之外)]取反
		//index[c]>0(字符已在表内)&& index[c]>l(该字符在当前字符串内部)
		if index[c] > 0 && index[c] > l {
			if r-l > maxCount {
				maxCount = r - l
				maxStr = l
			}
			l = index[c]
			//r = l 不需要重新从l开始计算
		}
		index[c] = r + 1 //不论什么情况下都需要重置index
	}

	if r-l > maxCount {
		maxCount = r - l
		maxStr = l
	}

	fmt.Println("result: ", s[maxStr:maxStr+maxCount])
	return maxCount
}

//space: O(1)
//time : O(n)
