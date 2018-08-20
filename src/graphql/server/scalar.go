package server

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

// Scalar 分级
func Scalar() {

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"products": &graphql.Field{
					Type: graphql.NewList(ScalarProductType),
					Args: graphql.FieldConfigArgument{
						"productor": &graphql.ArgumentConfig{
							Type: ScalarProductorType,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						products := []ScalarProduct{
							ScalarProduct{
								Location:  "lu-vl",
								Productor: NewScalarProductor("lu"),
							},
							ScalarProduct{
								Location:  "yao-v2",
								Productor: NewScalarProductor("yao"),
							},
						}

						ap := fmt.Sprintf("%v", p.Args["productor"])
						if 0 != len(ap) {
							tmpProducts := []ScalarProduct{}
							for _, y := range products {
								// fmt.Println(ap, "\t", y.Productor.Name)
								if ap == y.Productor.Name {
									tmpProducts = append(tmpProducts, y)
								}
							}
							return tmpProducts, nil
						}

						return products, nil
					},
				},
			},
		}),
	})
	if nil != err {
		log.Fatalf("fail to create schema: %v", err.Error())
	}
	execScalar(schema, `query{products{location, productor}}`)
	execScalar(schema, `query{products(productor:"lu"){location, productor}}`)
	execScalar(schema, `query($productor: ScalarProductorType){products(productor: $productor){productor}}`)
}

// ScalarProductor 生产者
type ScalarProductor struct {
	Name string
}

func (p *ScalarProductor) String() string {
	return p.Name
}

// NewScalarProductor 通过文本创建生产者
func NewScalarProductor(s string) *ScalarProductor {
	return &ScalarProductor{Name: s}
}

// ScalarProduct 产品
type ScalarProduct struct {
	Location  string           `json:"location"`
	Productor *ScalarProductor `json:"productor"`
}

// ScalarProductorType 用户 明细 GraphQL 类型
var ScalarProductorType = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "ScalarProductorType",
	Description: "ScalarProductorType for 明细 object",
	Serialize: func(v interface{}) interface{} {
		switch val := v.(type) {
		case ScalarProductor:
			return val.String()
		case *ScalarProductor:
			return (*val).String()
		default:
			return nil
		}
	},
	ParseValue: func(v interface{}) interface{} {
		switch val := v.(type) {
		case string:
			return NewScalarProductor(val)
		case *string:
			return NewScalarProductor(*val)
		default:
			return nil
		}
	},
	ParseLiteral: func(v ast.Value) interface{} {
		switch val := v.(type) {
		case *ast.StringValue:
			return NewScalarProductor(val.Value)
		default:
			return nil
		}
	},
})

// ScalarProductType 用户 GraphQL 类型
var ScalarProductType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ScalarProduct",
	Fields: graphql.Fields{
		"location": &graphql.Field{
			Type: graphql.String,
		},
		"productor": &graphql.Field{
			Type: ScalarProductorType,
		},
	},
})

func execScalar(schema graphql.Schema, query string) {

	r := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
		VariableValues: map[string]interface{}{
			"productor": "yao",
		},
	})
	if len(r.Errors) > 0 {
		log.Fatalf("fail to query[%v]: %v", query, r.Errors)
	}
	b, err := json.Marshal(r)
	if nil != err {
		log.Fatalf("fail to mashal(%v): %v", r, err.Error())
	}

	fmt.Println(string(b))
}
