package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/nsqio/go-nsq"
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

type pool struct {
	Options []string
}

func loadOptions() ([]string, error) {
	var options []string
	iter := db.DB("ballots").C("polls").Find(nil).Iter()
	var p pool
	for iter.Next(&p) {
		options = append(options, p.Options...)
	}
	iter.Close()
	return options, iter.Err()
}

func publishVotes(votes <-chan string) <-chan struct{} {
	stopchan := make(chan struct{}, 1)
	pub, _ := nsq.NewProducer("nsqd:4150", nsq.NewConfig())
	go func() {
		for vote := range votes {
			pub.Publish("votes", []byte(vote))
		}
		log.Println("Publisher: 停止中です")
		pub.Stop()
		log.Println("Publisher: 停止しました")
		stopchan <- struct{}{}
	}()
	return stopchan
}
