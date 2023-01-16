package db

import (
	"os"
)

type JsonDB struct{}

func NewJsonDB() *JsonDB {
	return &JsonDB{}
}

func (db *JsonDB) Load(dbName DBName) ([]byte, error) {
	jsonData, err := os.ReadFile("data/" + dbName + ".json")
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func (db *JsonDB) Save(dbName DBName, data Data) error {
	err := os.WriteFile("data/"+dbName+".json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}
