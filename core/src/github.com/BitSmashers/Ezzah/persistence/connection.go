package persistence

import . "github.com/BitSmashers/Ezzah/model"

type Connection interface {
	SaveArtist(a Artist)
	SaveArtists(artists []Artist)
	FindArtists(name string) []Artist
	DeleteArtist(id string)
	SaveAlbum(a Album)
	SaveAlbums(albums []Album)
	FindAlbum(title string) *Album
	ToString() string
}


