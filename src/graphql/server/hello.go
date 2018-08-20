package server

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

// Hello hello-world
func Hello() {

	// Schema
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(
				graphql.ObjectConfig{
					Name: "root",
					Fields: graphql.Fields{
						"hello": &graphql.Field{
							Type: graphql.String,
							Resolve: func(r graphql.ResolveParams) (interface{}, error) {
								return "world", nil
							},
						},
						"yao": &graphql.Field{
							Type: graphql.String,
							Resolve: func(r graphql.ResolveParams) (interface{}, error) {
								return "yqjdcyy", nil
							},
						},
					},
				},
			),
		},
	)
	if nil != err {
		log.Fatalf("fail to create schema: %v", err.Error())
	}

	// Query
	params := graphql.Params{Schema: schema, RequestString: `{yao,hello}`}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("fail to execute: %v", r.Errors)
	}
	j, _ := json.Marshal(r)
	fmt.Printf("%s", j)
}
