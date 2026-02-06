package utils

func MaskToken(token string) string {
	if len(token) <= 6 {
		return "***"
	}
	return token[:3] + "****" + token[len(token)-3:]
}
