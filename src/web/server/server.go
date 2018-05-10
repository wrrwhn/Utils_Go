package server

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var Port int

func init() {
	port := flag.Int("port", 8888, "input the port")
	flag.Parse()
	Port = *port
	// println("port=", Port, *port)
}
func Start() {
	http.HandleFunc("/", CommonHandler(IndexHandler))
	http.HandleFunc("/view/", CommonHandler(ViewHandler))
	http.HandleFunc("/edit/", CommonHandler(EditHandler))
	http.HandleFunc("/save/", CommonHandler(SaveHandler))

	listenLink := fmt.Sprintf(":%v", Port)
	log.Printf("Application start at [%v]\n", listenLink)
	http.ListenAndServe(listenLink, nil)
}
