package strings

import (
	"encoding/json"
	"fmt"
	"playground/utils"
	"testing"
)

func Test_maximumOddBinaryNumber(t *testing.T) {
	cases := []string{}
	if err := json.Unmarshal([]byte(`[
		"010", "001",
		"0101", "1001",
		"1", "1",
		"", ""
	]`), &cases); err != nil {
		t.Fatalf("json.Unmarshal error: %v\n", err)
	}
	funcs := []func(string) string{
		maximumOddBinaryNumber,
		maximumOddBinaryNumber2,
	}
	us := 2
	cnt := len(cases) / us
	for i := 0; i < cnt; i++ {
		input := cases[i*us]
		want := cases[i*us+1]
		for _, fn := range funcs {
			ans := fn(input)
			utils.AssertEqual(t, ans, want,
				fmt.Sprintf("Not equal. input=%q", input))
		}
	}
}
