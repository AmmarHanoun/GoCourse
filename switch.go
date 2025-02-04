package main
import (
  "fmt"
  "strings"
)
func switchCase(){
  var city string
  fmt.Println("Enter a city name")
  fmt.Scan(&city)
 switch strings.ToLower(city) {
   case "cairo":
   fmt.Println("You are in Egypt")
 case "London":
   fmt.Println("You are in England")
 case "paris","riyadh":
   fmt.Println("You are in France or Saudi Arabia ")
 default:
   fmt.Println("I don't know where you are")

 }
  
}