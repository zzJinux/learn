package tpack2

import (
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestFoo(t *testing.T) {
	t.Log(os.Getpid())
	if p := os.Getenv("OUTPUT_FILE"); p != "" {
		f, err := os.OpenFile(p, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			t.Fatal(err)
		}
		for i := 0; i < 5; i++ {
			_, _ = f.WriteString("from tpack2\n")
			jitterD := time.Second/2 + time.Duration(rand.Intn(201)-100)*time.Millisecond
			time.Sleep(jitterD)
		}
		_ = f.Close()
	}
}
