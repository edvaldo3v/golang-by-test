package integer

import "testing"

func TestAdd(t *testing.T) {

	verifyResult := func(t *testing.T, result, expected int) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%d', expected '%d'", result, expected)
		}
	}

	t.Run("Sum two number", func(t *testing.T) {
		result := Add(2, 2)
		expected := 4

		verifyResult(t, result, expected)
	})
}
