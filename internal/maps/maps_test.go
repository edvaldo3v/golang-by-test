package maps

import "testing"

func TestSearch(t *testing.T) {

	t.Run("Should return currect value from map", func(t *testing.T) {
		dictionary := map[string]string{"text": "htis is just text"}

		get := Seach(dictionary, "text")
		sent := "htis is just text"

		if get != sent {
			t.Errorf("get %q sent %q given %q", get, sent, "text")
		}
	})

}
