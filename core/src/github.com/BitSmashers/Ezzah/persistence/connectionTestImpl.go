package persistence
import . "github.com/BitSmashers/Ezzah/model"

/**
	Implements Connection
 */
type ConnectionTestImpl struct {
	artists []Artist
}

func NewConnectionTest() ConnectionTestImpl {
//	artists := new([]Artist)
	return ConnectionTestImpl{make([]Artist,10,10)}
}


func (c ConnectionTestImpl) SaveArtist(a Artist) {
	c.artists[0] = a
}

func (c ConnectionTestImpl) ToString() string {
	return ""
}

func (c ConnectionTestImpl) FindArtist(name string) Artist {
	return c.artists[0]
}
