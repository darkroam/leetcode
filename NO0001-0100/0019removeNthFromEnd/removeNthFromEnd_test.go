package removeNthFromEnd

import "testing"
import "reflect"

//TestTwoSum ...
func TestRemoveNthFromEnd(t *testing.T) {
	type testCase struct {
		head *ListNode
		n    int
		want []int
	}

	num1head := ArrayToList([]int{1, 2, 3, 4, 5})
	num2head := ArrayToList([]int{1, 2, 3, 4, 5})
	num3head := ArrayToList([]int{1, 2, 3, 4, 5})
	num4head := ArrayToList([]int{1, 2})
	num5head := ArrayToList([]int{1, 2})
	num6head := ArrayToList([]int{1})

	testGroup := map[string]testCase{
		"case1": testCase{num1head, 2, []int{1, 2, 3, 5}},
		"case2": testCase{num2head, 1, []int{1, 2, 3, 4}},
		"case3": testCase{num3head, 5, []int{2, 3, 4, 5}},
		"case4": testCase{num4head, 2, []int{2}},
		"case5": testCase{num5head, 1, []int{1}},
		"case6": testCase{num6head, 1, []int{}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := ListToArray(RemoveNthFromEnd(tc.head, tc.n))
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
