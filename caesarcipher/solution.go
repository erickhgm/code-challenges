package caesarcipher

import "fmt"

func Cipher(s string, shift int) string {
	aCode := int([]rune("a")[0])
	zCode := int([]rune("z")[0])

	var cipher string
	for _, v := range s {
		fmt.Println(string(v))

		code := int(v) + shift

		if code > zCode {
			code = aCode + (code % zCode) - 1
			if code > zCode {
				code = aCode + (code % zCode) - 1
			}
			cipher += string(rune(code))
		} else {
			cipher += string(rune(code))
		}
	}
	return cipher
}
