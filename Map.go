package main

import "fmt"

func main(){

//var x map[string]string // var x map[key]pair

x := make(map[string]string)
//Key = value

x["name"] = "Akib"
x["Height"] = "5.7"
x["address"] = "Dhaka"

delete(x , height)

fmt.Print(x["address"])

}

// A map is an unordered collection of key-value pairs
//Associative array / Hash table / dictionary
//panic error ==> program off