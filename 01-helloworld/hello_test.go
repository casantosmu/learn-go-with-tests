package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello Chris!!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello world!!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world!!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Elodie", "es")
		want := "Hola Elodie!!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in french", func(t *testing.T) {
		got := Hello("Elodie", "fr")
		want := "Bonjour Elodie!!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
