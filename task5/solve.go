package main

import (
	"math/rand"
	"net/http"
	"strings"
    "encoding/json"
    "io/ioutil"
    "github.com/gorilla/mux"
    "fmt"
)

var (
	shortToFull = make(map[string]string)
)

type Response struct {
    Key string `json:"key"`
}

type Request struct {
    Url string `json:"url"`
}

func getNewKey() string {
	result := make([]string, 0)
	for i := 0; i < 6; i++ {
		a := rand.Intn(36)
		if a < 10 {
			result = append(result, string('0'+a))
		} else {
			result = append(result, string('a'+a-10))
		}
	}
	ans := strings.Join(result, "")
	_, prs := shortToFull[ans]
	if prs {
		return getNewKey()
	} else {
		return ans
	}
}

func AddNewUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    var url Request
    bodyBytes, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(bodyBytes, &url)
    key := getNewKey()
    shortToFull[key] = url.Url
    var resp Response
    resp.Key = key
    j, err := json.Marshal(resp)
    if err != nil {
        panic(err)
    }
    w.WriteHeader(http.StatusOK)
    w.Write(j)
}

func GetPostedUrl(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["key"]
    w.Header().Set("Location", shortToFull[key])
    w.WriteHeader(http.StatusMovedPermanently)
    fmt.Println(shortToFull[key])
}

func main() {
	router := mux.NewRouter()

    router.HandleFunc("/", AddNewUrl).Methods("POST")
    router.HandleFunc("/{key}", GetPostedUrl)

    http.ListenAndServe(":8082", router)

}
