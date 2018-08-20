package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

// Context 对象结构查询
func Context() {

	initSchema()

	http.HandleFunc("/graphQL", handler)
	http.ListenAndServe(":8080", nil)
}

var objSchema graphql.Schema

func initSchema() {

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "query",
				Fields: graphql.Fields{
					"user": &graphql.Field{
						Type: graphql.NewObject(
							graphql.ObjectConfig{
								Name: "account",
								Fields: graphql.Fields{
									"id": &graphql.Field{
										Type: graphql.String,
									},
									"name": &graphql.Field{
										Type: graphql.String,
										// 需要校验+ 有权限时显示
										Resolve: func(p graphql.ResolveParams) (interface{}, error) {
											rootValue := p.Info.RootValue.(map[string]interface{})
											if !rootValue["auth-check"].(bool) || rootValue["login"].(bool) {
												return p.Source, nil
											}
											return nil, nil
										},
									},
								},
							}),
						Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							return p.Context.Value("user"), nil
						},
					},
				},
			},
		),
	})
	if nil != err {
		log.Fatalf("fail to create schema: %s\n", err.Error())
	}
	objSchema = schema
}

func handler(w http.ResponseWriter, r *http.Request) {
	yao := struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{1, "Yao"}
	rootObject := map[string]interface{}{
		"auth-check": true,
	}

	q := r.URL.Query().Get("query")
	l := r.URL.Query().Get("login")
	if "true" == l {
		rootObject["login"] = true
	}else{
		rootObject["login"] = false
	}

	rx := graphql.Do(graphql.Params{
		Schema:        objSchema,
		RootObject:    rootObject,
		RequestString: q,
		Context:       context.WithValue(context.Background(), "user", yao),
	})
	if len(rx.Errors) > 0 {
		log.Fatalf("fail to query[%v]: %v\n", q, rx.Errors)
	}
	json.NewEncoder(w).Encode(rx)
}
