package gopackages

import "strings"

func SplitString(inputString string, symbol string) []string {
	return strings.Split(inputString, symbol)
}
