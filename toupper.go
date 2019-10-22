package piscine

func ToUpper(s string) string {
	i := []rune(s)
	for index, d := range i {
		if d >= 'a' && d <= 'z' {
			i[index] = d - 32
		}
	}
	return string(i)
}
