package removeElement

import "testing"
import "reflect"

//TestTwoSum ...
func TestRemoveDuplicates(t *testing.T) {
	type testCase struct {
		nums []int
		val  int
		want int
	}

	testGroup := map[string]testCase{
		"case1": testCase{[]int{3, 2, 2, 3}, 3, 2},
		"case2": testCase{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := RemoveElement(tc.nums, tc.val)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
