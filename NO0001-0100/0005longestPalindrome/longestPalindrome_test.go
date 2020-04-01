package longestPalindrome

import "testing"
import "reflect"

//TestTwoSum ...
func TestLongestPalindrome(t *testing.T) {
	type testCase struct {
		sub  string
		want string
	}

	//testGroup := []testCase{
	//testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	//}

	testGroup := map[string]testCase{
		"case1": testCase{"babad", "bab"},
		"case2": testCase{"cbbd", "bb"},
		"case3": testCase{"bbbbbb", "bbbbbb"},
		"case4": testCase{"bbbbbbb", "bbbbbbb"},
		"case5": testCase{"b", "b"},
		"case6": testCase{"", ""},
		"case7": testCase{"ac", "a"},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := LongestPalindrome(tc.sub)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
