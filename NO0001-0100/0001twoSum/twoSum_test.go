package twosum

import "testing"
import "reflect"

//TestTwoSum ...
func TestTwoSum(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		want   []int
	}

	//testGroup := []testCase{
	//testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	//}

	testGroup := map[string]testCase{
		"case1": testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := TwoSum(tc.nums, tc.target)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}

//BenchmarkTwoSum ...
func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoSum([]int{2, 7, 11, 15}, 9)
	}
}
