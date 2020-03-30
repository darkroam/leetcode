package lengthOfLongestSubstring

import "testing"
import "reflect"

//TestTwoSum ...
func TestLengthOfLongestSubstring(t *testing.T) {
	type testCase struct {
		sub  string
		want int
	}

	//testGroup := []testCase{
	//testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	//}

	testGroup := map[string]testCase{
		"case1": testCase{"abcabcbb", 3},
		"case2": testCase{"bbbbb", 1},
		"case3": testCase{"pwwkew", 3},
		"case4": testCase{"bwf", 3},
		"case5": testCase{"wwkew", 3},
		"case6": testCase{"abcdcbef", 5},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := LengthOfLongestSubstring(tc.sub)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
