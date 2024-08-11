package array_slice

import (
	"reflect"
	"testing"
)

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

	t.Run("sum all from slices", func(t *testing.T) {
		arrange := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{0, 9}

		if reflect.DeepEqual(arrange, expected) {
			t.Errorf("result '%d', expected '%d'", arrange, expected)
		}
	})

	t.Run("", func(t *testing.T) {
		result := SumAllRest([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result %v, expected %v", result, expected)
		}
	})

}

func TestSumAllRest(t *testing.T) {

	verifysums := func(t *testing.T, result, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result %v, expected %v", result, expected)
		}
	}
	t.Run("Sum some slices", func(t *testing.T) {
		result := SumAllRest([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		verifysums(t, result, expected)
	})

	t.Run("Sum some empty slices", func(t *testing.T) {
		result := SumAllRest([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}

		verifysums(t, result, expected)
	})

}
