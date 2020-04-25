package longestValidParentheses

//32. 最长有效括号
//给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
//示例 1:
//	输入: "(()"
//	输出: 2
//	解释: 最长有效括号子串为 "()"
//示例 2:
//	输入: ")()())"
//	输出: 4
//	解释: 最长有效括号子串为 "()()"
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/string-to-integer-atoi
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//import "fmt"

func LongestValidParentheses(s string) int {
	stack, mark, res, tmp := []int{}, make([]bool, len(s)), 0, 0
	for j, v := range s {
		if v == '(' {
			tmp++
			stack = append(stack, j)
		} else { //v == ')'
			if tmp > 0 { //tmp = len(stack)
				tmp--
				mark[stack[tmp]], mark[j] = true, true
				stack = stack[:tmp] //pop(stack)
				continue
			}
		}
	}
	tmp = 0
	for _, v := range mark { //最长连续true
		if v {
			tmp++
		} else {
			if tmp > res {
				res = tmp
			}
			tmp = 0
		}
	}
	if tmp > res {
		res = tmp
	}
	return res
}

//space: O(?)
//time : O(2N)
