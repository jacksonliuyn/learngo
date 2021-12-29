package dao

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"reflect"
	"testing"
)

func Test_connectEs(t *testing.T) {
	tests := []struct {
		name    string
		want    *elastic.Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := connectEs()
			if (err != nil) != tt.wantErr {
				t.Errorf("connectEs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connectEs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndexMapping(t *testing.T) {
	userMapping := `{
    "mappings":{
        "properties":{
            "name":{
                "type":"keyword"
            },
            "age":{
                "type":"byte"
            },
            "birth":{
                "type":"date"
            }
        }
    }
}`
	client, _ := connectEs()
	// 检测索引是否存在
	indexName := "go-test"
	// 创建上下文
	ctx := context.Background()
	exist, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		t.Errorf("检测索引失败:%s", err)
		return
	}
	if exist {
		t.Error("索引已经存在，无需重复创建！")
		return
	}
	res, err := client.CreateIndex(indexName).BodyString(userMapping).Do(ctx)
	if exist {
		t.Errorf("创建索引失败:%s", err)
		return
	}
	fmt.Println("创建成功:", res)
}
