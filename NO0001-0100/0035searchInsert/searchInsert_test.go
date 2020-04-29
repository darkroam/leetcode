package searchInsert

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
		"case1": testCase{[]int{1, 3, 5, 6}, 5, 2},
		"case2": testCase{[]int{1, 3, 5, 6}, 2, 1},
		"case3": testCase{[]int{1, 3, 5, 6}, 7, 4},
		"case4": testCase{[]int{1, 3, 5, 6}, 0, 0},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := SearchInsert(tc.nums, tc.target)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
