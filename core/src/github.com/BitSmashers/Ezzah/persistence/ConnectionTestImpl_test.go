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
/*
	Id      string `json:"id"`
	Name    string `json:"name"`
	Details string `json:"details"`
	Albums  []Album `json:"albums"`
	Country string `json:"country"`
 */