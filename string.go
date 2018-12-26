package common

import "strconv"

//IsDigit - вернет true, если s является числом
func IsDigit(s string) bool {
	_, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		return false
	}
	return true
}
