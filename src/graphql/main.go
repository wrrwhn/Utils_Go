package main

import (
	server "./server"
)

func main() {

	// Hello-world
	// server.Hello()

	// 对象支持
	// http://localhost:8080/graphQL?query={user{id,name}}
	// http://localhost:8080/graphQL?query={user{id,name}}&login=true
	// server.Context()

	// CURD
	// 新增: `http://localhost:8080/graphQL?query=mutation+_{create(name:"new-5"){id,name}}`
	// 明细: `http://localhost:8080/graphQL?query={user(id:26611){id,name}}`
	// 列表: `http://localhost:8080/graphQL?query={users{id,name}}`
	// 修改: `http://localhost:8080/graphQL?query=mutation+_{update(id:26611, name:"26611"){id,name}}`
	// 删除: `http://localhost:8080/graphQL?query=mutation+_{delete(id:26611){id,name}}`
	// server.CRUD()

	server.Scalar()
	// server.ScalarExample()
}
