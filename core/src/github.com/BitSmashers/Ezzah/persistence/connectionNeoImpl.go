package persistence
import (
	"github.com/jmcvetta/neoism"
	. "github.com/BitSmashers/Ezzah/model"
	"github.com/BitSmashers/Ezzah/utils"
)

/**
	Implements Connection
 */
type ConnectionNeoImpl struct {
	dbpath        string
	db            *neoism.Database
	specialLabels []string //Used to tag sone data, useful to remove quickly testing data
}

var ARTIST_ALBUM_ARROW = "alb"

func NewConnectionNeo(dbpath string, db *neoism.Database, specialLabels... string) ConnectionNeoImpl {
	labels := make([]string, len(specialLabels), len(specialLabels))
	for i, l := range specialLabels {
		labels[i] = l
	}
	return ConnectionNeoImpl{dbpath, db, labels}
}

func (c ConnectionNeoImpl) SaveArtists(artists []Artist) {
	for _, a := range artists {
		c.SaveArtist(a)
	}
}

func (c ConnectionNeoImpl) SaveArtist(a Artist) {

	createNode(c, neoism.Props{"Id":a.Id, "Name":a.Name, "Country": a.Country, "Details":a.Details}, "Artist", "Person")
	c.SaveAlbums(a, a.Albums)
	//Create links
}

// Create a node into db
// Labels are used as filter into 'FindX()' functions
func createNode(c ConnectionNeoImpl, p neoism.Props, labels ...string) (*neoism.Node) {
	node, err := c.db.CreateNode(p)
	node.AddLabel(labels...)
	if err != nil {
		panic(err)
	}
	return node
}
type resStruct2 struct {
	Artist Artist `json:"artist"`
}

func (c ConnectionNeoImpl) FindArtists(name string) []Artist {
	utils.LOG.Info("Looking for artist name:" + name)
	artists := make([]Artist, 0, 100)
	// query results
	res := []Artist{}

	cq := neoism.CypherQuery{
		Statement:
		`
		MATCH (artist:Artist)
		WHERE artist.Name={name}
		RETURN artist
		`,
		Parameters: neoism.Props{"name": name},
		Result:     &res,
	}
	err := c.db.Cypher(&cq)

	utils.HandleError(err)
	utils.LOG.Info("Res len", len(res))
	utils.LOG.Info("Res", res)
	utils.LOG.Info("Result size : ", len(artists), " Result : ", artists)
	return artists
}
func (c ConnectionNeoImpl) findArtistById(id string) Artist {
	utils.LOG.Info("Looking for artist id:" + id)

	artistsDb := make([]Artist, 0, 100)

	cq := neoism.CypherQuery{
		Statement: "MATCH artist WHERE id={id} RETURN artist LIMIT 100",
		Parameters: neoism.Props{"id": id},
		Result:     &artistsDb,
	}
	err := c.db.Cypher(&cq)
	utils.HandleError(err)
	utils.LOG.Info("Result size : ", len(artistsDb), " Result : ", artistsDb)
	if len(artistsDb) != 0 {
		return artistsDb[0]
	}
	return Artist{}
	//[]Artist{"", "", "", nil, ""}
}

// Album
func (c ConnectionNeoImpl) SaveAlbums(artist Artist, albums []Album) {
	for _, a := range albums {

		c.SaveAlbum(artist, a)
	}
}

func (c ConnectionNeoImpl) SaveAlbum(artist Artist, al Album) {
	//	node := createNode(neoism.Props{"id":al.Id, "title":al.Title}, c)
	//	findArtistById("")
	//	node.Relate(ARTIST_ALBUM_ARROW, int(id), neoism.Props{})
}


func (c ConnectionNeoImpl) FindAlbum(title string) *Album {
	return nil
}

func (c ConnectionNeoImpl) FindAlbums(artist Artist) []Album {

	album := make([]Album, 0, 50)
	cq := neoism.CypherQuery{
		// Use backticks for long statements - Cypher is whitespace indifferent
		Statement: `
        MATCH (artist)-[:alb]->(album)
        WHERE artist.name={name}
        RETURN album`,
		Parameters: neoism.Props{"name": artist.Name, "rel": ARTIST_ALBUM_ARROW},
		Result:     &album,
	}
	err := c.db.Cypher(&cq)
	utils.HandleError(err)
	utils.LOG.Info("Album : ", album)
	return album
}

func (c ConnectionNeoImpl) ToString() string {
	return ""
}

func (c ConnectionNeoImpl) DeleteArtist(id string) {
}

func (c ConnectionNeoImpl) FindSongsByArtist(artistName string) []Song {
	songs := []Song{}
	cq := neoism.CypherQuery{
		//		 Use backticks for long statements - Cypher is whitespace indifferent
		Statement: `
        MATCH (song:Song)
        WHERE song.artist={name}
        RETURN song.title as title, song.artist as artist, song.id as id`,
		Parameters: neoism.Props{"name": artistName},
		Result:     &songs,
	}
	err := c.db.Cypher(&cq)
	utils.HandleError(err)
	utils.LOG.Info("Found %d songs", len(songs))
	return songs
}

func (c ConnectionNeoImpl) SaveSong(song Song) {
	utils.LOG.Info("special labels %s", c.specialLabels)
	createNode(c, neoism.Props{"id":song.Id, "artist":song.Artist, "title": song.Title},
		append(c.specialLabels, "Song")...)
}


func (c ConnectionNeoImpl) ClearDbPerLabel(label string) {
	var query string
	query = `MATCH (n:` + label + `) DELETE n`
	//	query = `MATCH (n:` + label + `) RETURN  count(*) as count`
	//	res := struct {
	//		Count string `json:"count"`
	//	}{}

	cq := neoism.CypherQuery{
		// Use backticks for long statements - Cypher is whitespace indifferent
		Statement: query,
		Parameters: neoism.Props{"label": label},
		// Result:     &res,
		Result:     nil,
	}
	err := c.db.Cypher(&cq)
	//	utils.LOG.Debug("Found %s", res)
	utils.HandleError(err)
}