package piscine

func ConcatParams(args []string) string {
	r := ""
	ln := 0
	for i := range args {
		ln = i
	}
	for i := 0; i <= ln; i++ {
		if i != ln {
			r = r + args[i] + "\n"
		} else {
			r = r + args[i]
		}
	}
	return r
}
