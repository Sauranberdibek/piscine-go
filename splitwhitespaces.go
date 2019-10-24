package piscine

func SplitWhiteSpaces(str string) []string {
	a := 0
	ln := 0

	ok := false
	for c := range str {

		if ok && c != 0 && (str[c-1] == '\n' || str[c-1] == '\t' || str[c-1] == ' ') && str[c] != '\n' && str[c] != '\t' && str[c] != ' ' {
			ln++
		}
		if str[c] != '\n' && str[c] != '\t' && str[c] != ' ' {
			ok = true
		}
	}
	ln++

	ans := make([]string, ln)
	index := 0
	r := ""
	for i, w := range str {
		if w == '\n' || w == ' ' || w == '\t' {
			if r != "" {
				ans[index] = r
				index++
				r = ""
				a = i
			}
		} else {
			if w != ' ' {
				r = r + string(w)
			}
		}
	}

	if str[a+1:] != "" && str[a+1:] != " " {
		ans[index] = str[a+1:]
	}
	return ans

}
