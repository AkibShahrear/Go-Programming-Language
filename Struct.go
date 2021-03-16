package main

import "fmt"

type Book struct{
Title string
Author string
ISBN string
Price float32
Pages int
}

type Student struct{
Color [ ]string
Name string
Age string
Roll int
}

func main(){

//Struct ==> structure
var b1 Book

b1.Title = "An Introduction to Programming in Go"
b1.Author = "CALEB DOXY"
b1.ISBN = "978-1478355823"
b1.Price = 10.50
b1.Pages = 165

fmt.Println(b1)
fmt.Println(b1.Title)
fmt.Println(b1.Author)
fmt.Println(b1.ISBN)
fmt.Println(b1.Price)
fmt.Println(b1.Pages)
}