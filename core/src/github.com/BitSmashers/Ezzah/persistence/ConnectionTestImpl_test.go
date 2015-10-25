package persistence
import (
	"testing"
	. "github.com/BitSmashers/Ezzah/model"
)

func TestSaveNewArtist(t *testing.T) {
	cnx := CreateNewConnectionForTest()
	var artists[] Artist
	artists = make([]Artist, 3)
	artists[0] = Artist{"123", "Pierre", "Details", nil, "France"}
	artists[1] = Artist{"456", "Paul", "Details2", nil, "France"}
	artists[2] = Artist{"789", "Jack", "Details3", nil, "France"}
	cnx = cnx.SaveArtists(artists)
//	cnx = cnx.SaveArtist(artists[1])
//	cnx = cnx.SaveArtist(artists[2])

	for _, a := range artists {
		retrievedArtist := cnx.FindArtist(a.Name)
		if retrievedArtist == nil || retrievedArtist.Name != a.Name {
			t.Errorf("%s %s %s %s", "Wrong artist retrieved ", retrievedArtist, "instead of ", a)
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