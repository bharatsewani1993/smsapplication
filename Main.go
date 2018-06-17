package main

import(
  "html/template"
  "net/http"
  "log"
  "os"
  "text/template"
  "fmt"
)

func login(w http.ResponseWriter, r *http.Request){
   t,_ := template.ParseFiles("login.html","main.html")
   t.Execute(w,nil)
}

func main(){
   http.HandleFunc("/",login)
   err := http.ListenAndServe(":9090",nil)
   if err != nil{
     log.Fatal("ListenandServe",err)
   }
}
