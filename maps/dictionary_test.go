package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "test value"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "test value"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		if got == nil {
			t.Fatal("expected to get an error.")
		}
		assertError(t, got, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "testValue definition"

		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "testValue definition"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExists)

		assertDefinition(t, dictionary, word, definition)

	})

}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "initValue definition"
	dictionary := Dictionary{word: definition}

	newDefinition := "newValue definition"
	dictionary.Update(word, newDefinition)

	assertDefinition(t, dictionary, word, newDefinition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}

	assertStrings(t, got, definition)
}
