package persistence

import . "github.com/BitSmashers/Ezzah/model"

type Connection interface {
	SaveArtist(a *Artist)
	FindArtist(name string) *Artist
	ToString() string
}


