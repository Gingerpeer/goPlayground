package main

import (
	"fmt"
)

func main(){
	var message string = "Hello from a fully declared and typed variable"
	inferred := "Hello from a inferred type variable"
	fmt.Println(message)
	fmt.Println(inferred)
}
