package main

import (
	"fmt"
	"reflect"
)

func main() {
	// check a type with reflect.TypeOf
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	// work with its value with reflect.ValueOf
	p := reflect.ValueOf(x)
	fmt.Println("type of p:", p.Type())
	// not all reflect value can be changed
	fmt.Println("settability of p:", p.CanSet())
	// settable reflect values changes the original
	p2 := reflect.ValueOf(&x)
	v := p2.Elem()
	fmt.Println("type of p2:", p2.Type())
	fmt.Println("settability of element v:", v.CanSet())
	v.SetFloat(3.14)
	fmt.Println("element v changed x:", x)
	fmt.Println("element v changed x:", v.Interface())
}
