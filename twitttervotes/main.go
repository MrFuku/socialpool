package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

func load_env() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env")
	}
}

func main() {
	load_env()
}

var db *mgo.Session

func dialdb() error {
	var err error
	log.Println("MongDBにダイヤル中： mongo")
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{"mongo"},
		Timeout:  10 * time.Second,
		Username: "root",
		Password: "example",
	}
	db, err = mgo.DialWithInfo(mongoDBDialInfo)
	return err
}

func closedb() {
	db.Close()
	log.Println("データベース接続が閉じられました")
}
