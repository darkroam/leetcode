package mergeKLists

import "testing"
import "reflect"

//TestTwoSum ...
func TestMergeKLists(t *testing.T) {
	type testCase struct {
		lists []*ListNode
		want  []int
	}

	num1head := ArrayToList([]int{1, 4, 5})
	num2head := ArrayToList([]int{1, 3, 4})
	num3head := ArrayToList([]int{2, 6})

	testGroup := map[string]testCase{
		"case1": testCase{[]*ListNode{num1head, num2head, num3head}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := ListToArray(MergeKLists(tc.lists))
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
