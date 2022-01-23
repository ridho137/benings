package config

import (
	"benings/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func Connect(database model.DatabaseConfig) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		database.Host, database.Port, database.User, database.Password, database.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}

func ReadDatabaseConfig() model.DatabaseConfig {
	var data model.DatabaseConfig

	file, errRead := ioutil.ReadFile("config/connection.json")

	if errRead != nil {
		log.Println(errRead.Error())
	} else {
		err := json.Unmarshal(file, &data)

		if err != nil {
			log.Println(err.Error())
		} else {
			log.Println("Load config database success !!")
		}
	}
	return data
}
