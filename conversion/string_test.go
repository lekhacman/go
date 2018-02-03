package conversion

import (
	"testing"
)

func TestToString(t *testing.T) {
	cases := []struct {
		input  int
		expect string
	}{
		{-9223372036854775808, "-9223372036854775808"},
		{-2147483648, "-2147483648"},
		{-5, "-5"},
		{0, "0"},
		{5, "5"},
		{2147483647, "2147483647"},
		{9223372036854775807, "9223372036854775807"},
	}

	for _, c := range cases {

		result := ToString(c.input)

		//fmt.Println(result)

		if c.expect != result {
			t.Errorf("ToString(%v) == %v. Expect: %v", c.input, result, c.expect)
		}
	}
}
