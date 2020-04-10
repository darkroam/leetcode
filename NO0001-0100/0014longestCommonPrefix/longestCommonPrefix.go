package longestCommonPrefix

//14. 最长公共前缀
//编写一个函数来查找字符串数组中的最长公共前缀。
//如果不存在公共前缀，返回空字符串 ""。
//示例 1:
//	输入: ["flower","flow","flight"]
//	输出: "fl"
//示例 2:
//	输入: ["dog","racecar","car"]
//	输出: ""
//	解释: 输入不存在公共前缀。
//说明:
//	所有输入只包含小写字母 a-z 。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/string-to-integer-atoi
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//import "fmt"

func LongestCommonPrefix(strs []string) string {
	var res string
	if len(strs) == 0 {
		return ""
	}
LOOP:
	for {
		if len(res) == len(strs[0]) {
			break LOOP
		}
		for i := 1; i < len(strs); i++ {
			if len(res) == len(strs[i]) {
				break LOOP
			}
			if strs[i][len(res)] != strs[0][len(res)] {
				break LOOP
			}
		}
		res = strs[0][:len(res)+1]
	}
	return res
}

//space: O(?)
//time : O(?)
