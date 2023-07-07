package piscine

func ToUpper(s string) string {
	char := []rune(s)

	for i := 0; i < len(char); i++ {
		if char[i] >= 'a' && char[i] <= 'z' {
			char[i] = char[i] - ('a' - 'A')
		}
	}
	return string(char)
}
