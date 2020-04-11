package threeSum

import "testing"
import "reflect"

//TestTwoSum ...
func TestReverse(t *testing.T) {
	type testCase struct {
		nums []int
		want [][]int
	}

	//testGroup := []testCase{
	//testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	//}

	testGroup := map[string]testCase{
		"case1": testCase{[]int{10, -1, 0, 1, 2, -1, -4}, [][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}}},
		"case2": testCase{[]int{0, 0, 0, 0}, [][]int{[]int{0, 0, 0}}},
		"case3": testCase{[]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}, [][]int{[]int{-4, -2, 6}, []int{-4, 0, 4}, []int{-4, 1, 3}, []int{-4, 2, 2}, []int{-2, -2, 4}, []int{-2, 0, 2}}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := ThreeSum(tc.nums)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
