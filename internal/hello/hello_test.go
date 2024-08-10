package hello

import "testing"

func TestAdd(t *testing.T) {

	verifyResult := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	}

	t.Run("Say hello to specific person", func(t *testing.T) {
		result := Hello("edvaldo")
		expected := "Hello, edvaldo"

		verifyResult(t, result, expected)
	})

	t.Run("Say hello world when any empty string send", func(t *testing.T) {
		result := Hello("")
		expected := "Hello, world"

		verifyResult(t, result, expected)
	})

}
