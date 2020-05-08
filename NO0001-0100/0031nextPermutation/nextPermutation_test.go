package nextPermutation

import "testing"
import "reflect"

//TestTwoSum ...
func TestSearch(t *testing.T) {
	type testCase struct {
		nums []int
		want []int
	}

	testGroup := map[string]testCase{
		"case1": testCase{[]int{1, 2, 3}, []int{1, 3, 2}},
		"case2": testCase{[]int{3, 2, 1}, []int{1, 2, 3}},
		"case3": testCase{[]int{1, 1, 5}, []int{1, 5, 1}},
		"case4": testCase{[]int{1, 3, 2}, []int{2, 1, 3}},
		"case5": testCase{[]int{2, 3, 1}, []int{3, 1, 2}},
		"case6": testCase{[]int{4, 2, 0, 2, 3, 2, 0}, []int{4, 2, 0, 3, 0, 2, 2}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			NextPermutation(tc.nums)
			got := tc.nums
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
