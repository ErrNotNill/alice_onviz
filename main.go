package main

import (
	"github.com/gorilla/schema"
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
	http.HandleFunc("/api/first_request", GetFirstAuthValues)
}

type FirstAuthValues struct {
	ClientId     string `json:"client_id"`
	RedirectUri  string `json:"redirect_uri"`
	ResponseType string `json:"response_type"`
	State        string `json:"state"`
}

func GetFirstAuthValues(w http.ResponseWriter, r *http.Request) {
	firstAuthValues := FirstAuthValues{
		ClientId:     r.URL.Query().Get("client_id"),
		RedirectUri:  r.URL.Query().Get("redirect_uri"),
		ResponseType: r.URL.Query().Get("response_type"),
		State:        r.URL.Query().Get("state"),
	}
	if err := schema.NewDecoder().Decode(firstAuthValues, r.Form); err != nil {
		log.Println("Error decoding filter")
	}
	//encodeValues := r.URL.Query().Encode()
	//fmt.Println("encodeValues:>", encodeValues)
}
