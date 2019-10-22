package piscine

func TrimAtoi(s string) int {

	var g int
	let := 'a'
	sing := 1
	isNumberMet := false

	for index := range s {
		let = rune(s[index])
		if let == '-' && !isNumberMet {
			sing = sing * (-1)
			continue
		} else if let < '0' || let > '9' {
			continue
		} else {
			isNumberMet = true
			numb := 0
			for i := '0'; i < let; i++ {
				numb++
			}
			g = g*10 + (numb)
		}
	}
	return g * sing
}
