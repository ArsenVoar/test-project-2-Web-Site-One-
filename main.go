package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Person struct {
	Name string
	Age  uint16
	Id   int16
}

/* func (p Person) getInformation() string {
	return fmt.Sprintf("Person Name is: %s, He is %d age, His Id: %d", p.Name, p.Age, p.Id)
} */

func main() {
	HandleFunc()
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	Tom := Person{Name: "Tom", Age: 28, Id: 1}
	//fmt.Fprintf(w, Tom.getInformation())

	tmpl, _ := template.ParseFiles("HTML/mainPage.html", "CSS/style1.css")
	tmpl.Execute(w, Tom)
}

func profilesPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Profiles")
}

func HandleFunc() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/Profiles", profilesPage)
	http.ListenAndServe(":8888", nil)
}
