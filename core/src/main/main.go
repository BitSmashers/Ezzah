// main
package main

import (
	"os"
	//"io"
	//"net/http"
	"log"
	//"strings"
	//"io/ioutil"
	"github.com/jmcvetta/neoism"
	//"github.com/daaku/go.httpgzip"
  //"encoding/json"
  "api"
)

var neo4jURL string

func main() {
	initEnv()
	//initDatabase()
	initBaseDbMock()
	api.StartServer()
	log.Println("Db URL : ", neo4jURL)
}

/*
	Handlers
*/

/*
	Init functions
		- Db
		- Mock data
		- Ezzah server (http handlers)
 */
func initEnv(){
	if os.Getenv("GRAPHENEDB_URL") != "" {
		neo4jURL = os.Getenv("GRAPHENEDB_URL")
	}else{
		neo4jURL = "http://localhost:7474" // Default config provided by neo4J
	}
}

func initDatabase() {
	log.Println("Trying to connect database...")
	//http://your_user:your_password@neo4j.yourdomain.com/db/data/
	db, err := neoism.Connect("http://neo4j:unsecure@localhost:7474/db/data")
	n, err := db.CreateNode(neoism.Props{"name": "Captain Kirk"})
	defer n.Delete()  // Deferred clean up
	n.AddLabel("Person") // Add a label

	res := []struct {
		// `json:` tags matches column names in query
		A   string `json:"a.name"`
	}{}
	cq := neoism.CypherQuery{
		// Use backticks for long statements - Cypher is whitespace indifferent
		Statement: `MATCH (a { name: "Captain Kirk"}) RETURN a LIMIT 100`,
		Parameters: neoism.Props{"name": "Dr McCoy"},
		Result:     &res,
	}
	err = db.Cypher(&cq)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(db,n)
	log.Println("Result => ", cq.Result) // Dont know yet how to tame results :)
}

func initBaseDbMock() {
	log.Println("Inserting mock data...")
}
