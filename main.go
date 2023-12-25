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

	rdr, _ := io.ReadAll(r.Body)
	fmt.Println(string(rdr))
	if rdr != nil {
		http.Redirect(w, r, "https://oauth.yandex.ru/authorize?client_id=4fed8408c435482b950afeb2d6e0f3cc&client_secret=dbb4420ab51f41fc86a2dedd37d2302b", http.StatusFound)
	}

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
