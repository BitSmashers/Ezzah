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

func (c *ConnectionTestImpl) SaveArtists(artists []Artist)  {
	for _, a := range artists {
		//Downcast, should exist a better way to keep state call after call,... pointers
		c.SaveArtist(a)
	}
}

func (c *ConnectionTestImpl) ToString() string {
	return "InMemory : "
}

func (c *ConnectionTestImpl) FindArtist(name string) *Artist {
	for _, a := range c.artists {
		if a.Name == name {
			return &a
		}
	}
	return nil
}

func (c *ConnectionTestImpl) SaveAlbum(a Album)  {
}

func (c *ConnectionTestImpl) SaveAlbums(albums []Album)  {
}

func (c *ConnectionTestImpl) FindAlbum(title string) *Album {
	return nil
}
