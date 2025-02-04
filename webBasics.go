package main
import(
  "net/http"
  "fmt"
)
  
func webInitialze(){

http.HandleFunc("/hello",helloHandleFunc) 
http.ListenAndServe(":8080",nil)
}



func helloHandleFunc(w http.ResponseWriter, r *http.Request){
fmt.Fprintf(w,"%d: %s\n",1,"Ammar")
}