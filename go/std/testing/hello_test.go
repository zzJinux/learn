package testing

import "testing"

func TestDoFail(t *testing.T) {
	helperAndError(t)
}

func TestDoFailNow(t *testing.T) {
	helperAndFatal(t)
}

func helperAndError(t *testing.T) {
	t.Helper()
	t.Error("t.Error()")
}

func helperAndFatal(t *testing.T) {
	t.Helper()
	t.Fatal("t.Fatal()")
}
