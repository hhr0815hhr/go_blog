package util

import "regexp"

/******************正则**************************/

//email
func RegexpEmail(email string) []int {
	comp := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)`)
	return comp.FindStringIndex(email)
}
