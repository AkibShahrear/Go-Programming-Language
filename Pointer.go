package main

import "fmt"

func update(a *int){
*a = *a + 10 
}

func main(){
var y *int  // special type ==>pointer type
var x int

fmt.Println("Value of x is: ",x)
fmt.Println("memory address of x is : ",  &x )

fmt.Println("Value of y is: ",y)
fmt.Println("memory address of x is : ",  &y )

x = 10 //assignment
y = &x //referencing

fmt.Println("Value of x is: ",x)
fmt.Println("y is : ",  y )  //memory address of y
fmt.Println("dereferencing value y is : ",  *y )//the value of memory address ==>dereferencing

update(&x)

fmt.Println("value of x: ",x) // x = 20

}

