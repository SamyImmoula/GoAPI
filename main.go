package main

import (
    "fmt"  
    "log"
    "net/http"
    "github.com/satori/go.uuid"
//    "time"

  
    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"
    _"github.com/jinzhu/gorm/dialects/postgres"
    
)

var db *gorm.DB
var err error

type User struct {
    UUID        uuid.UUID `gorm:"type:uuid;primary_key;"`
    //	AccessLevel int `json:"access_level"`
	FirstName string `json:"first_name" validate:"required,min=2"`
	LastName  string `json:"last_name" validate:"required,min=2"`
	Email string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
 //   DateOfBirth time.Time

}

 func InitialMigration() {
    db, err := gorm.Open("postgres","user=postgres-dev password=s3cr3tp4ssw0rd dbname=dev sslmode=disable")
    if err!= nil {
        fmt.Println(err.Error())
        panic("Failed to connect to database")
    }
    defer db.Close()

    if !db.HasTable(&User{}) {
        db.CreateTable(&User{})
    }
} 

func helloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world")
}
func HandleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/",helloWorld).Methods("GET")
    myRouter.HandleFunc("/users", AllUsers).Methods("GET")
    myRouter.HandleFunc("/user/{firstName}/{lastName}/{email}/{password}",NewUser).Methods("POST")
    myRouter.HandleFunc("/user/{uuid}",DeleteUser).Methods("DELETE")
    myRouter.HandleFunc("/user/{uuid}/{firstName}/{lastName}/{email}/{password}",UpdateUser).Methods("PUT")
    log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {

    db, err := gorm.Open("postgres","user=postgres-dev password=s3cr3tp4ssw0rd dbname=dev sslmode=disable")

    if err!= nil {
        panic(err.Error())
    }


    defer db.Close()

    database := db.DB();

    err = database.Ping()
    if err!= nil {
        panic(err.Error())
    }

    fmt.Println("Connection sucess")

    InitialMigration()

    HandleRequests()

}

