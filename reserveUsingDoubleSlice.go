package main
import(
  "fmt"
 "strings"
)

func reserveUsingDoubleSlice(){
  var userBookings uint
  var availableTickets uint =50
  var firstName string
  var lastName string
  var email string
 var booking = [][]string{}
////////////////////////////////////////
  
  for availableTickets !=0{
   availableTickets,firstName,lastName,email=userInput(availableTickets ,firstName ,lastName ,email )
 ////////////////////////////////////////
userBookings,availableTickets= userTicketsInput(userBookings ,availableTickets )
    fmt.Printf("userBookings is %v and Available tickets is %v\n", userBookings,availableTickets)
////////////////////////////////////////
 booking= addingUser(booking,firstName, lastName,userBookings, email)
//////////////////////////////////////
  printBookings(booking)
}
  fmt.Println("Sorry, there are no more tickets available")
}
// End Of Reserve()

func userInput(availableTickets uint, firstName string,lastName string,email string)(uint,string,string,string){
  fmt.Printf("Welcome to the booking system\n There are currently %v available tickets\n",availableTickets)
    invalidInput:
    fmt.Println("Enter your first name")
    fmt.Scan(&firstName)
    fmt.Println("Enter your last name")
    fmt.Scan(&lastName)
    fmt.Println("Enter your email")
    fmt.Scan(&email)
 // fmt.Println(firstName,lastName,email)
   isEmailValid := strings.Contains(email,"@")
    isNameValid := len(firstName)>1 && len(lastName)>1
    if !(isNameValid && isEmailValid){
      fmt.Println("Invalid input, Please try again")
      goto invalidInput}
  return availableTickets,firstName,lastName,email
}
// End Of userInput Function

func userTicketsInput(userBookings uint,availableTickets uint)(uint,uint){

label:
    fmt.Println("Enter the number of tickets you want to reserve")
    fmt.Scan(&userBookings)
    for userBookings > availableTickets{
      fmt.Printf("Sorry, there are not enough tickets available. currtenly there are %v tickest\n",availableTickets)
    goto label
    } 
availableTickets = availableTickets-userBookings
  return userBookings,availableTickets
}
// End Of userTicketsInput Function
func addingUser(booking[][] string,firstName string, lastName string,userBookings uint, email string)[][]string{

booking = append(booking,[]string{firstName+" "+lastName,fmt.Sprintf("%v", userBookings),email})
 return booking 
}

func printBookings(booking [][]string){
fmt.Printf("NAME\t\t\tBOOKING\t\t\tEMAIL\n")
  for _,v := range booking{
    for _,z := range v{
     fmt.Printf("%v\t\t",z)
    }
    fmt.Println()
  }
  fmt.Println("********************") 
}
// End Of printBookings Function