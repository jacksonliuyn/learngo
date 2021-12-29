package main

import (
	"database/sql"
	"fmt"
	_ "github.com/apache/calcite-avatica-go/v5"
	"github.com/brianvoe/gofakeit"
	"github.com/learngo/dao"
	//"github.com/learngo/genericlearn"
	"log"
	"math/rand"
	"sync"
)

func insertRainDaily() {

	stations := []string{"R111", "R112", "R113", "R811", "R211", "R222"}
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			db, error := sql.Open("avatica", "http://localhost:8765")
			if error != nil {
				fmt.Println(error)
			}
			defer db.Close()
			for j := 0; j < 100000; j++ {
				fakedate := gofakeit.Date()
				//username := gofakeit.Username()
				randomIndex := rand.Intn(len(stations))
				station := stations[randomIndex]
				_, error = db.Exec("upsert into MYDB.RAIN_DAILY values (0, ?, ?, 2.2, 2.2 , 3.5 )", station, fakedate)
				if error != nil {
					log.Fatalln(error)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
func main() {
	//println(genericlearn.Str2Number[int]("123"))
	//dao.MongoClientTest()
	//dao.PgCliTest()
	dao.TestConnectES()
}
