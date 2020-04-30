package countAndSay

import "testing"
import "reflect"

//TestTwoSum ...
func TestSearch(t *testing.T) {
	type testCase struct {
		n    int
		want string
	}

	testGroup := map[string]testCase{
		"case1": testCase{1, "1"},
		"case2": testCase{2, "11"},
		"case3": testCase{3, "21"},
		"case4": testCase{4, "1211"},
		"case5": testCase{5, "111221"},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := CountAndSay(tc.n)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
