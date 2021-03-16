package main

import "fmt"

type Doctor struct{
Name string
Education string
Age int
Salary float32
}

//method
func (d Doctor )Speak() {

fmt.Println(d.Name , "can speak")

}

//method
func (d Doctor )getName() string{

return d.Name
//fmt.Println(d.Name , "can speak")

}

//method
func (d Doctor )getSalaryInfo() float32{

return d.Salary
//fmt.Println(d.Name , "can speak")

}

func main(){

//var d = Doctor{"Tareq" , "MBBS", 30,500000,}
var d =  Doctor{Name:"Tareq" , Education: "MBBS", Age: 30, Salary:500000}
fmt.Println(d.Name, d.Education, d.Age)
d.Speak()
fmt.Println(d.getName())
fmt.Println(d.getSalaryInfo())


}
 