package convert

//将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
//比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
//	L   C   I   R
//	E T O E S I I G
//	E   D   H   N
//之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
//请你实现这个将字符串进行指定行数变换的函数：
//string convert(string s, int numRows);
//示例 1:
//	输入: s = "LEETCODEISHIRING", numRows = 3
//	输出: "LCIRETOESIIGEDHN"
//示例 2:
//	输入: s = "LEETCODEISHIRING", numRows = 4
//	输出: "LDREOEIIECIHNTSG"
//	解释:
//		L     D     R
//		E   O E   I I
//		E C   I H   N
//		T     S     G
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/zigzag-conversion
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

import "fmt"

//1       2n+1
//2  n+n  2n+2  ... 2n+N
//3  .    2n+3
//4  .    2n+4
//.  .    .
//.  n+4  .
//.  n+3  .
//n  n+2  3n
//   n+1
//   1, 2n+1, ..., 2n+N, 2, 2n, 2n+2, ..., 2n+N, 3, 2n-+1...
//  可以先排列N*n个
//一组打印 1, 2n+1, ... 2n+N, 2, 2n+2..
//另一组打印 n+n, 4n+n...倒排
//补充一行，再交叉打印

func Convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	lenStr := len(s)
	N := lenStr / (numRows - 1)

	ptr, ptrRes := 0, 0
	res := make([]byte, numRows*(N+1))
	resString := make([]byte, lenStr)
LOOP:
	for i := 0; i <= N; i++ {
		for j := 0; j < numRows; j++ {
			if i%2 == 0 {
				if j == numRows-1 {
					fmt.Printf("  ")
					res[ptrRes] = '\x00'
					ptrRes++
				} else if (numRows-1)*i+j < lenStr {
					fmt.Printf(" %c", s[(numRows-1)*i+j])
					res[ptrRes] = s[(numRows-1)*i+j]
					ptrRes++
					ptr++
				}
			} else {
				if j == 0 {
					fmt.Printf("  ")
					res[ptrRes] = '\x00'
					ptrRes++
				} else if (numRows-1)*(i+1)-j < lenStr {
					fmt.Printf(" %c", s[(numRows-1)*(i+1)-j])
					res[ptrRes] = s[(numRows-1)*(i+1)-j]
					ptrRes++
					ptr++
				} else {
					fmt.Printf("  ")
					res[ptrRes] = '\x00'
					ptrRes++
				}

			}
			if ptr == lenStr {
				fmt.Println()
				break LOOP
			}
		}
		fmt.Println()
	}

	ptr = 0
	fmt.Println("---------------")
	for i := 0; i < numRows; i++ {
		for j := 0; j <= N; j++ {
			if res[j*numRows+i] != '\x00' {
				fmt.Printf(" %c", res[(numRows)*j+i])
				resString[ptr] = res[(numRows)*j+i]
				ptr++
			}
		}
		fmt.Println("|")
	}
	fmt.Println(" res: ", string(resString))
	fmt.Println()
	return string(resString)
}

//执行用时为 0 ms 的范例
//
//func convert(s string, numRows int) string {
//	if numRows == 1 {
//		return s
//	}
//
//	var result strings.Builder
//	n := len(s)
//
//	add := 2*numRows - 2
//	for i := 0; i < n; i += add {
//		result.WriteByte(s[i])
//	}
//
//	add = 2*numRows - 4
//	for row := 1; row < numRows/2; row++ {
//		for i := row; i < n; i = i + add + 2*row {
//			result.WriteByte(s[i])
//			if i+add < n {
//				result.WriteByte(s[i+add])
//			}
//		}
//		add -= 2
//	}
//
//	if numRows%2 == 1 {
//		add = numRows - 1
//		for i := numRows / 2; i < n; i += add {
//			result.WriteByte(s[i])
//		}
//	}
//
//	if numRows%2 == 0 {
//		add = numRows - 2
//	} else {
//		add = numRows - 3
//	}
//	for row := (numRows + 1) / 2; row < numRows-1; row++ {
//		for i := row; i < n; i = i + add + 2*row {
//			result.WriteByte(s[i])
//			if i+add < n {
//				result.WriteByte(s[i+add])
//			}
//		}
//		add -= 2
//	}
//
//	add = 2*numRows - 2
//	for i := numRows - 1; i < n; i += add {
//		result.WriteByte(s[i])
//	}
//
//	return result.String()
//}

//space: O(?)
//time : O(?)
