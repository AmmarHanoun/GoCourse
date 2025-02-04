package main
import (
  "fmt"
)

func displayMap(code string,conference string,availableTickets uint,unitPrice uint,displayMap map[string]map[string]map[uint]uint ){

  displayMap = map[string]map[string]map[uint]uint{
    code: {
      conference: {
         availableTickets: unitPrice,
},
    },
  }
  printMap(displayMap)
}
// End Of displayMap()
func printMap(myMap map[string]map[string]map[uint]uint){
  for i,v := range myMap{
    for j,z := range v{
      for k,y := range z{
        fmt.Printf("code: %v\t conference: %v\t AvailavleTickets: %v\t unitPrice: %v\t",i,j,k,y)
      }
    }
    
  }
}
// End Of printMap()

