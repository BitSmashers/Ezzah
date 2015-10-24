package persistence

import (
	"log"
	"github.com/jmcvetta/neoism"
)

func CreateNewConnection() Connection {
	log.Println("Trying to connect database...")
	//http://your_user:your_password@neo4j.yourdomain.com/db/data/
	dbpath := "http://neo4j:unsecure@localhost:7474/db/data"
	db, err := neoism.Connect(dbpath)

	if err != nil {
		log.Fatal(err)
	}
	return Connection(NewConnectionNeo(dbpath, db))

}

func CreateNewConnectionForTest() Connection {
	return Connection(NewConnectionTest())
}
