package threeSumClosest

import "testing"
import "reflect"

//TestTwoSum ...
func TestThreeSumClosest(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		want   int
	}

	//testGroup := []testCase{
	//testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	//}

	testGroup := map[string]testCase{
		"case1": testCase{[]int{-1, 2, 1, -4}, 1, 2},
		"case2": testCase{[]int{1, 1, 1, 0}, -100, 2},
		"case3": testCase{[]int{1, 1, -1, -1, 3}, -1, -1},
		"case4": testCase{[]int{6, -18, -20, -7, -15, 9, 18, 10, 1, -20, -17, -19, -3, -5, -19, 10, 6, -11, 1, -17, -15, 6, 17, -18, -3, 16, 19, -20, -3, -17, -15, -3, 12, 1, -9, 4, 1, 12, -2, 14, 4, -4, 19, -20, 6, 0, -19, 18, 14, 1, -15, -5, 14, 12, -4, 0, -10, 6, 6, -6, 20, -8, -6, 5, 0, 3, 10, 7, -2, 17, 20, 12, 19, -13, -1, 10, -1, 14, 0, 7, -3, 10, 14, 14, 11, 0, -4, -15, -8, 3, 2, -5, 9, 10, 16, -4, -3, -9, -8, -14, 10, 6, 2, -12, -7, -16, -6, 10}, -52, -52},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := ThreeSumClosest(tc.nums, tc.target)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
