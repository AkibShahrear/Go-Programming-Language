package main

import "fmt"
//example-1
/*
func add(a int, y int ) int  //func identifier(parameter) return type{
//body of function
r := x + y
return r
}

*/
//method 2
/*
func add(x , y int)int{
//body
r := x + y
return r
}*/

//example-3
/*
func add(x , y int)(r int){
//body
r = x + y
return r
}
*/
//example-4

func add(x , y int) (r int){
//body
r = x + y
return 
}

//return multiple values
func rectangle(l int, b int) (area int , parameter int){ 
parameter = 2 * (l + b)
area = l * b
return 
}

func main( ){
//write code here
x := add(10 , 30) //output 40
fmt.Println(x)
a , p := rectangle(10 , 10)
fmt.Println(a , p)

}