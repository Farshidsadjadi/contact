package main

import (
	"net/http"
	"html/template"
	"strconv"
	"log"
)

type Contact struct {
	ID int
	FirstName string
	LastName string
	PhoneNumber int
	Age int
}


func main () {
	contacts := []Contact{
		{1,"farshid", "sadjadi", 123123123, 25},
		{2,"alireza", "Molaee", 987745, 23},
		{3,"hojat", "haghighat", 5406546, 22},
		{4,"parisa", "movahed", 687894020, 26},
	}
	lastID := 5

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		tmpl, _ := template.New("index.html").ParseFiles("templates/index.html")
		context := Context{
			contacts,
		}
		err := tmpl.Execute(w, context)
		log.Println("err execute:", err)

	})

	http.HandleFunc("/add", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			tmpl, err := template.New("add.html").ParseFiles("templates/add.html")
			log.Println("err", err)
			err = tmpl.Execute(w, nil)
			log.Println("error", err)
		} else if req.Method == "POST" {
			req.ParseForm()
			phonenumber , _ := strconv.Atoi(req.FormValue("PhoneNumber"))
			age , _ := strconv.Atoi(req.FormValue("Age"))
			contacts = append(contacts, Contact{
				ID:lastID,
				FirstName: req.FormValue("FirstName"),
				LastName: req.FormValue("LastName"),
				PhoneNumber:phonenumber,
				Age: age})
			lastID += 1
			w.Write([]byte("OK"))
		}
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			req.ParseForm()
			ID , _ := strconv.Atoi(req.FormValue("ID"))
			for i:=0 ; i < len(contacts) ; i++ {
				if ID == contacts[i].ID {
					contacts = append(contacts[:i], contacts[i+1:]...)
				}
			}

		}
	} )


	http.ListenAndServe(":8080", nil)
}


type Context struct {
	MyContact []Contact
}
