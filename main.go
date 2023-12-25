package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

var (
	ClientId     string
	RedirectUri  string
	ResponseType string
	State        string
	UserEmail    string
)

func main() {

	InitRouter()
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Println("Error listening...")
	}

}

func CheckAccessForEndpoint(w http.ResponseWriter, r *http.Request) {
	rdr, _ := io.ReadAll(r.Body)
	fmt.Println("string(rdr) CheckAccessForEndpoint", string(rdr))
}
func CallThatUserUnlink(w http.ResponseWriter, r *http.Request) {
	rdr, _ := io.ReadAll(r.Body)
	fmt.Println("string(rdr) CallThatUserUnlink", string(rdr))
}
func InfoAboutUserDevices(w http.ResponseWriter, r *http.Request) {
	rdr, _ := io.ReadAll(r.Body)
	fmt.Println("string(rdr) InfoAboutUserDevices", string(rdr))
}
func InfoAboutUserDevicesState(w http.ResponseWriter, r *http.Request) {
	rdr, _ := io.ReadAll(r.Body)
	fmt.Println("string(rdr) InfoAboutUserDevicesState", string(rdr))
}
func ChangeDevicesState(w http.ResponseWriter, r *http.Request) {
	rdr, _ := io.ReadAll(r.Body)
	fmt.Println("string(rdr) ChangeDevicesState", string(rdr))
}

func InitRouter() {
	http.HandleFunc("/api/v1.0/", CheckAccessForEndpoint)
	http.HandleFunc("/api/v1.0/user/unlink", CallThatUserUnlink)
	http.HandleFunc("/api/v1.0/user/devices", InfoAboutUserDevices)
	http.HandleFunc("/api/v1.0/user/devices/query", InfoAboutUserDevicesState)
	http.HandleFunc("/v1.0/user/devices/action", ChangeDevicesState)

	http.HandleFunc("/api/first_request", GetFirstAuthValues)
	http.HandleFunc("/api/login", LoginPage)
	http.HandleFunc("/api/email", ReadEmailFromLoginPage)
	http.HandleFunc("/api/access_token", AccessToken)

}

func GetFirstAuthValues(w http.ResponseWriter, r *http.Request) {
	ClientId = r.URL.Query().Get("client_id")
	RedirectUri = r.URL.Query().Get("redirect_uri")
	ResponseType = r.URL.Query().Get("response_type")
	State = r.URL.Query().Get("state")

	http.Redirect(w, r, "https://onviz-api.ru/api/login", http.StatusFound)
}

func AccessToken(w http.ResponseWriter, r *http.Request) {
	codeForAccessToken := r.URL.Query().Get("code")
	fmt.Println("codeForAccessToken:>", codeForAccessToken)
	rdr, _ := io.ReadAll(r.Body)
	fmt.Println("string(rdr), access_token:>", string(rdr))
}

func ReadEmailFromLoginPage(w http.ResponseWriter, r *http.Request) {
	code, _ := generateRandomString()
	urlForRedirect := fmt.Sprintf("https://social.yandex.net/broker/redirect?%v&state=%v&client_id%v", code, State, ClientId)
	fmt.Println("urlForRedirect:>", urlForRedirect)
	rdr, _ := io.ReadAll(r.Body)
	if rdr != nil {
		http.Redirect(w, r, urlForRedirect, http.StatusFound)
	}
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	email := r.Form.Get("email")
	UserEmail = email
	ts, err := template.ParseFiles("login.html")
	if err != nil {
		log.Println("error:", err)
	}
	ts.Execute(w, r)
}

func generateRandomString() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
