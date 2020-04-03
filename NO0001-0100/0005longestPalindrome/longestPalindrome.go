package longestPalindrome

//5. 最长回文子串
//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//示例 1：
//	输入: "babad"
//	输出: "bab"
//注意: "aba" 也是一个有效答案。
//示例 2：
//	输入: "cbbd"
//	输出: "bb"
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-palindromic-substring
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

import "fmt"

func LongestPalindrome(s string) string {
	var lenStr = len(s)
	if lenStr == 0 || lenStr == 1 {
		return s
	}
	var c1, c2 int
	maxStr, maxLen := 0, 1

	for c1, c2 = 0, 1; c2 < lenStr; {
		if s[c1] == s[c2] {
			lPtr, rPtr := c1-1, c2+1
			for lPtr >= 0 && rPtr < lenStr && s[lPtr] == s[rPtr] {
				lPtr--
				rPtr++
			}
			lPtr++
			rPtr--
			if rPtr+1-lPtr > maxLen {
				maxLen = rPtr - lPtr + 1
				maxStr = lPtr
			}

		}
		if c1+1 == c2 {
			c2++
		} else {
			c1++
		}
	}
	fmt.Printf("result : %s\n", s[maxStr:maxStr+maxLen])
	return s[maxStr : maxStr+maxLen]
}

//space: O(?)
//time : O(?)
