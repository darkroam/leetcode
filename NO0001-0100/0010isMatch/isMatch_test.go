package isMatch

import "testing"
import "reflect"

//TestTwoSum ...
func TestIsMatch(t *testing.T) {
	type testCase struct {
		s, p string
		want bool
	}

	//testGroup := []testCase{
	//testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	//}

	testGroup := map[string]testCase{
		"case1":  testCase{"aa", "a", false},
		"case2":  testCase{"aa", "a*", true},
		"case3":  testCase{"ab", ".*", true},
		"case4":  testCase{"aab", "c*a*b", true},
		"case5":  testCase{"mississippi", "mis*is*p*.", false},
		"case6":  testCase{"a", ".", true},
		"case7":  testCase{"aa", "..", true},
		"case8":  testCase{"a", "a*", true},
		"case9":  testCase{"abc", "abc", true},
		"case10": testCase{"abc", "ab.", true},
		"case11": testCase{"abc", "a.c", true},
		"case12": testCase{"abc", "..c", true},
		"case13": testCase{"abc", "...", true},
		"case14": testCase{"abc", "..*", true},
		"case15": testCase{"abc", "*", false},
		"case16": testCase{"aaa", "ab*a*c*a", true},
		"case17": testCase{"aaaaa", "aaaa", false},
		"case18": testCase{"aaa", "aaaa", false},
		"case19": testCase{"aaa", "aa*aa", true},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := IsMatch(tc.s, tc.p)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
