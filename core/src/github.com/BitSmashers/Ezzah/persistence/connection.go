package persistence

import . "github.com/BitSmashers/Ezzah/model"

type Connection interface {
	SaveArtist(a Artist) Connection
	SaveArtists(artists []Artist) Connection
	FindArtist(name string) *Artist
	ToString() string
}


