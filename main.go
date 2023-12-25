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
	encodeValues := r.URL.Query().Encode()
	fmt.Println("encodeValues:>", encodeValues)
}
