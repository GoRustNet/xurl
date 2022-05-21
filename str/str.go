package str

const Empty string = ""

func FixLenWithDesc(s string, maxLen int) (fixedStr string, isFixed bool) {
	r := []rune(s)
	sLen := len(r)
	if sLen <= maxLen {
		return s, false
	}
	return string(r[:maxLen]), true
}

func FixLen(s string, maxLen int) string {
	ss, _ := FixLenWithDesc(s, maxLen)
	return ss
}
func FixLenRef(s *string, maxLen int) {
	*s = FixLen(*s, maxLen)
}

func IsEmpty(s string) bool {
	return s == Empty
}
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}
