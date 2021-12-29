package dao

import (
	"fmt"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	// Setup the mgm default config
	err := mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI("mongodb://admin:123456@localhost:27017"))
	if err != nil {
		fmt.Println(err)
	}
}

type Book struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Pages            int    `json:"pages" bson:"pages"`
}

func NewBook(name string, pages int) *Book {
	return &Book{
		Name:  name,
		Pages: pages,
	}
}
func MongoClientTest() {
	book := NewBook("test", 5)
	// Make sure to pass the model by reference (to update the model's "updated_at", "created_at" and "id" fields by mgm).
	err := mgm.Coll(book).Create(book)
	if err != nil {
		fmt.Println(err)
	}
}
