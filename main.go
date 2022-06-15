package main

import (
	_ "github.com/apache/calcite-avatica-go/v5"
	"github.com/learngo/dao"
	"github.com/learngo/genericlearn"
)

func main() {
	println(genericlearn.Str2Number[int]("123"))
	dao.MongoClientTest()
	dao.PgCliTest()
	dao.TestConnectES()
}
