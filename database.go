package main

import (
	"log"
	"os"

	"github.com/boltdb/bolt"
)

func openDatabase(conf Configuration) *bolt.DB {
	homeDir, _ := os.UserHomeDir()
	_, err := os.Stat(conf.DatabaseDirectory)
	if os.IsNotExist(err) {
		db, err := bolt.Open(homeDir+"/.trae.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		return db
	} else {
		db, err := bolt.Open(conf.DatabaseDirectory+"/.trae.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		return db
	}
}
