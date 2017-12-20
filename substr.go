package strtools

// Substr returns part of a string
func Substr(s string, start int, length int) string {
	if length > len(s) {
		length = len(s)
	}
	if start < 0 {
		start = length + start
	}
	return s[start:length]
}
