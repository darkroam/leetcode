package mergeTwoLists

import "testing"
import "reflect"

//TestTwoSum ...
func TestMergeTwoLists(t *testing.T) {
	type testCase struct {
		l1   *ListNode
		l2   *ListNode
		want []int
	}

	num1head := ArrayToList([]int{1, 2, 4})
	num2head := ArrayToList([]int{1, 3, 4})
	num3head := ArrayToList([]int{})
	num4head := ArrayToList([]int{1, 2})
	num5head := ArrayToList([]int{1, 2})
	num6head := ArrayToList([]int{})
	num7head := ArrayToList([]int{})
	num8head := ArrayToList([]int{})
	num9head := ArrayToList([]int{1})
	num10head := ArrayToList([]int{})
	num11head := ArrayToList([]int{})
	num12head := ArrayToList([]int{1})
	num13head := ArrayToList([]int{2, 4, 6, 8})
	num14head := ArrayToList([]int{1, 3, 5, 7})

	testGroup := map[string]testCase{
		"case1": testCase{num1head, num2head, []int{1, 1, 2, 3, 4, 4}},
		"case2": testCase{num3head, num4head, []int{1, 2}},
		"case3": testCase{num5head, num6head, []int{1, 2}},
		"case4": testCase{num7head, num8head, []int{}},
		"case5": testCase{num9head, num10head, []int{1}},
		"case6": testCase{num11head, num12head, []int{1}},
		"case7": testCase{num13head, num14head, []int{1, 2, 3, 4, 5, 6, 7, 8}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := ListToArray(MergeTwoLists(tc.l1, tc.l2))
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
