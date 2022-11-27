package Test
func len(str string) int {
	i := 0
	for str[i] != 0 {
		i++
	}
  //冲突
	return i
}
