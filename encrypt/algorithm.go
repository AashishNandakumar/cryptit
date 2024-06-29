package encrypt

// Nimbus :the fxn name/identifier should start with capital letter to be used outside
func Nimbus(str string) string {
	encryptedString := ""
	for _, c := range str {
		asciiCode := int(c)
		character := string(rune(asciiCode + 3))
		encryptedString += character
	}

	return encryptedString
}
