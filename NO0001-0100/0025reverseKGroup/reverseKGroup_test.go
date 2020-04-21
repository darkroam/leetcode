package reverseKGroup

import "testing"
import "reflect"

//TestTwoSum ...
func TestMergeKLists(t *testing.T) {
	type testCase struct {
		head *ListNode
		k    int
		want []int
	}

	num1head := ArrayToList([]int{1, 2, 3, 4, 5})
	num2head := ArrayToList([]int{1, 2, 3, 4, 5})
	num3head := ArrayToList([]int{1, 2, 3, 4, 5})
	num4head := ArrayToList([]int{1, 2, 3, 4, 5})
	num5head := ArrayToList([]int{})

	testGroup := map[string]testCase{
		"case1": testCase{num1head, 2, []int{2, 1, 4, 3, 5}},
		"case2": testCase{num2head, 3, []int{3, 2, 1, 4, 5}},
		"case3": testCase{num3head, 5, []int{5, 4, 3, 2, 1}},
		"case4": testCase{num4head, 6, []int{1, 2, 3, 4, 5}},
		"case5": testCase{num5head, 9, []int{}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := ListToArray(ReverseKGroup(tc.head, tc.k))
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
