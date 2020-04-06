package isMatch

//10. 正则表达式匹配
//给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
//	'.' 匹配任意单个字符
//	'*' 匹配零个或多个前面的那一个元素
//所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
//说明:
//	s 可能为空，且只包含从 a-z 的小写字母。
//	p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
//示例 1:
//	输入:
//		s = "aa"
//		p = "a"
//	输出: false
//解释: "a" 无法匹配 "aa" 整个字符串。
//示例 2:
//	输入:
//		s = "aa"
//		p = "a*"
//	输出: true
//解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
//示例 3:
//	输入:
//		s = "ab"
//		p = ".*"
//	输出: true
//解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
//示例 4:
//	输入:
//		s = "aab"
//		p = "c*a*b"
//	输出: true
//解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
//示例 5:
//	输入:
//		s = "mississippi"
//		p = "mis*is*p*."
//	输出: false

//import "fmt"

//貌似只有递归或者动态规划方法能解,需要确认匹配的情况存在，存在可能不为一
//从左右往右和从右往做都存在问题

func IsMatch(s string, p string) bool {

	if len(p) == 0 && len(s) == 0 {
		return true
	}

	if len(p) == 0 && len(s) != 0 {
		return false
	}

	//if p[0] == '*' { //如果需要考虑*前必须有符号
	//return false
	//}

	if len(s) == 0 {
		if len(p)%2 == 1 {
			return false
		} else if p[1] == '*' {
			return IsMatch(s, p[2:])
		} else {
			return false
		}
	}

	if len(p) >= 2 && p[1] == '*' {
		if p[0] == s[0] || p[0] == '.' {
			return IsMatch(s[1:], p) || IsMatch(s[1:], p[2:]) || IsMatch(s, p[2:])
		}
		if p[0] != s[0] {
			return IsMatch(s, p[2:])
		}
	}
	if p[0] == s[0] || p[0] == '.' {
		return IsMatch(s[1:], p[1:])
	}
	return false

	//partens := make([]string, 0, 10)
	//for i := 0; i < len(p); i++ {
	//if p[i] != '*' {
	//partens = append(partens, string(p[i]))
	//} else {
	//partens[len(partens)-1] = partens[len(partens)-1] + "*"
	//}
	//}
	//fmt.Println("partens = ", partens)

	//----------------------------------
	//i, j := len(s)-1, len(p)-1

	//for i >= 0 && j >= 0 {
	//fmt.Printf("s[%2d] = %c , p[%2d] = %c\n", i, s[i], j, p[j])
	//if p[j] == s[i] || p[j] == '.' { //处理字符相同和.占位符号的情况
	//i--
	//j--
	//continue
	//}
	//if p[j] == '*' {
	//if j-1 == 0 && p[j-1] == '.' {
	//return true
	//}
	////if j-1 != 0 && p[j-1] == '.' {
	////if j-1
	////}
	//if p[j-1] != s[i] { //处理不存在a.的情况
	//j -= 2
	//continue
	//} else if p[j-1] == s[i] { //处理匹配成功的情况
	//i--
	//continue
	//}
	//}
	//break
	//}
	//for j >= 0 && p[j] == '*' {
	//if p[j] == '*' {
	//j -= 2
	//}
	//}
	////fmt.Printf("i = %d , j = %d\n", i, j)
	//if i == -1 && j == -1 {
	//return true
	//}
	//return false

}

//func isMatch(s string, p string) bool {
//	const (
//		T_MATCH_ANY = '*'
//		T_MATCH_ONE = '.'
//	)
//	var (
//		match  func(si, pi int) bool
//		sLen   = len(s)
//		pLen   = len(p)
//		record = make([][]int, sLen+1) // improve cache speed
//	)
//	for i := 0; i < sLen+1; i++ {
//		record[i] = make([]int, pLen+1)
//	}
//
//	match = func(si, pi int) bool {
//		if si == sLen && pi == pLen {
//			return true
//		}
//		if pi >= pLen {
//			return false
//		}
//
//		if record[si][pi] == 1 {
//			return false
//		}
//		record[si][pi] = 1
//
//		switch p[pi] {
//		case T_MATCH_ONE:
//			if pi == pLen-1 && si == sLen-1 {
//				return true
//			}
//			if pi+1 < pLen && p[pi+1] == T_MATCH_ANY {
//				return match(si, pi+1)
//			}
//			if si >= sLen {
//				return false
//			}
//			return match(si+1, pi+1)
//		case T_MATCH_ANY:
//			if si >= sLen {
//				return match(si, pi+1)
//			}
//			switch p[pi-1] {
//			case T_MATCH_ONE:
//				if r := match(si, pi+1); r { // fix test_4
//					return true
//				}
//				for si < sLen {
//					si += 1 // fix test_5
//					if r := match(si, pi+1); r {
//						return true
//					}
//				}
//				return false
//			default:
//				if r := match(si, pi+1); r {
//					return true
//				}
//				for si < sLen && s[si] == p[pi-1] {
//					si += 1
//					if r := match(si, pi+1); r {
//						return true
//					}
//				}
//				return false
//			}
//		default:
//			if pi+1 < pLen && p[pi+1] == T_MATCH_ANY {
//				return match(si, pi+1)
//			}
//			if si >= sLen || p[pi] != s[si] {
//				return false
//			}
//			return match(si+1, pi+1)
//		}
//	}
//	return match(0, 0)
//}

//space: O(?)
//time : O(?)
