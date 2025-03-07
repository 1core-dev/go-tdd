package structs

import "testing"

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 5.0}

	got := Perimeter(rectangle)
	want := 30.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{10.0, 5.0}, hasArea: 50.0},
		{shape: Circle{10.0}, hasArea: 314.1592653589793},
		{shape: Triangle{6.0, 2.0}, hasArea: 6.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}

}
