package multiply

//43. 字符串相乘
//给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
//示例 1:
//	输入: num1 = "2", num2 = "3"
//	输出: "6"
//示例 2:
//	输入: num1 = "123", num2 = "456"
//	输出: "56088"
//说明：
//	num1 和 num2 的长度小于110。
//	num1 和 num2 只包含数字 0-9。
//	num1 和 num2 均不以零开头，除非是数字 0 本身。
//	不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/string-to-integer-atoi
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//import "fmt"

func Multiply4(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	resStr := make([]byte, len(num1)+len(num2))
	for n2 := len(num2) - 1; n2 >= 0; n2-- {
		for n1 := len(num1) - 1; n1 >= 0; n1-- {
			tmp := resStr[n2+n1+1] + (num2[n2]-48)*(num1[n1]-48)
			resStr[n2+n1+1] = tmp % 10
			resStr[n2+n1] += tmp / 10
		}
	}
	for resStr[0] == 0 {
		resStr = resStr[1:]
	}
	for i := 0; i < len(resStr); i++ {
		resStr[i] += 48
	}
	return string(resStr)
}

func Multiply3(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	l1, l2 := len(num1), len(num2)
	res := make([]byte, l1+l2)
	l1--
	l2--
	for ; l1 >= 0; l1-- {
		n1 := num1[l1] - '0'
		for j := l2; j >= 0; j-- {
			sum := res[l1+j+1] + n1*(num2[j]-'0')
			res[l1+j+1] = sum % 10
			res[l1+j] += sum / 10
		}
	}
	l1 = 0
	if res[l1] == 0 {
		l1++
	}
	for l2 := l1; l2 < len(res); l2++ {
		res[l2] += '0'
	}
	return string(res[l1:])
}

func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n, resStr, length, count1, count2 := 9, "", len(num1)+len(num2), 0, 0

	for n2 := len(num2) - n; n2+n > 0; n2 -= n {
		count1 = 0
		for n1 := len(num1) - n; n1+n > 0; n1 -= n {
			if n1 >= 0 && n2 >= 0 {
				tmp := MyAtoi(num2[n2:n2+n]) * MyAtoi(num1[n1:n1+n])
				if len(resStr) <= length-n1-n2-n*2 {
					for i, loop := 0, length-n1-n2-n*2-len(resStr); i < loop; i++ {
						resStr = "0" + resStr
					}
					resStr = MyItoa(tmp) + resStr
				} else {
					resStr = addTwo(resStr[:len(resStr)-length+n1+n2+n*2], MyItoa(tmp)) + resStr[len(resStr)-length+n1+n2+n*2:]
				}
			} else if n1+n >= 0 && n2 >= 0 {
				tmp := MyAtoi(num2[n2:n2+n]) * MyAtoi(num1[:n1+n])
				if len(resStr) < count1+count2 {
					for i, loop := 0, count1+count2-len(resStr); i < loop; i++ {
						resStr = "0" + resStr
					}
					resStr = MyItoa(tmp) + resStr
				} else {
					resStr = addTwo(resStr[:len(resStr)-count1-count2], MyItoa(tmp)) + resStr[len(resStr)-count1-count2:]
				}
			} else if n1 >= 0 && n2+n >= 0 {
				tmp := MyAtoi(num2[:n2+n]) * MyAtoi(num1[n1:n1+n])
				if len(resStr) < count1+count2 {
					for i, loop := 0, count1+count2-len(resStr); i < loop; i++ {
						resStr = "0" + resStr
					}
					resStr = MyItoa(tmp) + resStr
				} else {
					resStr = addTwo(resStr[:len(resStr)-count1-count2], MyItoa(tmp)) + resStr[len(resStr)-count1-count2:]
				}
			} else if n1+n >= 0 && n2+n >= 0 {
				tmp := MyAtoi(num2[:n2+n]) * MyAtoi(num1[:n1+n])
				resStr = addTwo(resStr[:len(resStr)-count1-count2], MyItoa(tmp)) + resStr[len(resStr)-count1-count2:]
			}
			count1 += n
		}
		count2 += n
	}
	return resStr
}

func addTwo(num1 string, num2 string) (res string) {
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}
	n, carry := 18, false
	for n2, n1 := len(num2)-n, len(num1)-n; n2+n > 0 || n1+n > 0; {
		val, tmp := 0, ""
		if carry {
			val++
			carry = false
		}
		if n1 >= 0 && n2 >= 0 {
			val += MyAtoi(num1[n1:n1+n]) + MyAtoi(num2[n2:n2+n])
			tmp = MyItoa(val)
			if len(tmp) < n {
				loop := n - len(tmp)
				for i := 0; i < loop; i++ {
					tmp = "0" + tmp
				}
			}
		} else if n1+n >= 0 && n2 >= 0 {
			val += MyAtoi(num1[:n1+n]) + MyAtoi(num2[n2:n2+n])
			tmp = MyItoa(val)
			if len(tmp) < n {
				for i, loop := 0, n-len(tmp); i < loop; i++ {
					tmp = "0" + tmp
				}
			}
		} else if n1+n < 0 && n2 >= 0 {
			val += MyAtoi(num2[n2 : n2+n])
			tmp = MyItoa(val)
			if len(tmp) < n {
				for i, loop := 0, n-len(tmp); i < loop; i++ {
					tmp = "0" + tmp
				}
			}
			//} else if n1 >= 0 && n2+n >= 0 {
			//val += MyAtoi(num1[n1:n1+n]) + MyAtoi(num2[:n2+n])
			//tmp = MyItoa(val)
			//if len(tmp) < n {
			//for i, loop := 0, n-len(tmp); i < loop; i++ {
			//tmp = "0" + tmp
			//}
			//}
			//} else if n1 >= 0 && n2+n < 0 {
			//val += MyAtoi(num1[n1 : n1+n])
			//tmp = MyItoa(val)
			//if len(tmp) < n {
			//for i, loop := 0, n-len(tmp); i < loop; i++ {
			//tmp = "0" + tmp
			//}
			//}
		} else if n1+n >= 0 && n2+n >= 0 {
			val += MyAtoi(num1[:n1+n]) + MyAtoi(num2[:n2+n])
			tmp = MyItoa(val)
		} else if n1+n < 0 && n2+n >= 0 {
			val += MyAtoi(num2[:n2+n])
			tmp = MyItoa(val)
			//} else if n1+n >= 0 && n2+n < 0 {
			//val += MyAtoi(num1[:n1+n])
			//tmp = MyItoa(val)
		}
		if len(tmp) > n {
			carry = true
			tmp = tmp[1:]
		}
		res, n1, n2 = tmp+res, n1-n, n2-n
	}
	if carry {
		res = "1" + res
	}
	return res
}

func MyAtoi(str string) (res int) {
	for _, v := range str {
		res = res<<1 + res<<3 + int(v) - 48
	}
	return
}

func MyItoa(n int) (res string) {
	for ; n > 0; n /= 10 {
		res = string(n%10+48) + res
	}
	return
}

func Multiply1(num1 string, num2 string) string {
	//思路：123*89 = (9*3*1 + 9*2*10 + 9*1*100)*1 + (8*3*1 + 8*2*10 + 8*1*100)*10
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	resStr, count1, count2 := "", 0, 0
	for n2 := len(num2) - 1; n2 >= 0; n2-- {
		count1 = 0
		for n1 := len(num1) - 1; n1 >= 0; n1-- {
			tmp := int(num2[n2]-48) * int(num1[n1]-48)
			a, b := tmp/10, tmp%10
			if a != 0 {
				resStr = addTwoString(resStr, string(a+48)+string(b+48), count1+count2)
			} else {
				resStr = addTwoString(resStr, string(b+48), count1+count2)
			}
			count1++
		}
		count2++
	}
	return resStr
}

func addTwoString(num1 string, num2 string, time int) (res string) {
	carry := false
	for i := 0; i < time; i++ {
		num2 = num2 + "0"
	}
	if len(num2) < len(num1) {
		num2, num1 = num1, num2
	}
	for n2, n1 := len(num2)-1, len(num1)-1; n2 >= 0 || n1 >= 0; {
		if n1 < 0 && !carry {
			res = num2[:n2+1] + res
			break
		}
		val := 0
		if n1 >= 0 {
			val = int(num1[n1]-48) + int(num2[n2]-48)
		} else {
			val = int(num2[n2] - 48)
		}
		if carry {
			val++
			carry = false
		}
		if val >= 10 {
			carry = true
			val -= 10
		}
		res = string(val+48) + res
		n1--
		n2--
	}
	if carry {
		res = "1" + res
	}
	return res
}

//space: O(?)
//time : O(?)
