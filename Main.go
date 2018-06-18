package main

import(
  "html/template"
  "net/http"
  "log"
  "fmt"
  "database/sql"
  _"github.com/go-sql-driver/mysql"
)

func login(w http.ResponseWriter, r *http.Request){
   t,err := template.ParseFiles("login.html")
   if err != nil{
     fmt.Println("Error: ",err)
   }
   fmt.Println(t.Execute(w,nil))
}

func verifyuser(w http.ResponseWriter, r *http.Request){
  r.ParseForm()
  fmt.Println("Form Elements ",r.Form)
  //form validation
       var validationerror []string
      //check for required fields
      if len(r.Form["email"][0]) == 0 {
        validationerror = append(validationerror,"Email Id is required")
        fmt.Println("Username is required")  //print on console
        fmt.Fprintf(w,"failed") //print on client machine
        return
      }
      if len(r.Form["password"][0]) == 0{
        validationerror = append(validationerror,"Password required")
        fmt.Println("Password is required") //Print on console
        fmt.Fprintf(w,"failed") //print on client machine
        return
      }

      //validate username and password from database
        validatecreds(r.Form["email"][0],r.Form["password"][0])
}

func validatecreds(email string, password string) bool {
   db := dbconnection()
   stmt, err := db.Prepare("Select * from login_table where username=? and password=?")
   if err != nil {
     fmt.Printf("Error \n",err)
   }

   result, err := stmt.Exec(email,password)
   if err != nil{
     fmt.Printf("Error \n", err)
   }

   effect,err := result.RowsAffected()
   if err != nil {
     fmt.Println("Error \n",err)
   }
   fmt.Println("Effect ", effect)
   return true
}

func dbconnection() (*sql.DB){
  db, err := sql.Open("mysql","root:P3NT3ST3R@/twilioapp")
  if err != nil {
    fmt.Printf("Database Error \n",err)
  }
  return db
}

func main(){
   http.HandleFunc("/",login)
   http.HandleFunc("/verifyuser/",verifyuser)
   err := http.ListenAndServe(":9090",nil)
   if err != nil{
     log.Fatal("ListenandServe",err)
   }
}
