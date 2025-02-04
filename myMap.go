package main
import ( 
  "fmt"
//  "strconv"
  )

func creatMap1(){
  // Creating map using make
  var myMap = make(map[int]string)
  myMap[1] = "Ammar"
  myMap[2] = "Arwa"
  myMap[3] = "Yasser"
  myMap[4] = "Mohammad"
  myMap[5] = "Juliana"
  myMap[6] = "Lana"
 myMap= appendingToMap(7,"Fatimah",myMap)
 printMap1(myMap)
}
// End Of creatMap1()

func creatMap2(){
  // Creating map second way
  myMap2 := map[string]string{
  "name":   "Ammar",
  "country": "Brazil",
  "Age": "44",
  }
  fmt.Printf("lenght of myMap2 is: %v\n",len(myMap2))
myMap2= appendingToMap2("Height","1.59",myMap2)
  printMap2(myMap2)  
}
// End Of creatMap2()

func appendingToMap(key int,value string,myMap map[int]string)map[int]string{
  myMap[key]=value
  return myMap
}
// End Of appendingToMap()

func appendingToMap2(key string,value string,myMap map[string]string)map[string]string{
  myMap[key]=value
  return myMap
}
// End Of appendingToMap2()

func creatSliceOfMaps(){
  var value string
  
  fmt.Println("Enter a name to add to map")
  fmt.Scan(&value)

  // First way to create a slice of maps
  mySliceOfMaps := []map[int]string{}

  
  mySliceOfMaps = addMapToSlice(mySliceOfMaps,map[int]string{len(mySliceOfMaps)+1 : value,})
fmt.Println(mySliceOfMaps)
  fmt.Println("Enter a name to add to map")
  fmt.Scan(&value)
  mySliceOfMaps = addMapToSlice(mySliceOfMaps,map[int]string{len(mySliceOfMaps)+1 : value,})
  fmt.Println(mySliceOfMaps)

  

  
  printSliceOfMaps(mySliceOfMaps)

  // Second way to create a slice of maps
  fmt.Println("Enter a second name to add to map")
  fmt.Scan(&value)
 var mySliceOfMaps2 =   make([]map[int]string,0)
  mySliceOfMaps2=addMapToSlice(mySliceOfMaps2,map[int]string{len(mySliceOfMaps2)+1 : value,})
fmt.Println(mySliceOfMaps2)
printSliceOfMaps(mySliceOfMaps2)
}
// End Of creatSliceOfMaps()

func creatSliceOfMaps2(){
  var name string
  var age string
  var country string
  fmt.Println("Enter your name")
  fmt.Scan(&name)
  fmt.Println("Enter your age")
  fmt.Scan(&age)
  fmt.Println("Enter your country")
  fmt.Scan(&country)
  // First way to create a slice of maps
  mySliceOfMaps := []map[string]string{}
  mySliceOfMaps= addMapToSlice2(mySliceOfMaps,name,age,country)
  printSliceOfMaps2(mySliceOfMaps)

  // Second way to create a slice of maps
 var mySliceOfMaps2 =   make([]map[string]string,0)
  mySliceOfMaps2=addMapToSlice2(mySliceOfMaps2,"Ammar","44","Brazil")
printSliceOfMaps2(mySliceOfMaps2)
}
// End Of creatSliceOfMaps2()

func addMapToSlice2(mySliceOfMaps []map[string]string,name string, age string, country string)[]map[string]string{
  mySliceOfMaps = append(mySliceOfMaps,map[string]string{ "Name": name, "Age": age, "Country":country,})
return mySliceOfMaps

}
// End Of addMapToSlice2()
func addMapToSlice(mySliceOfMaps []map[int]string,myMap map[int]string)[]map[int]string{
  mySliceOfMaps = append(mySliceOfMaps,myMap)
  return mySliceOfMaps
}
// End Of addMapToSlice()

func printMap1(myMap map[int]string){
  fmt.Printf("Default format of the map is:\n %v\n", myMap)
fmt.Printf("first value in the map is: %v\n",myMap[1])
fmt.Printf("The length of the map is %v\n", len(myMap))
  
}
// End Of printMap1()


func printMap2(myMap2 map[string]string){
  for i,v := range myMap2{
    fmt.Printf("Key: %v, Value: %v\n",i,v)
  }
  
}
// End Of printMap2()
func printSliceOfMaps(mySliceOfMaps []map[int]string){
  for _,v := range mySliceOfMaps{
    for k,z := range v{
     fmt.Printf("Key: %v, Value: %v\n",k,z)
      //fmt.Printf("i is: %v, v is: %v, k is: %v, z is: %v\n",i,v,k,z)
    }

  }
}
// End Of printSliceOfMaps()

func printSliceOfMaps2(mySliceOfMaps []map[string]string){
  for _,v := range mySliceOfMaps{
    for k,z := range v{
      fmt.Printf("Key: %v, Value: %v\n",k,z)
    }

  }
}
// End Of printSliceOfMaps2()

// End Of myMap.go

  