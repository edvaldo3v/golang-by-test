package array_slice

import "testing"

func TestSum(t *testing.T) {

	verifyResult := func(t *testing.T, result, expected int, numbers []int) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%d', expected '%d' numbers '%v'", result, expected, numbers)
		}
	}

	t.Run("Sum numbers on array", func(t *testing.T) {
		// Arrange
		numbers := []int{1, 2, 3, 4, 5}

		// Act
		result := Sum(numbers)
		expected := 15

		//Assert
		verifyResult(t, result, expected, numbers)
	})

	t.Run("Sum numbers on slice", func(t *testing.T) {
		// Arrange
		numbers := []int{1, 2, 3, 4, 5}

		// Act
		result := Sum(numbers)
		expected := 15

		//Assert
		verifyResult(t, result, expected, numbers)
	})

}
