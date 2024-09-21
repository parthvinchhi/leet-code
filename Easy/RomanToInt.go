package easy

func RomanToInt(s string) int {
	l := len(s)
	roman := make(map[byte]int)

	roman['I'] = 1
	roman['V'] = 5
	roman['X'] = 10
	roman['L'] = 50
	roman['C'] = 100
	roman['D'] = 500
	roman['M'] = 1000

	if l == 0 {
		return 0
	}

	if l == 1 {
		return roman[s[0]]
	}

	sum := roman[s[l-1]]

	for i := l - 2; i >= 0; i-- {
		if roman[s[i]] < roman[s[i+1]] {
			sum -= roman[s[i]]
		} else {
			sum += roman[s[i]]
		}
	}

	return sum

}
