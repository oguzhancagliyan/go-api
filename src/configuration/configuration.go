package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"yemek-sepeti-case-api/src/models"
)

type Database struct {
	Db map[string]string
	*sync.RWMutex
}

func NewDatabase() *Database {
	return &Database{Db: make(map[string]string), RWMutex: &sync.RWMutex{}}
}

type RestoreDataModel struct {
	Data []models.Data
}

func RestoreDatabase(db *Database) {
	files, err := iOReadDir("./src/backups")

	fmt.Printf("db restore started")
	if err != nil {
		panic(err.Error())
	}

	for _, file := range files {

		data, err := os.ReadFile("./src/backups/" + file)

		if err != nil {
			panic(err.Error())
		}

		var dataArray []models.Data

		err = json.Unmarshal(data, &dataArray)

		if err != nil {
			panic(err.Error())
		}

		for _, item := range dataArray {
			_, isOk := db.Db[item.Key]

			if isOk {
				fmt.Printf("Item exist do not insert. Item is %s\n", item.Key)
				continue
			}

			fmt.Printf("item added key : %s value : %s \n", item.Key, item.Value)
			db.Db[item.Key] = item.Value
		}
	}
}

// IOReadDir https://stackoverflow.com/questions/14668850/list-directory-in-go
func iOReadDir(root string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}
