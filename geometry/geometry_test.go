package geometry

import "testing"

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{
		Width:  10.0,
		Height: 10.0,
	}

	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

// func TestArea(t *testing.T) {

// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()

// 		if got != want {
// 			t.Errorf("got %g, want %g", got, want)
// 		}
// 	}

// 	t.Run("Area for rectangle", func(t *testing.T) {

// 		rectangle := Rectangle{
// 			Width:  12.0,
// 			Height: 6.0,
// 		}

// 		checkArea(t, rectangle, 72)
// 	})
// 	t.Run("Area for circle", func(t *testing.T) {

// 		circle := Circle{
// 			Radius: 10.0,
// 		}

// 		checkArea(t, circle, 314.1592653589793)
// 	})
// }

// Table drive tests
func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}

}
