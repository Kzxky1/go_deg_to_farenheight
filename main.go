package main


import (
  "fmt"
  "reflect"
)


func main(){
  var input int 
  var conv string
  
  fmt.Printf(" Pls enter d for degree to farenheight f for other way round: ")
  fmt.Scan(&conv)
 
  fmt.Printf(" Pls enter input temprature: ")
  fmt.Scan(&input)
  
  inType := reflect.TypeOf(input)
  println("type of input is :", inType)

    if conv == "d" {
      result := (input*9/5) +32
      fmt.Println("Temprature in farenheight is : ", result)
    }else{
      result := (input*5/9) -32
      fmt.Println("Temprature in degrees is : ", result)
    }
}

