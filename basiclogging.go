package main

import (
	"fmt"
	"errors"
)

func main() {

	var x int
	y := 3.57895 
	x = 2
	var truth bool
	truth = false
	var err error
	err = errors.New("invalid error ")
	fmt.Printf("The value is: %d\n", x)
	fmt.Printf("The float value is: %0.2f\n", y)
	fmt.Printf("The boolean value is: %t\n", truth)
	fmt.Printf("The error is %v %T\n", err, err)


	b := []byte{72, 101, 108, 108, 111} // "Hello"

	// Prints the byte array as a string (assumes UTF-8)
	fmt.Printf("prints Hello: %s\n", b)     // prints Hello
	fmt.Println(string(b))   //  

	u := User{ID: 1, Name: "Surya"}

	fmt.Printf("%v\n", u)   // {1 Surya}       (values)
	fmt.Printf("%+v\n", u)  // {ID:1 Name:Surya} (with field names)
	fmt.Printf("%#v\n", u)  // main.User{ID:1, Name:"Surya"} (Go-syntax)
	fmt.Printf("%T\n", u)   // main.User (type)


	maps1 := map[int]string{ 1 : "Sai", 2 : "leo"}
	fmt.Printf("map: %v\n", maps1)
	fmt.Print(maps1)

}

type User struct {
    ID   int
    Name string
}


