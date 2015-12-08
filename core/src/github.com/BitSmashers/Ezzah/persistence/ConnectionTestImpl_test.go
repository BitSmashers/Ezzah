package persistence
import (
	"testing"
	. "github.com/BitSmashers/Ezzah/model"
	"github.com/stretchr/testify/assert"

)

func TestSaveNewArtist(t *testing.T) {
	cnx := CreateNewConnectionForTest()
	var artists[] Artist
	artists = make([]Artist, 3)
	artists[0] = Artist{"123", "Pierre", "Details", nil, "France"}
	artists[1] = Artist{"456", "Paul", "Details2", nil, "France"}
	artists[2] = Artist{"789", "Jack", "Details3", nil, "France"}
	cnx.SaveArtists(artists)

	for _, a := range artists {
		retrievedArtist := cnx.FindArtists(a.Name)
		for _, ra := range retrievedArtist {
			assert.Equal(t, ra.Name, &a.Name, "Wrong artist retrieved ", retrievedArtist, "instead of ", a)
		}
	}
}

func TestSaveSongs(t *testing.T) {
	cnx := CreateNewConnectionForTest()
	var songs[] Song
	songs = make([]Song, 3)
	songs[0] = Song{"123","Good good", "Pierre"}
	songs[1] = Song{"456","Water", "Pierre"}
	songs[2] = Song{"789","Sunny", "Paul"}
	cnx.SaveSong(songs[0])
	cnx.SaveSong(songs[1])
	cnx.SaveSong(songs[2])

}
/*
	Id      string `json:"id"`
	Name    string `json:"name"`
	Details string `json:"details"`
	Albums  []Album `json:"albums"`
	Country string `json:"country"`
 */