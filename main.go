package main

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
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
	log.Fatal(http.ListenAndServe(":9090", nil))
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
	http.HandleFunc("/api/email", ReadEmailFromLoginPageAndRedirect)
	http.HandleFunc("/api/access_token", AccessToken)
	http.HandleFunc("/api/refresh_token", RefreshToken)

	http.HandleFunc("/api/yandex_id_token", YandexIdToken)

}

// here we get token
func YandexIdToken(w http.ResponseWriter, r *http.Request) {
	//https://onviz-api.ru/api/yandex/token
}

var AuthCode string

func ReadEmailFromLoginPageAndRedirect(w http.ResponseWriter, r *http.Request) {
	reqId := r.Header.Get("x-request-id")
	fmt.Println("reqId:>", reqId)
	http.Redirect(w, r, "https://oauth.yandex.ru/authorize?response_type=code&client_id=4fed8408c435482b950afeb2d6e0f3cc", http.StatusFound)
	body := []byte(``)
	fmt.Println("code:>", r.URL.Query().Get("code"))
	//code := r.URL.Query().Get("code")
	//AuthCode = code

	fmt.Println("state:>", r.URL.Query().Get("state"))
	fmt.Println("scope:>", r.URL.Query().Get("scope"))
	fmt.Println("expires_in:>", r.URL.Query().Get("expires_in"))

	encodedString := base64.StdEncoding.EncodeToString([]byte(`4fed8408c435482b950afeb2d6e0f3cc:dbb4420ab51f41fc86a2dedd37d2302b`))
	req, _ := http.NewRequest("POST", "https://oauth.yandex.ru/grant_type=authorization_code&code="+AuthCode+"&client_id="+"4fed8408c435482b950afeb2d6e0f3cc", bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+encodedString)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	tbs, err := io.ReadAll(resp.Body)
	fmt.Println("string(tbs)>>>", string(tbs))

	var i interface{}
	bs, _ := io.ReadAll(r.Body)
	fmt.Println("string(bs)>", string(bs))
	err = json.Unmarshal(bs, &i)
	if err != nil {
		log.Println("Err ", err.Error())
		return
	}
	log.Println("i:", i)
	log.Println("resp:", string(bs))
	defer resp.Body.Close()
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	reqId := r.Header.Get("x-request-id")
	fmt.Println("reqId:>", reqId)
	fmt.Println("code:>", r.URL.Query().Get("code"))

	AuthCode = r.URL.Query().Get("code")
	//http.Redirect(w, r, "https://oauth.yandex.ru/authorize?response_type=code&client_id=4fed8408c435482b950afeb2d6e0f3cc", http.StatusFound)
	email := r.Form.Get("email")
	UserEmail = email
	ts, err := template.ParseFiles("login.html")
	if err != nil {
		log.Println("error:", err)
	}
	ts.Execute(w, r)
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

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	codeForAccessToken := r.URL.Query().Get("code")
	fmt.Println("codeForRefreshToken:>", codeForAccessToken)
	rdr, _ := io.ReadAll(r.Body)
	fmt.Println("string(rdr), access_token:>", string(rdr))
}

func generateRandomString() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
