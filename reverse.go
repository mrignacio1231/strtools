package strtools

func Reverse(s string) string {
	length := len(s)
	runes := make([]rune, length)
	for i := 0; i < length; i++ {
		runes[i] = rune(s[length-1-i])
	}
	return string(runes)
}
