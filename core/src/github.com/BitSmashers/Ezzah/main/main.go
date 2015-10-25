// main
package main

import (
	"os"
	"log"
//"github.com/BitSmashers/Ezzah/persistence"
	"github.com/BitSmashers/Ezzah/api"
)

var neo4jURL string

func main() {
	initEnv()
	api.StartServer()
	log.Println("Db URL : ", neo4jURL)
}

func initEnv() {
	if os.Getenv("GRAPHENEDB_URL") != "" {
		neo4jURL = os.Getenv("GRAPHENEDB_URL")
	}else {
		neo4jURL = "http://localhost:7474" // Default config provided by neo4J
	}
}


