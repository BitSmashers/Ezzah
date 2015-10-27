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

func (c *ConnectionNeoImpl) SaveArtists(artists []Artist) {
	for _, a := range artists {
		c.SaveArtist(a)
	}
}

func (c ConnectionNeoImpl) SaveArtist(a Artist) {
	createNode(neoism.Props{"id":a.Id, "name":a.Name, "country": a.Country, "details":a.Details}, c)
	c.SaveAlbums(a.Albums)
	//Create links
}

func createNode(p neoism.Props, c ConnectionNeoImpl) (*neoism.Node) {
	node, err := c.db.CreateNode(p)
	if err != nil {
		panic(err)
	}
	return node
}

func (c ConnectionNeoImpl) FindArtist(name string) *Artist {
	return &Artist{"", "", "", nil, ""}
}

func (c ConnectionNeoImpl) SaveAlbum(al Album) {
	createNode(neoism.Props{"id":al.Id, "title":al.Title}, c)
}

func (c *ConnectionNeoImpl) SaveAlbums(albums []Album) {
	for _, a := range albums {
		c.SaveAlbum(a)
	}
}

func (c *ConnectionNeoImpl) FindAlbum(title string) *Album {
	return nil
}

func (c *ConnectionNeoImpl) ToString() string {
	return ""
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