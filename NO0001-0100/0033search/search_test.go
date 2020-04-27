package search

import "testing"
import "reflect"

//TestTwoSum ...
func TestSearch(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		want   int
	}

	testGroup := map[string]testCase{
		"case1": testCase{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		"case2": testCase{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		"case3": testCase{[]int{4, 1}, 1, 1},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Search(tc.nums, tc.target)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
