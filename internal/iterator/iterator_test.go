package iterator

import "testing"

func TestRepeat(t *testing.T) {

	verifyResult := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	}

	t.Run("Repeat character", func(t *testing.T) {
		// Arrange
		// Act
		result := Repeat("a")
		expected := "aaa"

		// Assert
		verifyResult(t, result, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
