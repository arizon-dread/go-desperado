package businesslayer

import (
	"example/ws-test/models"
	"strings"
)

func GetTextAsDesperado(input models.Desperado) string {

	var output string = ""
	consonants_lower := "bcdfghjklmnpqrstvxz"
	consonants_upper := "BCDFGHJKLMNPQRSTVXZ"

	for _, c := range input.Text {
		if strings.ContainsRune(consonants_lower, c) {
			c_str := string(c)
			output = output + c_str + "o" + c_str
		} else if strings.ContainsRune(consonants_upper, c) {
			c_str := string(c)
			output = output + c_str + "O" + c_str
		} else {
			output = output + string(c)
		}
	}
	return output
}
