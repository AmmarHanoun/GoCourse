package main

import (
  "fmt"
)

func doubleSlice(){
  var booking = [][]int{
    {0,2,4,6},
    {1,3,5,7},
    {8,10,12,14},
    {9,11,13,15},
  }

 
  fmt.Printf("size of the Double array is: %vx%v\n",len(booking),len(booking[0]))

  printSlice(booking,"Print before Append")
  // Appending a new row to the double slice
  booking = append(booking,[]int{16,18,20,22})
  printSlice(booking,"Print after Append")
}


func printSlice(s [][]int, t string){
  fmt.Println(t)
  // Printing Double slice, row by row.
  for i,v := range s{
    for j,z := range v{
     fmt.Printf("element (%v,%v): %v\n",i,j,z) 
    }
  }
  
}
