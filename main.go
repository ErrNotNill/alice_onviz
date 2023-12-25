package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {

	InitRouter()
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Println("Error listening...")
	}

}

func InitRouter() {

	http.HandleFunc("/api/login", LoginPage)
	http.HandleFunc("/api/first_request", GetFirstAuthValues)
}

type FirstAuthValues struct {
	ClientId     string `json:"client_id"`
	RedirectUri  string `json:"redirect_uri"`
	ResponseType string `json:"response_type"`
	State        string `json:"state"`
}

var FirstAuthValuesInterface interface{}

func GetFirstAuthValues(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
	rdr, _ := io.ReadAll(r.Body)
	fmt.Println(string(rdr))

	firstAuthValues := FirstAuthValues{
		ClientId:     r.URL.Query().Get("client_id"),
		RedirectUri:  r.URL.Query().Get("redirect_uri"),
		ResponseType: r.URL.Query().Get("response_type"),
		State:        r.URL.Query().Get("state"),
	}
	FirstAuthValuesInterface = firstAuthValues

	//fmt.Println("firstAuthValues:", firstAuthValues)
	//encodeValues := r.URL.Query().Encode()
	//fmt.Println("encodeValues:>", encodeValues)
}

var UserEmail string

func LoginPage(w http.ResponseWriter, r *http.Request) {
	email := r.Form.Get("email")
	UserEmail = email
	ts, err := template.ParseFiles("login.html")
	if err != nil {
		log.Println("error:", err)
	}
	ts.Execute(w, r)
}
