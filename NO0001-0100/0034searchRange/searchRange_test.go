package searchRange

import "testing"
import "reflect"

//TestTwoSum ...
func TestSearch(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		want   []int
	}

	testGroup := map[string]testCase{
		"case1": testCase{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		"case2": testCase{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		"case3": testCase{[]int{1}, 1, []int{0, 0}},
		"case4": testCase{[]int{1, 2, 3, 3, 3, 3, 4, 5, 9}, 3, []int{2, 5}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := SearchRange(tc.nums, tc.target)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
