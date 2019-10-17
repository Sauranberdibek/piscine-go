package piscine

func StrRev(s string) string {
	runes := []rune(s)
	counter := 0 
	for range s {
		counter++ 
	}
	for i, j := 0, counter ; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)

}
