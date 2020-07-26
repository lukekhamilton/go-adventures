package main

import (
	"log"
	"time"

	"github.com/Kamva/mgm/v3"
	"github.com/davecgh/go-spew/spew"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DateFields struct contain `created_at` and `updated_at`
// fields that autofill on insert/update model.
type DateFields struct {
	CreatedAt time.Time `json:"created" bson:"created"`
	UpdatedAt time.Time `json:"updated" bson:"updated"`
}

//--------------------------------
// DateField methods
//--------------------------------

// Creating hook used here to set `created_at` field
// value on inserting new model into database.
func (f *DateFields) Creating() error {
	f.CreatedAt = time.Now().UTC()
	return nil
}

// Saving hook used here to set `updated_at` field value
// on create/update model.
func (f *DateFields) Saving() error {
	f.UpdatedAt = time.Now().UTC()
	return nil
}

type Book struct {
	// mgm.DefaultModel `bson:",inline"` // TODO write custom
	mgm.IDField    `bson:",inline"`
	DateFields `bson:",inline"`
	Name string `json:"name" bson:"name"`
	Pages int `json:"pages" bson:"pages"`
}

func NewBook(name string, pages int) *Book {
	return &Book{
		Name:         name,
		Pages:        pages,
	}
}


func main() {
	err := mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI("mongodb://root:Netapp01!@localhost:27017"))
	if err != nil {
		log.Println(err)
	}

	log.Println("STARTED")

	book := NewBook("Pride and Prejudice", 345)

	spew.Dump(book)

	err = mgm.Coll(book).Create(book)
	if err != nil {
		log.Println(err)
	}

	spew.Dump(book)



}
