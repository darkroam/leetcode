package removeDuplicates

import "testing"
import "reflect"

//TestTwoSum ...
func TestRemoveDuplicates(t *testing.T) {
	type testCase struct {
		nums []int
		want int
	}

	testGroup := map[string]testCase{
		"case1": testCase{[]int{1, 1, 2}, 2},
		"case2": testCase{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := RemoveDuplicates(tc.nums)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
