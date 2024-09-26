package database

import (
	"io"
	"os"
	"encoding/json"
	"errors"
)


type Database map[string]interface{}

func LoadDatabase() *Database {
	// Open our jsonFile
	jsonFile, err := os.Open("database/data.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := io.ReadAll(jsonFile)

	// we initialize our Database struct
	var database Database

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'database' which we defined above
	json.Unmarshal(byteValue, &database)


	return &database

}

func (database *Database) GetItem(databaseField string, id int) (any, error) {
	
	items, ok := (*database)[databaseField].([]interface{})
	if !ok {
		return nil, errors.New("invalid database field")
	}

	for _, item := range items {
		itemMap := item.(map[string]interface{})
		if int(itemMap["id"].(float64)) == id {
			return item, nil
		}
	}

	return nil, nil

}
