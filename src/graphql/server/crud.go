package server

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
)

// CURD 增删改查操作
func CRUD() {

	initUsers()

	http.HandleFunc("/graphQL", func(w http.ResponseWriter, r *http.Request) {
		rst := exec(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(rst)
	})
	http.ListenAndServe(":8080", nil)
}

// User 用户信息
type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

var users []User

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type:        userType,
				Description: "/user/{userId}",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						for _, u := range users {
							if int(u.ID) == id {
								return u, nil
							}
						}
					}
					return nil, nil
				},
			},
			"users": &graphql.Field{
				Type:        graphql.NewList(userType),
				Description: "/user",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return users, nil
				},
			},
		},
	},
)

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type: userType,
			Description: "/user	POST",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				rand.Seed(time.Now().UnixNano())

				u := User{
					ID:   int64(rand.Intn(100000)),
					Name: p.Args["name"].(string),
				}
				users = append(users, u)

				return u, nil
			},
		},
		"update": &graphql.Field{
			Type: userType,
			Description: "/user/{userId}	PUT",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				id, _ := p.Args["id"].(int)
				name, _ := p.Args["name"].(string)
				user := User{}

				for i, u := range users {

					if int64(id) == u.ID {
						users[i].Name = name
						user = users[i]
						break
					}
				}

				return user, nil
			},
		},
		"delete": &graphql.Field{
			Type: userType,
			Description: "/user/{userId}	DELETE",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				id, _ := p.Args["id"].(int)
				user := User{}

				for i, u := range users {

					if int64(id) == u.ID {
						user = users[i]
						users = append(users[:i], users[i+1:]...)
						break
					}
				}

				return user, nil
			},
		},
	},
})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)

func exec(query string, schema graphql.Schema) *graphql.Result {

	r := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(r.Errors) > 0 {
		fmt.Printf("fail to query[%v]: %v\n", query, r.Errors)
	}
	return r
}

func initUsers() {
	u1 := User{ID: 1, Name: "User-1"}
	u2 := User{ID: 2, Name: "User-2"}
	u3 := User{ID: 3, Name: "User-3"}
	users = append(users, u1, u2, u3)
}
