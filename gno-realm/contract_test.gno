package demomod

import "testing"

func Test(t *testing.T) {
	{
		got := SayHello()
		expected := "Hello World!"
		if got != expected {
			t.Fatalf("expected %q, got %q.", expected, got)
		}
	}

    {
		got := Sum(2, 5)
		expected := 7
		if got != expected {
			t.Fatalf("expected %d, got %d.", expected, got)
		} 
    }
}
