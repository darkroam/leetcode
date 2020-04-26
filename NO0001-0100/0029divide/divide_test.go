package divide

import "testing"
import "reflect"

//TestTwoSum ...
func TestIsValid(t *testing.T) {
	type testCase struct {
		dividend int
		divisor  int
		want     int
	}

	testGroup := map[string]testCase{
		"case1": testCase{10, 3, 3},
		"case2": testCase{7, -3, -2},
		"case3": testCase{-7, -3, 2},
		"case4": testCase{-2147483648, -1, 2147483647},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Divide(tc.dividend, tc.divisor)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
