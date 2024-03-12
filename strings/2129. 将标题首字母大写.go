package strings

import "strings"

func capitalizeTitle(title string) string {
	words := strings.Split(title, " ")
	var res []string
	for _, word := range words {
		if len(word) > 2 {
			res = append(res, strings.ToUpper(string(word[0]))+strings.ToLower(word[1:]))
		} else {
			res = append(res, strings.ToLower(word))
		}
	}
	return strings.Join(res, " ")
}

func capitalizeTitle_2(title string) string {
	res := ""
	word := []rune{}
	for i, c := range title {
		if c != ' ' {
			word = append(word, c)
		}
		if c == ' ' || i == len(title)-1 {
			for ii, cc := range word {
				if int(cc) >= 65 && int(cc) <= 90 {
					word[ii] = rune(cc + 32)
				}
			}
			if len(word) > 2 {
				word[0] = rune(word[0] - 32)
			}
			res += string(word)
			word = []rune{}
		}
		if c == ' ' {
			res += " "
		}
	}
	return res
}
