package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var (
	// ReqUrl 请求接口
	ReqUrl string
	// Files 测试文件列表
	Files string
	// Separator 文件分隔符
	Separator string
)

func init() {

	initReqUrl :=
		"http://localhost:8080/resource/chunk?file=%s&offset=%d"
	// "http://dev-live.yunkai.com/basic-api/resource/chunk?file=%s&offset=%d"

	flag.StringVar(&ReqUrl, "reqUrl", initReqUrl, "request.url")
	flag.StringVar(&Files, "fiels", "files/1.txt;files/2.txt;files/3.txt", "file.list,spit with [;]")
	flag.StringVar(&Separator, "separator", ";", "file.separator")
}

func main() {

	// init
	flag.Parse()
	files := strings.Split(Files, Separator)

	// blockChuck(files[0])
	// lineChuck(files[0])
	modelChunk(files)
}

// chunk by line
func lineChuck(path string) {

	// read
	f, err := os.Open(path)
	if nil != err {
		fmt.Printf("\topen(%v) fail: %v\n", path, err.Error())
		return
	}
	defer f.Close()

	// init
	filename := "new.txt"
	offset := 0
	cli := &http.Client{}
	bdy := &bytes.Buffer{}

	// fill
	sperator := "\r\n"
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := scanner.Text() + "\r"
		bdy.Write([]byte(fmt.Sprintf("%x%s", len(txt), sperator)))
		bdy.Write([]byte(txt + sperator))
	}
	bdy.Write([]byte(fmt.Sprintf("0%s%s", sperator, sperator)))

	// request
	req, err := http.NewRequest("POST", fmt.Sprintf(ReqUrl, filename, offset), bdy)
	fmt.Printf("send:\n%v\n", bdy.String())
	if nil != err {
		fmt.Println(err.Error())
		return
	}
	req.Header.Set("Transfer-Encoding", "chunked")
	resp, err := cli.Do(req)
	if nil != err {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	// reponse
	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(body))
}

// chunk by block
func blockChuck(path string) {

	// read
	f, err := os.Open(path)
	if nil != err {
		fmt.Printf("\topen(%v) fail: %v\n", path, err.Error())
		return
	}
	defer f.Close()

	// init
	filename := "new.txt"
	offset := 0
	buffer := make([]byte, 10)

	// fill
	sperator := "\r\n"
	for {
		n, err := f.Read(buffer)
		if err == io.EOF {
			break
		}

		bdy := &bytes.Buffer{}
		bdy.Write([]byte(fmt.Sprintf("%x%s", n, sperator)))
		bdy.Write(buffer[0:n])
		bdy.Write([]byte(sperator))

		bdy.Write([]byte(fmt.Sprintf("0%s%s", sperator, sperator)))

		// chunk
		chunk(filename, offset, bdy, "")

		// update
		offset += n
	}
}

func modelChunk(files []string) {

	if 0 == len(files) {
		fmt.Errorf("params[files] is empty\n")
		return
	}

	for _, p := range files {
		f, err := os.Open(p)
		if nil != err {
			fmt.Printf("\topen(%v) fail: %v\n", p, err.Error())
			return
		}
		defer f.Close()

		// init
		filename := "model-chunk.txt"
		offset := 0
		buffer := make([]byte, 1024)

		// fill
		sperator := "\r\n"
		for {
			n, err := f.Read(buffer)
			if err == io.EOF {
				break
			}

			bdy := &bytes.Buffer{}
			bdy.Write([]byte(fmt.Sprintf("%x%s", n, sperator)))
			bdy.Write(buffer[0:n])
			bdy.Write([]byte(sperator))
			bdy.Write([]byte(fmt.Sprintf("0%s%s", sperator, sperator)))

			// chunk
			chunk(filename, offset, bdy, "&model=0")

			// update
			offset += n
		}
	}
}

func chunk(filename string, offset int, bdy *bytes.Buffer, args ...string) {

	// init
	cli := &http.Client{}
	var arg string
	if len(args) >= 0 {
		arg = args[0]
	} else {
		arg = ""
	}

	// request
	req, err := http.NewRequest("POST", fmt.Sprintf(ReqUrl, filename, offset)+arg, bdy)
	fmt.Printf("\tsend: \n%v\n", bdy.String())
	if nil != err {
		fmt.Println(err.Error())
		return
	}
	req.Header.Set("Transfer-Encoding", "chunked")
	resp, err := cli.Do(req)
	if nil != err {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	// reponse
	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(body))
}
