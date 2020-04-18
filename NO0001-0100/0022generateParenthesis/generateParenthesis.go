package generateParenthesis

//22. 括号生成
//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//示例：
//	输入：n = 3
//	输出：[
//		"((()))",
//		"(()())",
//		"(())()",
//		"()(())",
//		"()()()"
//	      ]
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/string-to-integer-atoi
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//import "fmt"

func GenerateParenthesis(n int) []string {
	ans := make([]string, 0)
	var gen func(l, r int, s string)
	gen = func(l, r int, s string) {
		if l == n && r == n {
			ans = append(ans, s)
			return
		}
		if l < n {
			gen(l+1, r, s+"(")
		}
		if r < l {
			gen(l, r+1, s+")")
		}
		//fmt.Print("l ", l, " r ", r, " s ", s, "\n")

	}
	gen(0, 0, "")
	return ans

}

//space: O(?)
//time : O(?)
