package xtools

func isAscii(r byte) bool {
	if r > 32 && r < 127 {
		return true
	}
	return false
}
