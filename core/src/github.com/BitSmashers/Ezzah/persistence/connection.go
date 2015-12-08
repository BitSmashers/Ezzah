package persistence

import (
	"log"
	"github.com/jmcvetta/neoism"
	"github.com/BitSmashers/Ezzah/utils"
	. "github.com/BitSmashers/Ezzah/model"
)

type Connection interface {
	//Artist
	SaveArtist(a Artist)
	SaveArtists(artists []Artist)
	FindArtists(name string) []Artist
	DeleteArtist(id string)
	//Album
	SaveAlbum(artist Artist, a Album)
	SaveAlbums(artist Artist, albums []Album)
	FindAlbum(title string) *Album
	FindAlbums(artist Artist) []Album
	//Songs
	FindSongsByArtist(artistName string) []Song
	SaveSong(song Song)
	//Clear
	ClearDbPerLabel(label string)
	//Others
	ToString() string
}

func CreateNewConnection(specialLabel... string) Connection {
	log.Println("Trying to connect database...")
	//http://your_user:your_password@neo4j.yourdomain.com/db/data/
	//	dbpath := "http://neo4j:rirh@localhost:7474/db/data"
	dbpath := "http://neo4j:rirh@localhost:7474/db/data"
	db, err := neoism.Connect(dbpath)
	utils.HandleError(err)

	c := NewConnectionNeo(dbpath, db, specialLabel...)
	return Connection(&c)

}

func CreateNewConnectionForTest() Connection {
	c := NewConnectionTest()
	return Connection(&c)
}


