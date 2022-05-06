package iteration

func Repeat(char string, repeatCount int) (respated string) {
	for i := 0; i < repeatCount; i++ {
		respated += char
	}
	return
}
