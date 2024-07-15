package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("Print Hello World", func(t *testing.T) {
		got := Hello()
		expected := "Hello World"

		Assert(t, got, expected)
	})
}

func Assert(t testing.TB, got, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}
