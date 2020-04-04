package myAtoi

import "testing"
import "reflect"

//TestTwoSum ...
func TestReverse(t *testing.T) {
	type testCase struct {
		str  string
		want int
	}

	//testGroup := []testCase{
	//testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	//}

	testGroup := map[string]testCase{
		"case1": testCase{"42", 42},
		"case2": testCase{"     -42", -42},
		"case3": testCase{"4193 with words", 4193},
		"case4": testCase{"words and 987", 0},
		"case5": testCase{"-91283472332", -2147483648},
		"case6": testCase{"+-2", 0},
		"case7": testCase{"9223372036854775808", 2147483647},
		"case8": testCase{"0-1", 0},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := MyAtoi(tc.str)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
