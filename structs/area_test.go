package main

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		name string
		shape Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12.0, 6.0}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12,6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea{
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
			}
		})
	}
}
