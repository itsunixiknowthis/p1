package p1

import "testing"

func TestF(t *testing.T) {
    want := 123
    if got := A; got != want {
        t.Errorf("Hello() = %d, want %d", got, want)
    }
}

