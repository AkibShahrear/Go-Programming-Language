
package main

import "fmt"
import "reflect"

func main(){

//slice

//var students [3]string

//students[0] ="Asgor"
//students[1] = "Mainul"
//students[2] = "Anonnya"

//x := students[0:2]
//2nd way:
//x := make([]string , 5) 

//long hand format:
var fruits []string

fruits = append(fruits, "Apple","Banana")

fmt.Println(fruits, len(fruits))

fmt.Printf("%T",fruits)

b := reflect.TypeOf(fruits).Kind().String()
fmt.Println(b)
  
}