package bench

var digitMap = map[rune]struct{}{
	'1': {},
	'2': {},
	'3': {},
	'4': {},
	'5': {},
	'6': {},
	'7': {},
	'8': {},
	'9': {},
	'0': {},
}

func IsDigit(r rune) bool {
	_, ok := digitMap[r]
	return ok
}
