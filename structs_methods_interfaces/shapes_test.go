package structs_methods_interfaces

import (
	"reflect"
	"testing"
)

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 62.83185307179586},
	}

	for _, areaTest := range areaTests {
		t.Run("calculates area of shape"+reflect.TypeOf(areaTest.shape).String(), func(t *testing.T) {
			got := areaTest.shape.Area()
			if got != areaTest.want {
				t.Errorf("got %g want %g", got, areaTest.want)
			}
		})
	}

}
