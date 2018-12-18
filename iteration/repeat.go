package iteration

var defaultRepeat = 5

// Repeat return duplication message
func Repeat(s string) (r string) {
	for i := 0; i < defaultRepeat; i++ {
		r = r + s
	}
	return r
}
