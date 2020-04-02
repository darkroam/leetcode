package convert

import "testing"
import "reflect"

//TestTwoSum ...
func TestConvert(t *testing.T) {
	type testCase struct {
		sub   string
		lines int
		want  string
	}

	//testGroup := []testCase{
	//testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	//}

	testGroup := map[string]testCase{
		"case1": testCase{"LEETCODEISHIRING", 3, "LCIRETOESIIGEDHN"},
		"case2": testCase{"LEETCODEISHIRING", 4, "LDREOEIIECIHNTSG"},
		"case3": testCase{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		"case4": testCase{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Convert(tc.sub, tc.lines)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
