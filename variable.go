package main

import "fmt"

func main( ){
var name1 string = "Akib" // Variable [name] variable type

//var name string ==> decleration

// name = "Akib"(value) ==>initialization

var name , city string = "Sakib" , "Chittagong"//multiple variable

fmt.Println(name , city , name1 )
//OUTPUT: Akib Chittagong,Akib

//char (rune)

var c rune = 'A'

//int

var age int = 20

//float

var result float32 = 50.95

fmt.Printf("%c\n",c)
fmt.Println(c , age , result)

}