package main

import (
	"fmt"
	"reflect"
)

func main() {
	var short_number int8 = 1
	number := 1
	fmt.Println(reflect.TypeOf(short_number))
	fmt.Println(reflect.TypeOf(number))

	var name string = "f"
	fmt.Println(name)
	sprintf := fmt.Sprintf("Number: %X is cool", 12345567890)
	fmt.Println(sprintf)
}
