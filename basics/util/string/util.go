package string

func Reverse(s string) string {
	r := []rune(s)
	i, j := 0, len(r)-1
	for i < j {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}

	return string(r)
}
