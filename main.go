package main

import (
	"github.com/learngo/genericlearn"
	"github.com/learngo/mq"
)

func main() {
	println(genericlearn.Str2Number[int]("123"))
	//dao.MongoClientTest()
	//dao.PgCliTest()
	//dao.TestConnectES()
	mq.Producer("test", 10)
}
