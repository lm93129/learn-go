package MapAndPackage

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}
	want := "this is just a test"

	t.Run("Search", func(t *testing.T) {
		got := Search(dictionary, "test")
		asserStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		got := UnkownWord(dictionary, "unknown")
		if got {
			t.Error(got)
		}
	})
}

func asserStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s' given, '%s'", got, want, "test")
	}
}
