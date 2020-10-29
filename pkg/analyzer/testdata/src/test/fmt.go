package test

import "fmt"

func Foo() {
	a := "%d"
	fmt.Sprintf(a, 10) // want "avoid tail format"
	fmt.Sprint(10)
}
