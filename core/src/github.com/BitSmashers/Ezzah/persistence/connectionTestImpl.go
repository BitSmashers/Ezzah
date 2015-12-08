package persistence
import (
	. "github.com/BitSmashers/Ezzah/model"
	"log"
)

/**
	Implements Connection
 */
type ConnectionTestImpl struct {
	artists []Artist
}

var limit = 50

func NewConnectionTest() ConnectionTestImpl {
	return ConnectionTestImpl{make([]Artist, 0, 10)}
}

func addArtist(slice []Artist, element Artist) []Artist {
	if (len(slice) == cap(slice)) {
		log.Println("Slice is full : ", cap(slice))
	}else {
		n := len(slice)
		slice = slice[0 : n + 1]
		slice[n] = element
	}
	return slice

}

func (c *ConnectionTestImpl) SaveArtist(a Artist) {
	c.artists = addArtist(c.artists, a)
}

func (c *ConnectionTestImpl) SaveArtists(artists []Artist) {
	for _, a := range artists {
		//Downcast, should exist a better way to keep state call after call,... pointers
		c.SaveArtist(a)
	}
}

func (c *ConnectionTestImpl) ToString() string {
	return "InMemory : "
}

func (c *ConnectionTestImpl) FindArtists(name string) []Artist {
	result := make([]Artist, limit, limit)
	idx := 0
	for _, a := range c.artists {
		if a.Name == name {
			result[idx] = a
			idx++
		}
	}
	return result[0:idx]
}

func (c *ConnectionTestImpl) SaveAlbum(artist Artist, a Album) {
}

func (c *ConnectionTestImpl) SaveAlbums(artist Artist, albums []Album) {
}

func (c *ConnectionTestImpl) FindAlbum(title string) *Album {
	return nil
}

func (c *ConnectionTestImpl) FindAlbums(artist Artist) []Album {
	return nil
}

func (c *ConnectionTestImpl) DeleteArtist(id string) {
}

func (c *ConnectionTestImpl)  FindSongsByArtist(artistName string) []Song {

return nil
}
func (c *ConnectionTestImpl)  SaveSong(song Song) {
}
func (c *ConnectionTestImpl) ClearDbPerLabel(label string) {
}

