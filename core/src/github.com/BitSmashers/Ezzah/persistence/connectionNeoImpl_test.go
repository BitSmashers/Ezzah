package persistence



import (
	"testing"
	. "github.com/BitSmashers/Ezzah/model"
//	"github.com/stretchr/testify/assert"

	"log"
	"github.com/stretchr/testify/assert"
)

func TestSaveNewArtistN(t *testing.T) {
	cnx := CreateNewConnection()
	var artists[] Artist
	artists = make([]Artist, 3)
	albums0 := []Album{Album{Id:"1", Title:"Soleil"}, Album{Id:"2", Title:"Lune"}, Album{Id:"3", Title:"Pluto"}}
	albums1 := []Album{Album{Id:"4", Title:"Pierre"}, Album{Id:"5", Title:"Feuille"}, Album{Id:"6", Title:"Ciseaux"}}
	albums2 := []Album{Album{Id:"7", Title:"BestOf"}}

	artists[0] = Artist{"123", "Pierre", "Details", albums0, "France"}
	artists[1] = Artist{"456", "Paul", "Details2", albums1, "France"}
	artists[2] = Artist{"789", "Jack", "Details3", albums2, "France"}
	cnx.SaveArtists(artists)

	//try to retrieve them
	artistsRet := cnx.FindArtists("Pierre")
	log.Println(artistsRet)
	assert.Equal(t,len(artistsRet),1, "")
}

func TestSaveNewArtistN2(t *testing.T) {
	cnx := CreateNewConnectionForTest()
	var artists[] Artist
	artists = make([]Artist, 3)
	artists[0] = Artist{"123", "Pierre", "Details", nil, "France"}
	artists[1] = Artist{"456", "Paul", "Details2", nil, "France"}
	artists[2] = Artist{"789", "Jack", "Details3", nil, "France"}
	cnx.SaveArtists(artists)
	defer cnx.DeleteArtist("123")
	defer cnx.DeleteArtist("456")
	defer cnx.DeleteArtist("789")

	for _, a := range artists {
		retrievedArtist := cnx.FindArtists(a.Name)
		for _, ra := range retrievedArtist {
			assert.Equal(t, ra.Name, &a.Name, "Wrong artist retrieved ", retrievedArtist, "instead of ", a)
		}
	}
}