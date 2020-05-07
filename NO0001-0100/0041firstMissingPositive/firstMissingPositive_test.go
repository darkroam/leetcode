package firstMissingPositive

import "testing"
import "reflect"

//TestTwoSum ...
func TestSearch(t *testing.T) {
	type testCase struct {
		nums []int
		want int
	}

	testGroup := map[string]testCase{
		"case1": testCase{[]int{1, 2, 0}, 3},
		"case2": testCase{[]int{3, 4, -1, 1}, 2},
		"case3": testCase{[]int{7, 8, 9, 11, 12}, 1},
		"case4": testCase{[]int{1, 2, 3, 4, 5}, 6},
		"case5": testCase{[]int{1, 5, 2, 3, 4}, 6},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := FirstMissingPositive(tc.nums)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
