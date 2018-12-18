package iteration

// Repeat return duplication message
func Repeat(s string, t int) (r string) {
	for i := 0; i < t; i++ {
		r = r + s
	}
	return r
}
