package generateParenthesis

import "testing"
import "reflect"

//TestTwoSum ...
func TestGenerateParenthesis(t *testing.T) {
	type testCase struct {
		n    int
		want []string
	}

	testGroup := map[string]testCase{
		"case1": testCase{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := GenerateParenthesis(tc.n)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
