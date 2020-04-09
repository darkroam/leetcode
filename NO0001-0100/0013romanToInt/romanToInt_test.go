package romanToInt

import "testing"
import "reflect"

//TestTwoSum ...
func TestReverse(t *testing.T) {
	type testCase struct {
		num  string
		want int
	}

	//testGroup := []testCase{
	//testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	//}

	testGroup := map[string]testCase{
		"case1": testCase{"III", 3},
		"case2": testCase{"IV", 4},
		"case3": testCase{"IX", 9},
		"case4": testCase{"LVIII", 58},
		"case5": testCase{"MCMXCIV", 1994},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := RomanToInt(tc.num)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
