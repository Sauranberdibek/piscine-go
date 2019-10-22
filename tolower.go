package piscine

func ToLower(s string) string {
	i := []rune(s)
	for index, d := range i {
		if d >= 'A' && d <= 'Z' {
			i[index] = d + 32
		}
	}
	return string(i)
}
