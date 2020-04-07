package maxArea

import "testing"
import "reflect"

//TestTwoSum ...
func TestReverse(t *testing.T) {
	type testCase struct {
		height []int
		want   int
	}

	//testGroup := []testCase{
	//testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	//}

	testGroup := map[string]testCase{
		"case1": testCase{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := MaxArea(tc.height)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
