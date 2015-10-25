package persistence
import (
	"github.com/jmcvetta/neoism"
	. "github.com/BitSmashers/Ezzah/model"
)

/**
	Implements Connection
 */
type ConnectionNeoImpl struct {
	dbpath string
	db     *neoism.Database
}

func NewConnectionNeo(dbpath string, db *neoism.Database) ConnectionNeoImpl {
	return ConnectionNeoImpl{dbpath, db}
}

func (c ConnectionNeoImpl) SaveArtist(a Artist) Connection {
	return c
}
func (c ConnectionNeoImpl) SaveArtists(a []Artist) Connection {
	return c
}

func (c ConnectionNeoImpl) ToString() string {
	return ""
}

func (c ConnectionNeoImpl) FindArtist(name string) *Artist {
	return &Artist{"", "", "", nil, ""}
}


/*

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
 */