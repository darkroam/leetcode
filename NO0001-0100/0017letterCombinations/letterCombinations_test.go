package letterCombinations

import "testing"
import "reflect"

//TestTwoSum ...
func TestLetterCombinations(t *testing.T) {
	type testCase struct {
		digits string
		want   []string
	}

	//testGroup := []testCase{
	//testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	//}

	testGroup := map[string]testCase{
		"case1": testCase{"23", []string{"ad", "bd", "cd", "ae", "be", "ce", "af", "bf", "cf"}},
		"case2": testCase{"", []string{}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := LetterCombinations(tc.digits)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
