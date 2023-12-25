package main

import (
	"fmt"
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
	http.HandleFunc("/api/first_request", HandleAlice)
}

func HandleAlice(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())
	client_id := r.URL.Query().Get("client_id")
	redirect_uri := r.URL.Query().Get("redirect_uri")
	response_type := r.URL.Query().Get("response_type")
	state := r.URL.Query().Get("state")
	fmt.Println("client_id:>", client_id)
	fmt.Println("redirect_uri:>", redirect_uri)
	fmt.Println("response_type:>", response_type)
	fmt.Println("state:>", state)

	//encodeValues := r.URL.Query().Encode()
	//fmt.Println("encodeValues:>", encodeValues)
}
