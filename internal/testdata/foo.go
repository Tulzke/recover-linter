package testdata

import (
	"fmt"
)

func foo() {
	go func() {
		defer recover()
	}()

	go func() {
		defer func() {
			if msg := recover(); msg != nil {
				fmt.Println(msg)
			}
		}()
		go func() {}()
	}()

	go func() {
		fmt.Println("hi")
	}()

	go foo1()

	go foo2[int](10)
}

var foo1 = func() {
	defer recover()
}

func foo2[T any](_ T) {
	go func() {

	}()
}

type Struct struct {
}

type String string

type Inter interface{}
