package fourSum

import "testing"
import "reflect"

//TestTwoSum ...
func TestFourSum(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		want   [][]int
	}

	//testGroup := []testCase{
	//testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	//}

	testGroup := map[string]testCase{
		"case1": testCase{[]int{1, 0, -1, 0, -2, 2}, 0, [][]int{[]int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}, []int{-1, 0, 0, 1}}},
		"case2": testCase{[]int{0, 0, 0, 0}, 0, [][]int{[]int{0, 0, 0, 0}}},
		"case3": testCase{[]int{-1, 0, 1, 2, -1, -4}, -1, [][]int{[]int{-4, 0, 1, 2}, []int{-1, -1, 0, 1}}},
		"case4": testCase{[]int{-5, -4, -3, -2, -1, 0, 0, 1, 2, 3, 4, 5}, 0, [][]int{[]int{-5, -4, 4, 5}, []int{-5, -3, 3, 5}, []int{-5, -2, 2, 5}, []int{-5, -2, 3, 4}, []int{-5, -1, 1, 5}, []int{-5, -1, 2, 4}, []int{-5, 0, 0, 5}, []int{-5, 0, 1, 4}, []int{-5, 0, 2, 3}, []int{-4, -3, 2, 5}, []int{-4, -3, 3, 4}, []int{-4, -2, 1, 5}, []int{-4, -2, 2, 4}, []int{-4, -1, 0, 5}, []int{-4, -1, 1, 4}, []int{-4, -1, 2, 3}, []int{-4, 0, 0, 4}, []int{-4, 0, 1, 3}, []int{-3, -2, 0, 5}, []int{-3, -2, 1, 4}, []int{-3, -2, 2, 3}, []int{-3, -1, 0, 4}, []int{-3, -1, 1, 3}, []int{-3, 0, 0, 3}, []int{-3, 0, 1, 2}, []int{-2, -1, 0, 3}, []int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}, []int{-1, 0, 0, 1}}},
		//[]int{-5, -4, 4, 5}, []int{-5, -3, 3, 5}, []int{-5, -2, 2, 5}, []int{-5, -2, 3, 4}, []int{-5, -1, 1, 5}, []int{-5, -1, 2, 4}, []int{-5, 0, 0, 5}, []int{-5, 0, 1, 4}, []int{-5, 0, 2, 3}, []int{-4, -3, 2, 5}, []int{-4, -3, 3, 4}, []int{-4, -2, 1, 5}, []int{-4, -2, 2, 4}, []int{-4, -1, 0, 5}, []int{-4, -1, 1, 4}, []int{-4, -1, 2, 3}, []int{-4, 0, 0, 4}, []int{-4, 0, 1, 3}, []int{-3, -2, 0, 5}, []int{-3, -2, 1, 4}, []int{-3, -2, 2, 3}, []int{-3, -1, 0, 4}, []int{-3, -1, 1, 3}, []int{-3, 0, 0, 3}, []int{-3, 0, 1, 2}, []int{-2, -1, 0, 3}, []int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}, []int{-1, 0, 0, 1}}},
		"case5": testCase{[]int{1, -2, -5, -4, -3, 3, 3, 5}, -11, [][]int{[]int{-5, -4, -3, 1}}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := FourSum(tc.nums, tc.target)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
