package iteration



func Repeat(char string,count int) string {
	var repeat string
	for i := 0; i < count; i++ {
		repeat += char
	}
	return repeat

}
