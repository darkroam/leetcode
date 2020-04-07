package findMedianSortedArrays

import "testing"
import "reflect"

//TestTwoSum ...
func TestFindMedianSortedArrays(t *testing.T) {
	type testCase struct {
		nums1 []int
		nums2 []int
		want  float64
	}

	//testGroup := []testCase{
	//testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	//}

	testGroup := map[string]testCase{
		"case1":  testCase{[]int{1, 3}, []int{2}, float64(2)},
		"case2":  testCase{[]int{1, 2}, []int{3, 4}, float64(2.5)},
		"case3":  testCase{[]int{1, 2, 3}, []int{}, float64(2)},
		"case4":  testCase{[]int{}, []int{1, 2, 3}, float64(2)},
		"case5":  testCase{[]int{1, 2, 3, 4}, []int{}, float64(2.5)},
		"case6":  testCase{[]int{}, []int{1, 2, 3, 4}, float64(2.5)},
		"case7":  testCase{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}, float64(2.5)},
		"case8":  testCase{[]int{1, 2, 3, 4}, []int{1, 2, 3}, float64(2)},
		"case9":  testCase{[]int{1}, []int{}, float64(1)},
		"case10": testCase{[]int{2, 3}, []int{}, float64(2.5)},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := FindMedianSortedArrays(tc.nums1, tc.nums2)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
