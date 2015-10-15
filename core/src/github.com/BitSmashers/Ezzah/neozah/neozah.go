package neozah

import (
	"log"
	"github.com/jmcvetta/neoism"
)

func InitDatabase() {
	log.Println("Trying to connect database...")
	//http://your_user:your_password@neo4j.yourdomain.com/db/data/
	db, err := neoism.Connect("http://neo4j:unsecure@localhost:7474/db/data")
	n, err := db.CreateNode(neoism.Props{"name": "Captain Kirk"})
	defer n.Delete()  // Deferred clean up
	n.AddLabel("Person") // Add a label

	res := []struct {
		// `json:` tags matches column names in query
		A string `json:"a.name"`
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
	log.Println(db, n)
	log.Println("Result => ", cq.Result) // Dont know yet how to tame results :)
}

func InitBaseDbMock() {
	//log.Println("Inserting mock data...")
}