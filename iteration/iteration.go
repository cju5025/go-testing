package iteration

func Repeat(character string, times int) string {
	var characterToRepeat string
	for i := 0; i < times; i++ {
		characterToRepeat += character
	}
	return characterToRepeat
}
