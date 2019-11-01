package main

import (	
	"github.com/satori/go.uuid"
    "fmt"  
//    "time"

	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
    _"github.com/jinzhu/gorm/dialects/postgres"
)

func AllUsers(w http.ResponseWriter, r *http.Request){
	db, err := gorm.Open("postgres","user=postgres-dev password=s3cr3tp4ssw0rd dbname=dev sslmode=disable")
    if err!= nil {
        panic("Could not connect to the database")
    }
	defer db.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request){
	db, err := gorm.Open("postgres","user=postgres-dev password=s3cr3tp4ssw0rd dbname=dev sslmode=disable")
    if err!= nil {
        panic("Could not connect to the database")
    }
	defer db.Close()

	vars := mux.Vars(r)

//	accesslevel := vars["accesslevel"]
	u := uuid.NewV4()
	firstName := vars["firstName"]
	lastName := vars["lastName"]
	email := vars["email"]
	password := vars["password"]
//	dateofbirth := vars["dateofbirth"]

	db.Create(&User{ UUID: u, FirstName: firstName, LastName: lastName, Email: email, Password: password})

	fmt.Fprintf(w, "New User Successfully created")
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	db, err := gorm.Open("postgres","user=postgres-dev password=s3cr3tp4ssw0rd dbname=dev sslmode=disable")
    if err!= nil {
        panic("Could not connect to the database")
    }
	defer db.Close()

	vars := mux.Vars(r)
	uuid := vars["uuid"]
	fmt.Fprintf(w, uuid)


	var user User
	db.Where("uuid = ?", uuid).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "User Successfully Deleted")
}

func UpdateUser(w http.ResponseWriter, r *http.Request){

	db, err := gorm.Open("postgres","user=postgres-dev password=s3cr3tp4ssw0rd dbname=dev sslmode=disable")
    if err!= nil {
        panic("Could not connect to the database")
    }
	defer db.Close()

	vars := mux.Vars(r)
	uuid := vars["uuid"]	
	fmt.Fprintf(w, "Successfully updated")

	firstName := vars["firstName"]
	lastName := vars["lastName"]
	email := vars["email"]
	password := vars["password"]


	var user User	

	db.Where("uuid = ?", uuid).Find(&user)

	user.FirstName = firstName
	user.LastName = lastName
	user.Email = email
	user.Password = password

	db.Save(&user)	




}


  