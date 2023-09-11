package lang

import (
	"fmt"
	"testing"
)

func Test_map_ww(_ *testing.T) {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a" // First conflicting access.
		c <- true
	}()
	m["2"] = "b" // Second conflicting access.
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func Test_map_rw(_ *testing.T) {
	c := make(chan bool)
	m := make(map[string]string)
	m["1"] = "a"
	go func() {
		_, ok := m["1"] // First conflicting access.
		c <- ok
	}()
	m["2"] = "b" // Second conflicting access.
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}
