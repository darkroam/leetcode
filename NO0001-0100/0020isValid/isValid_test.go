package isValid

import "testing"
import "reflect"

//TestTwoSum ...
func TestIsValid(t *testing.T) {
	type testCase struct {
		str  string
		want bool
	}

	//testGroup := []testCase{
	//testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	//}

	testGroup := map[string]testCase{
		"case1": testCase{"()", true},
		"case2": testCase{"()[]{}", true},
		"case3": testCase{"(]", false},
		"case4": testCase{"([)]", false},
		"case5": testCase{"{[]}", true},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := IsValid(tc.str)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
