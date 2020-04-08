package intToRoman

import "testing"
import "reflect"

//TestTwoSum ...
func TestReverse(t *testing.T) {
	type testCase struct {
		num  int
		want string
	}

	//testGroup := []testCase{
	//testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	//}

	testGroup := map[string]testCase{
		"case1": testCase{3, "III"},
		"case2": testCase{4, "IV"},
		"case3": testCase{9, "IX"},
		"case4": testCase{58, "LVIII"},
		"case5": testCase{1994, "MCMXCIV"},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := IntToRoman(tc.num)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
