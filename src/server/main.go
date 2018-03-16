package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var (
	// Port 服务端口
	Port int
	// Path 文件保存目录
	Path string
)

func init() {
	flag.IntVar(&Port, "port", 8888, "server start port")
	flag.StringVar(&Path, "path", "D:\\data\\cdn\\nfs\\ppt", "file save dirtory")
}

func main() {

	addr := fmt.Sprintf(":%d", Port)
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

		// info
		var pkgID string
		pkgIds, ok := args["pkgId"]
		if !ok || 0 == len(pkgIds) {
			pkgID = "nil"
		}
		pkgID = pkgIds[0]
		fmt.Printf("Convert finish: pkgID= %v\n", pkgID)

		// save
		path, err := save(pkgID, r.Body)
		if nil != err {
			fmt.Printf("fail to save file[%s] into local", pkgID)
		}
		fmt.Printf("save file[%s] to path[%s]", pkgID, path)
	} else {

		// info
		bs, err := ioutil.ReadAll(r.Body)
		if nil != err {
			fmt.Println("fail to parse reponse body")
			return
		}
		fmt.Printf("convert fail: error= %s\n", string(bs))
	}
}

func save(pkgID string, reader io.ReadCloser) (filePath string, err error) {

	defer reader.Close()
	// build file name and file path
	filename := fmt.Sprintf("%s.%s", pkgID, "zip")
	filePath, err = filepath.Abs(filepath.Join(Path, time.Now().Format("20060102"), filename))

	// create all dir
	path := filepath.Dir(filePath)
	err = os.MkdirAll(path, os.ModeDir)
	if err != nil {
		fmt.Printf("fail to mkdir path[%v]: %v", path, err.Error())
		return
	}

	// save
	f, err := os.Create(filePath)
	if nil != err {
		fmt.Printf("fail to create file[%v]\n", filePath)
		return
	}
	defer f.Close()

	// save
	bufReader := bufio.NewReader(reader)
	_, err = bufReader.WriteTo(f)
	fmt.Printf("\tsave(%s) to [%s]\n", filename, path)
	return
}
