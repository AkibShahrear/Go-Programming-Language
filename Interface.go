package main

import "fmt"

type Email interface{

Write(string, string , string)

send( )

Read( )

}

type Test struct{
To string
From string
Subject string
Message string
}

func (t Test) Write( to string, from string, subject string){
fmt.Println(to, from, subject)
}

func (t Test) send(){
fmt.Println(t.To,  "email sent")
}

func (t Test) read( ){
fmt.Println(t.From,"receives from" )
}

func main( ){

//var e Email
var tst Test

tst.To ="Sakib@gmail.com"

tst.From ="Akib@gmail.com"

tst.Subject = "Test email"

tst.Message ="Hello this is a test email"

tst.Write(tst.To , tst.From, tst.Subject)


}

