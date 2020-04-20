package swapPairs

import "testing"
import "reflect"

//TestTwoSum ...
func TestMergeKLists(t *testing.T) {
	type testCase struct {
		head *ListNode
		want []int
	}

	num1head := ArrayToList([]int{1, 2, 3, 4})
	num2head := ArrayToList([]int{1, 2})
	num3head := ArrayToList([]int{1, 2, 3})

	testGroup := map[string]testCase{
		"case1": testCase{num1head, []int{2, 1, 4, 3}},
		"case2": testCase{num2head, []int{2, 1}},
		"case3": testCase{num3head, []int{2, 1, 3}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := ListToArray(SwapPairs(tc.head))
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
