package reverse

import "testing"
import "reflect"

//TestTwoSum ...
func TestReverse(t *testing.T) {
	type testCase struct {
		x    int
		want int
	}

	//testGroup := []testCase{
	//testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	//}

	testGroup := map[string]testCase{
		"case1": testCase{123, 321},
		"case2": testCase{-123, -321},
		"case3": testCase{120, 21},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Reverse(tc.x)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
