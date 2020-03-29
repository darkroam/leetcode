package addtwonumbers

import "testing"
import "reflect"

//TestTwoSum ...
func TestAddTwoNumbers(t *testing.T) {
	type testCase struct {
		l1   *ListNode
		l2   *ListNode
		want []int
	}

	//testGroup := []testCase{
	//testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	//}

	num1head := ArrayToList([]int{2, 4, 3})
	num2head := ArrayToList([]int{5, 6, 4})

	num3head := ArrayToList([]int{5})
	num4head := ArrayToList([]int{5})

	num5head := ArrayToList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	num6head := ArrayToList([]int{5, 6, 4})

	testGroup := map[string]testCase{
		"case1": testCase{num1head, num2head, []int{7, 0, 8}},
		"case2": testCase{num3head, num4head, []int{0, 1}},
		"case3": testCase{num5head, num6head, []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := ListToArray(AddTwoNumbers(tc.l1, tc.l2))
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
