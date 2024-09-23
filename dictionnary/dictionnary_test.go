package dictionnary

import "testing"

func TestSearch(t *testing.T) {
	dictionnary := Dictionary{"test": "don't panic, this is a test"}
	t.Run("known word", func(t *testing.T) {

		got, _ := dictionnary.Search("test")
		want := "don't panic, this is a test"

		assertString(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionnary.Search("wellwellwellwhatdowehavehere")
		want := ErrNotFound

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	dictionnary := Dictionary{}
	t.Run("new word", func(t *testing.T) {
		word := "test"
		dictionnary.Add(word, "don't panic, this is a test")

		want := "don't panic, this is a test"
		assertDefinition(t, dictionnary, word, want)
	})
	t.Run("already existing word", func(t *testing.T) {
		word := "test"
		err := dictionnary.Add(word, "don't panic, this is a test")

		want := "don't panic, this is a test"
		assertError(t, err, ErrAlreadyExists)
		assertDefinition(t, dictionnary, word, want)
	})
}

func TestUpdate(t *testing.T) {
	dictionnary := Dictionary{}
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "don't panic, this is a test"
		newDef := "TIME TO PANIC, THE TEST DROPED THE PROD"
		dictionnary.Add(word, def)
		err := dictionnary.Update(word, newDef)
		assertError(t, err, nil)
		assertDefinition(t, dictionnary, word, newDef)
	})
	t.Run("new word", func(t *testing.T) {
		dictionnary := Dictionary{}
		word := "test"
		def := "don't panic, this is a test"

		err := dictionnary.Update(word, def)

		assertError(t, err, ErrDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "don't panic, this is a test"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	assertError(t, err, ErrNotFound)
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	t.Helper()

	got, err := dict.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertString(t, got, definition)
}
