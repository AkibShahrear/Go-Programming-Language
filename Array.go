package main 

import "fmt"

func main(){

//lengthy way
//var students [3]string  // same type

//data set
//students[0] = "Asgor"
//students[1] = "Mainul"
//students[2] = "Anonnya"

//data get
 //fmt.Println(students[2])
//fmt.Println(len(students))
//Implicit
//short hand way , string literals
 student :=  [...]string{"Akib","Sakib","Fahad","Rakib","Bakib"}
fmt.Println(student , len(student))


}