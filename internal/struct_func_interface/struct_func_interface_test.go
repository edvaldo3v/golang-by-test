package struct_func_interface

import "testing"

func TestPerimeter(t *testing.T) {

	verifyResult := func(t *testing.T, result, expected float64) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%.2f', expected '%.2f'", result, expected)
		}
	}

	verifyArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		result := shape.Area()

		if result != expected {
			t.Errorf("resultado %.2f, esperado %.2f", result, expected)
		}
	}

	t.Run("retângulos", func(t *testing.T) {
		retangulo := Rectangle{12.0, 6.0}
		verifyArea(t, retangulo, 72.0)
	})

	t.Run("círculos", func(t *testing.T) {
		circulo := Circle{10}
		verifyArea(t, circulo, 314.1592653589793)
	})

	t.Run("Rectanglw", func(t *testing.T) {
		t.Run("Perimeter", func(t *testing.T) {
			rectangle := Rectangle{10.0, 10.0}
			result := rectangle.Parimeter()
			expected := 40.0

			verifyResult(t, result, expected)
		})
	})

	t.Run("Circle", func(t *testing.T) {
		t.Run("Area", func(t *testing.T) {
			circle := Circle{10}
			result := circle.Area()
			expected := 314.1592653589793

			verifyResult(t, result, expected)
		})
	})

	t.Run("anonimus struct", func(t *testing.T) {
		testsArea := []struct {
			name     string
			shape    Shape
			expected float64
		}{
			{name: "Rectangle", shape: Rectangle{12.0, 6.0}, expected: 72.0},
			{name: "Circle", shape: Circle{ray: 10}, expected: 314.1592653589793},
			{name: "Triangle", shape: Trianglw{12.0, 6.0}, expected: 36.0},
		}

		for _, tt := range testsArea {
			t.Run(tt.name, func(t *testing.T) {
				result := tt.shape.Area()
				if result != tt.expected {
					t.Errorf("resultado %.2f, esperado %.2f", result, tt.expected)
				}
			})

		}
	})

}
