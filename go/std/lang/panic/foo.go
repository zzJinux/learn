package panic

import (
	"errors"
	"fmt"
	"log"
	"runtime/debug"
)

func panicHere() {
	panic(errors.New("hi"))
}

func aaa() {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("capture in aaa")
			log.Print(string(debug.Stack()))
			if err, ok := e.(error); ok {
				log.Printf("err: %v", err)
				panic(fmt.Errorf("aaa: %w", err))
			}
		}
	}()
	panicHere()
}

func bbb() {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("capture in bbb")
			log.Print(string(debug.Stack()))
			if err, ok := e.(error); ok {
				log.Printf("err: %v", err)
				panic(fmt.Errorf("bbb: %w", err))
			}
		}
	}()
	aaa()
}

func ccc() {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("capture in ccc")
			log.Print(string(debug.Stack()))
			if err, ok := e.(error); ok {
				log.Printf("err: %v", err)
				panic(fmt.Errorf("ccc: %w", err))
			}
		}
	}()
	bbb()
}
