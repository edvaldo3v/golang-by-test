package iterator

const atdRepeat = 3

func Repeat(char string) string {
	var repeat string
	for i := 0; i < atdRepeat; i++ {
		repeat += char
	}
	return repeat
}
