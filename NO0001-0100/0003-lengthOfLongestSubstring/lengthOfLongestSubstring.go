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
	//	var maxLength int
	//	var index [128]int
	//	var l, r int
	//	for ; r < len(s); r++ {
	//		c := int(s[r])
	//		if index[c] > 0 && index[c] > l {
	//			if length := r - l; length > maxLength {
	//				maxLength = length
	//			}
	//			l = index[c]
	//		}
	//		index[c] = r + 1
	//	}
	//	if length := r - l; length > maxLength {
	//		maxLength = length
	//	}
	//	return maxLength
	//}
	//
	//func a(s string) int {

	var maxStr, maxCount int
	var tmpStr, tmpCount int
	postion := make(map[byte]int, 26)

	for i := 0; i < len(s); i++ {
		tmpIdx, ok := postion[s[i]]
		//三种情况++
		//字典中没有，重新回到了重复的字符位置，重现的字符不在字符串中
		if !ok || tmpIdx == i || tmpIdx < tmpStr {
			postion[s[i]] = i
			tmpCount++
			continue
		}

		if tmpCount > maxCount {
			maxCount = tmpCount
			maxStr = tmpStr
		}
		postion[s[i]] = i   //更新字典最新出现的位置
		tmpStr = tmpIdx + 1 //重置初始字符串位置为重现重现过的字符下一个位置
		tmpCount = 1        //字符串重新值1,新字符串肯定是从下一个开始
		i = tmpIdx + 1      //i重置为字符上一次出现的位置下一个位置，又因为下一个肯定没问题，且循环结束i++，所以实际是从字符冲第二个开始重新开始检查
	}

	if tmpCount > maxCount {
		maxCount = tmpCount
		maxStr = tmpStr
	}

	fmt.Println("result: ", s[maxStr:maxStr+maxCount])
	return maxCount
}

//space: O(1)
//time : O(n^2)
