package strStr

import "testing"
import "reflect"

//TestTwoSum ...
func TestStrStr(t *testing.T) {
	type testCase struct {
		haystack string
		needle   string
		want     int
	}

	testGroup := map[string]testCase{
		"case1": testCase{"hello", "ll", 2},
		"case2": testCase{"aaaaa", "bba", -1},
		"case3": testCase{"aaaaa", "", 0},
		"case4": testCase{"", "a", -1},
		"case5": testCase{"a", "a", 0},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := StrStr(tc.haystack, tc.needle)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
