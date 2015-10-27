package persistence

import (
	"log"
	"github.com/jmcvetta/neoism"
)

func CreateNewConnection() Connection {
	log.Println("Trying to connect database...")
	//http://your_user:your_password@neo4j.yourdomain.com/db/data/
	dbpath := "http://neo4j:rirh@localhost:7474/db/data"
	db, err := neoism.Connect(dbpath)

	if err != nil {
		log.Fatal(err)
	}
	c := NewConnectionNeo(dbpath, db)
	return Connection(&c)

}

func CreateNewConnectionForTest() Connection {
	c := NewConnectionTest()
	return Connection(&c)
}
