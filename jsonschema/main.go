package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/qri-io/jsonschema"
)

func main() {
	var schemaData = []byte(`{
		 "title": "Person",
      "type": "object",
      "properties": {
          "firstName": {
              "type": "string"
          },
          "lastName": {
              "type": "string"
          },
          "age": {
              "description": "Age in years",
              "type": "integer",
              "minimum": 0
          },
          "friends": {
            "type" : "array",
            "items" : { "title" : "REFERENCE", "$ref" : "#" }
          }
      },
      "required": ["firstName", "lastName"]
    }`)

	rs := &jsonschema.RootSchema{}
	if err := json.Unmarshal(schemaData, rs); err != nil {
		panic("Unmarshall schemaData: " + err.Error())
	}

	var valid = []byte(`{
		"firstName":	"George",
		"lastName":		"Michael"
	}`)

	valError, err := rs.ValidateBytes(valid)
	if err != nil {
		log.Panic("err: ", err)
	}

	if len(valError) > 0 {
		log.Panic("ValError: ", valError)
	}

	var invalid = []byte(`{
		"firstName": "Prince"
	}`)

	valError, err = rs.ValidateBytes(invalid)
	if err != nil {
		log.Panic("err: ", err)
	}

	if len(valError) > 0 {
		// log.Fatal("valError: ", valError)
		fmt.Println("errors: ", valError)
	}

	var invalid1 = []byte(`{
		"firstName": "Jay",
		"lastName": "Z",
			"friends" : [{
				"firstName": "Nas"
			}]
		}`)

	valError, err = rs.ValidateBytes(invalid1)
	if err != nil {
		log.Panicln("err: ", err)
	}

	if len(valError) > 0 {
		fmt.Println("valError: ", valError)
	}
}
