package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Prathamesh")
		want := "Hello, Prathamesh"
		
		assertCorrectMessage(t, got, want)


	})
	t.Run("saying hello word when empty string is given", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		//comment
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string)  {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}