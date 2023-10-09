package main

import (
	"fmt"
)

// struct
type User struct {
	id string
	name string
	surname string
	age int
	access Access
	favoriteFoods map[string]string
}
type Access struct {
	admin map[string] bool
	paid map[string] bool
}
func main(){
	var message string = "Hello from a fully declared and typed variable"
	inferred := "Hello from a inferred type variable"
	const constVar string = "This variable cannot be changed"

	var favFood = make(map[string]string)

	favFood["Take-Out"] = "Pizza"
	favFood["BBQ"] = "Ribs"
	favFood["Comfort"] = "Spaghetti and Mince"

	var John User

	John.id = "asdaswewq213123124shfgASZ_sadas"
	John.name = "John"
	John.surname = "Doe"
	John.age = 35
	John.access.admin = map[string]bool{"admin": true}
	John.access.paid = map[string]bool{"paid": true}
	John.favoriteFoods = favFood

	


	fmt.Println(message)
	fmt.Println(inferred)
	fmt.Println(constVar)
	fmt.Print(John)
}
