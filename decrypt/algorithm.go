package decrypt

func Nimbus(str string) string {
	decryptedString := ""
	for _, c := range str {
		asciiCode := int(c)
		character := string(rune(asciiCode - 3))
		decryptedString += character
	}

	return decryptedString
}
