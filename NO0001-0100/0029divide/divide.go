package divide

//29. 两数相除
//给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
//返回被除数 dividend 除以除数 divisor 得到的商。
//整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2
//示例 1:
//	输入: dividend = 10, divisor = 3
//	输出: 3
//	解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
//示例 2:
//	输入: dividend = 7, divisor = -3
//	输出: -2
//	解释: 7/-3 = truncate(-2.33333..) = -2
//提示：
//	被除数和除数均为 32 位有符号整数。
//	除数不为 0。
//	假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,  2^31 − 1]。本题中，如果除法结果溢出，则返回 2^31 − 1。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/string-to-integer-atoi
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//import "fmt"

func Divide(dividend int, divisor int) int {
	if dividend == -(1<<31) && divisor == -1 {
		return 1<<31 - 1
	}
	if dividend < 0 && divisor < 0 {
		return Divide(-dividend, -divisor)
	}
	if dividend < 0 && divisor > 0 {
		return -Divide(-dividend, divisor)
	}
	if dividend > 0 && divisor < 0 {
		return -Divide(dividend, -divisor)
	}
	if dividend == 0 { //此判断前移，导致耗时增加 ??
		return 0
	}
	result := 0
	for i := 31; i >= 0; i-- {
		if (dividend >> i) >= divisor { //找出足够大的数2^n*divisor
			result += 1 << i         //将结果加上2^n
			dividend -= divisor << i //将被除数减去2^n*divisor
		}
	}
	return result

}

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend == -(1<<31) && divisor == -1 {
		return 1<<31 - 1
	}
	negative, result := (dividend^divisor) < 0, 0 //用异或来计算是否符号相异
	if dividend < 0 {
		dividend *= -1
	}
	if divisor < 0 {
		divisor *= -1
	}
	for i := 31; i >= 0; i-- {
		if (dividend >> i) >= divisor { //找出足够大的数2^n*divisor
			result += 1 << i         //将结果加上2^n
			dividend -= divisor << i //将被除数减去2^n*divisor
		}
	}
	if negative { //符号相异取反
		return -result
	} else {
		return result
	}

}

//space: O(?)
//time : O(?)
