package main

import(
  "html/template"
  "net/http"
  "log"
  "fmt"
  "database/sql"
  _"github.com/go-sql-driver/mysql"
  "regexp"
  "github.com/subosito/twilio"
)

var (
	AccountSid = ""
	AuthToken  = ""
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
       loginsuccess := validatecreds(r.Form["email"][0],r.Form["password"][0])
       if loginsuccess {
         fmt.Fprintf(w,"loginsuccess")
       } else {
         fmt.Fprintf(w,"failed")
       }
}

func validatecreds(email string, password string) bool {
   db := dbconnection()
   fmt.Println("Username and Password\n",email,password)

   stmt, err := db.Prepare("Select username,password from login_table where username=? and password=?")
   if err != nil {
     fmt.Printf("Error \n",err)
   }
   defer stmt.Close()

   rows, err := stmt.Query(email,password)
   if err != nil {
     fmt.Printf("Error \n",err)
   }
   defer rows.Close()

   count := 0
   for rows.Next(){
    // var dbuser string
    // var dbpassword string
    // _ = rows.Scan(&dbuser,&dbpassword)
    // fmt.Println(dbuser)
     count ++
   }

   if count == 1 {
     return true
   } else {
    return false
   }

}

func dbconnection() (*sql.DB){
  db, err := sql.Open("mysql","root:P3NT3ST3R@/twilioapp")
  if err != nil {
    fmt.Printf("Database Error \n",err)
  }
  return db
}

func sendsms(w http.ResponseWriter, r *http.Request){
  t, err := template.ParseFiles("sendsms.html")
  if err != nil{
    log.Fatal(err)
  }
  fmt.Println(t.Execute(w,nil))
}

func forwardsms(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w,"<br> Forward SMS called!")
  r.ParseForm()

  //form validation
    var validationerror = ""
    //phone number required
    if len(r.Form["number"][0]) == 0 {
      validationerror += "<p> Phone Number Required! </p>"
    }

    //phone number must be 12 digit with country code
    if len(r.Form["number"][0]) != 12 {
      validationerror += "<p> Please enter a valid phone number </p>"
    }

    //phone number should only contain number.
    if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("number")); !m {
      validationerror += "<p> Phone Number should Contain only Number</p>"
    }

    //Message required
    if len(r.Form["message"][0]) == 0 {
      validationerror += "<p> Please enter valid message </p>"
    }

    if validationerror == "" {
       //no validation error send sms
       fmt.Fprintf(w,"<br> No validation error found")
        c := twilio.NewClient(AccountSid, AuthToken, nil)
        params := twilio.MessageParams{
          Body: r.Form["message"][0],
        }
        m, response, err := c.Messages.Send("+FROMNUMBER", "+"+r.Form["number"][0], params)
        if err != nil {
          log.Fatal(m, response, err)
        }
        fmt.Fprintf(w,m.Status)
    } else {
      //validation error
        fmt.Fprintf(w,"<br>" + validationerror)
    }
}

func main(){
   http.HandleFunc("/",login)
   http.HandleFunc("/verifyuser/",verifyuser)
   http.HandleFunc("/sendsms/",sendsms)
   http.HandleFunc("/forwardsms/",forwardsms)
   err := http.ListenAndServe(":9090",nil)
   if err != nil{
     log.Fatal("ListenandServe",err)
   }
}
