package dao

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"time"
)

// 连接Es
func connectEs() (*elastic.Client, error) {
	return elastic.NewClient(
		// 设置Elastic服务地址
		elastic.SetURL("http://127.0.0.1:9200"),
		// 是否转换请求地址，默认为true,当等于true时 请求http://ip:port/_nodes/http，将其返回的url作为请求路径
		elastic.SetSniff(false),
		// 心跳检查,间隔时间
		elastic.SetHealthcheckInterval(time.Second*5),
		// 设置错误日志
		elastic.SetErrorLog(log.New(os.Stderr, "ES-ERROR ", log.LstdFlags)),
		// 设置info日志
		elastic.SetInfoLog(log.New(os.Stdout, "ES-INFO ", log.LstdFlags)),
	)
}

// 测试连接
func TestConnectES() {
	client, err := connectEs()
	if err != nil {
		return
	}
	// 健康检查
	do, _ := client.ClusterHealth().Index().Do(context.TODO())
	fmt.Println("健康检查:", do)
}
