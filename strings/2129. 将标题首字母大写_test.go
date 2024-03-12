package strings

import (
	"encoding/json"
	"fmt"
	"playground/utils"
	"testing"
)

func Test_capitalizeTitle(t *testing.T) {
	cases := []string{}
	if err := json.Unmarshal([]byte(`[
		"capiTalIze tHe titLe", "Capitalize The Title",
		"First leTTeR of EACH Word", "First Letter of Each Word",
		"i lOve leetcode", "i Love Leetcode"
	]`), &cases); err != nil {
		t.Fatalf("json.Unmarshal error: %v\n", err)
	}
	us := 2
	cnt := len(cases) / us
	for i := 0; i < cnt; i++ {
		input := cases[i*us]
		want := cases[i*us+1]
		testEach_capitalizeTitle(t, capitalizeTitle, input, want)
		testEach_capitalizeTitle(t, capitalizeTitle_2, input, want)
	}
}

func testEach_capitalizeTitle(t *testing.T, fn func(string) string, input, want string) {
	ans := fn(input)
	utils.AssertEqual(t, ans, want,
		fmt.Sprintf("Not equal. input=%q", input))
}
