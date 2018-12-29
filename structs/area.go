package main

type Rectangle struct {
	width, height float64
}

func Area(r Rectangle) float64 {
	return r.width * r.height
}
