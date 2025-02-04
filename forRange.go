package main

import (
	"fmt"
	)

 var name string
func forRange() {

A1 := []string{"Ammar","Yasser","Mohammad","Arwa","Juliana","Lana"}

 A1= addToList(A1,"Fatimah")
	//A1 = append(A1,"Fatimah")
	fmt.Println("Updated List..")
	for i,v := range A1{
		fmt.Printf("%v: %v\n",i+1,v)
	}
	
}

func addToList(l1 []string,s1 string)[]string{
fmt.Println("Enter a name you like to add to the list..")
	fmt.Scan(&name)
	l1 = append(l1,s1,name)
	return l1
}