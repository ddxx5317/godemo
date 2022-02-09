package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	ID   int
	Name string
}

func (user User) String() string {
	format := `ID : %d Name: %s`
	return fmt.Sprintf(format, user.ID, user.Name)
}
func saveUser(w http.ResponseWriter, req *http.Request) {
	userId, _ := strconv.Atoi(req.FormValue("id"))
	name := req.FormValue("name")
	out := strconv.Itoa(userId) + "-" + name
	u := &User{ID: userId, Name: name}
	log.Println(out, ";", u.String())
	//fmt.Fprintf(w, u.String())
	http.Redirect(w, req, "/", http.StatusFound)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl/result.html")
	t.Execute(w, "GOGO")
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", addUser)
	http.HandleFunc("/user", saveUser)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}