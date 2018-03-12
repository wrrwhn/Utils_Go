package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"time"
)

var (
	port   int
	secret []byte
)

func init() {
	port = 12306
	secret = []byte("yqjdcyy")
}

func main() {

	addr := fmt.Sprintf(":%d", port)
	r := mux.NewRouter()
	r.HandleFunc("/new", LoginHandler).Methods("GET")
	r.Handle("/logic",
		&RouteFilter{
			filters: []HTTPFunc{jwtCheckHandler},
			hdlr:    LogicHandler,
		}).Methods("GET")
	// NewRouteFilter().AddFilter(jwtCheckHandler).Handler(LogicHandler)).Methods("GET")

	// server.start
	http.ListenAndServe(addr, r)
}

// HTTPFunc 判断请求
type HTTPFunc func(w http.ResponseWriter, r *http.Request) bool

// RouteFilter struct
type RouteFilter struct {
	filters []HTTPFunc
	hdlr    http.HandlerFunc
}

func (f *RouteFilter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, hdlr := range f.filters {
		if !hdlr(w, req) {
			return
		}
	}
	f.hdlr(w, req)
}

// LoginHandler 登录
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, YunkaiClaims{
		"name",
		"info",
		jwt.StandardClaims{
			ExpiresAt: time.Date(2018, 3, 12, 22, 0, 0, 0, time.UTC).Unix(),
			Issuer:    "yqj",
			Subject:   "lmm",
		},
	})
	fmt.Println(token)
	tokenString, err := token.SignedString(secret)

	if nil != err {
		fmt.Println(err)
		fmt.Fprintf(w, err.Error())
	} else {
		fmt.Println(tokenString)
		fmt.Fprintf(w, tokenString)
	}
}

// jwtCheckHandler JWT 服务校验
func jwtCheckHandler(w http.ResponseWriter, r *http.Request) bool {

	// check.args
	var sjwt string
	authes, ok := r.Header["Authorization"]
	if ok && len(authes) >= 1 {
		sjwt = strings.TrimPrefix(authes[0], "Bearer ")
	}

	// check.jwt
	fmt.Println(sjwt)

	token, err := jwt.Parse(sjwt, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %s", token.Header["alg"])
		}
		return secret, nil
	})
	if nil != err {

		fmt.Fprintf(w, "fail to parse jwt-string:%s", err.Error())
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if !claims.VerifyIssuer("yqj", true) {
			fmt.Fprintln(w, "issue not fit")
			return false
		}
		fmt.Printf("name=%s, info=%s, exp=%s", claims["name"], claims["info"], claims["exp"])
		return true
	}

	return false
}

// LogicHandler 逻辑处理
func LogicHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "logic")
}

// YunkaiClaims 请求消息休
type YunkaiClaims struct {
	Name string `json:"name"`
	Info string `json:"info"`
	jwt.StandardClaims
}
