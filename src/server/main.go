package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var (
	port int
)

func init() {
	port = 8888
}

func main() {

	addr := fmt.Sprintf(":%d", port)
	r := mux.NewRouter()
	r.HandleFunc("/callback", CallbackHandler).Methods("GET", "POST")
	http.ListenAndServe(addr, r)
}

// CallbackHandler 回调
func CallbackHandler(w http.ResponseWriter, r *http.Request) {

	// int
	args := r.URL.Query()
	code := "1"

	// check
	codes, ok := args["code"]
	if ok && len(codes) > 0 {
		code = codes[0]
	}
	if "1" == code {
		var pkgID string
		pkgIds, ok := args["pkgId"]
		if !ok || 0 == len(pkgIds) {
			pkgID = "nil"
		}
		pkgID = pkgIds[0]
		fmt.Printf("Convert finish: pkgID= %v\n", pkgID)
	} else {
		bs, err := ioutil.ReadAll(r.Body)
		if nil != err {
			fmt.Println("fail to parse reponse body")
			return
		}
		fmt.Printf("convert fail: error= %s\n", string(bs))
	}
}
