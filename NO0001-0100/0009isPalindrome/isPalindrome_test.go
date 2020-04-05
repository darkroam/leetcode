package isPalindrome

import "testing"
import "reflect"

//TestTwoSum ...
func TestReverse(t *testing.T) {
	type testCase struct {
		x    int
		want bool
	}

	//testGroup := []testCase{
	//testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	//}

	testGroup := map[string]testCase{
		"case1": testCase{121, true},
		"case2": testCase{-121, false},
		"case3": testCase{10, false},
		"case4": testCase{22, true},
		"case5": testCase{2, true},
		"case6": testCase{2222222, true},
		"case7": testCase{22222222, true},
		"case8": testCase{0, true},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := IsPalindrome(tc.x)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
