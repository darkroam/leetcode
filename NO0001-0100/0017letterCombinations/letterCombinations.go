package letterCombinations

//17. 电话号码的字母组合
//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//示例:
//	输入："23"
//	输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
//说明:
//	尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/string-to-integer-atoi
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

import "fmt"

func LetterCombinations(digits string) []string {
	key := [...]string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	if len(digits) == 0 {
		return []string{}
	}
	res := make([]string, 0)
	if len(digits) == 1 {
		for _, v := range key[int(digits[0])-50] {
			res = append(res, string(v))
		}
		return res
	} else {
		for _, v := range key[int(digits[0])-50] {
			res = append(res, string(v))
		}
	}

	for _, v := range digits[1:] {
		tmp := res
		res = make([]string, 0)
		for _, w := range key[int(v)-50] {
			for _, vv := range tmp {
				res = append(res, vv+string(w))
			}
		}
	}
	fmt.Println(res)
	return res
}

//space: O(?)
//time : O(?)
