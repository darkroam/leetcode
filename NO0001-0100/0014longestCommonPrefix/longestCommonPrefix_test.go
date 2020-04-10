package longestCommonPrefix

import "testing"
import "reflect"

//TestTwoSum ...
func TestReverse(t *testing.T) {
	type testCase struct {
		strs []string
		want string
	}

	//testGroup := []testCase{
	//testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	//}

	testGroup := map[string]testCase{
		"case1": testCase{[]string{"flower", "flow", "flight"}, "fl"},
		"case2": testCase{[]string{"dog", "racecar", "car"}, ""},
		"case3": testCase{[]string{"", "dog", "racecar", "car"}, ""},
		"case4": testCase{[]string{"dog", "racecar", "", "car"}, ""},
		"case5": testCase{[]string{}, ""},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := LongestCommonPrefix(tc.strs)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
