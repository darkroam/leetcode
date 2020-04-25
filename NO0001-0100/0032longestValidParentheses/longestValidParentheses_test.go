package longestValidParentheses

import "testing"
import "reflect"

//TestTwoSum ...
func TestIsValid(t *testing.T) {
	type testCase struct {
		s    string
		want int
	}

	testGroup := map[string]testCase{
		"case1": testCase{"(()", 2},
		"case2": testCase{")()())", 4},
		"case3": testCase{"()())", 4},
		"case4": testCase{"()()(", 4},
		"case5": testCase{"(()())()(", 8},
		"case6": testCase{"()(()", 2},
		"case7": testCase{"))))())()()(()", 4},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := LongestValidParentheses(tc.s)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
