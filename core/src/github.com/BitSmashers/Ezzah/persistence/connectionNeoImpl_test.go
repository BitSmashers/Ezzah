package persistence



import (
	"testing"
	. "github.com/BitSmashers/Ezzah/model"
//	"github.com/stretchr/testify/assert"

	"log"
	"github.com/stretchr/testify/assert"
	"github.com/BitSmashers/Ezzah/utils"
)

func TestSaveNewArtistN(t *testing.T) {
	label := "TestSaveNewArtistN"
	cnx := CreateNewConnection(label)
	defer cnx.ClearDbPerLabel(label)

	var artists[] Artist
	artists = make([]Artist, 3)
	albums0 := []Album{Album{Id:"1", Title:"Soleil"}, Album{Id:"2", Title:"Lune"}, Album{Id:"3", Title:"Pluto"}}
	albums1 := []Album{Album{Id:"4", Title:"Pierre"}, Album{Id:"5", Title:"Feuille"}, Album{Id:"6", Title:"Ciseaux"}}
	albums2 := []Album{Album{Id:"7", Title:"BestOf"}}

	artists[0] = Artist{"123", "Pierre", "Details", albums0, "France"}
	artists[1] = Artist{"456", "Paul", "Details2", albums1, "France"}
	artists[2] = Artist{"789", "Jack", "Details3", albums2, "France"}
	utils.LOG.Debug("ok")
	cnx.SaveArtists(artists)

	//try to retrieve them
	artistsRet := cnx.FindArtists("Pierre")
	utils.LOG.Debug("Artistes : ", artistsRet)
	assert.Equal(t, len(artistsRet), 1)

	//TODO Remove artist and check

}
func TestSaveFindAlbum(t *testing.T) {
	label := "TestSaveFindAlbum"
	cnx := CreateNewConnection(label)
	defer cnx.ClearDbPerLabel(label)

	var artists[] Artist
	artists = make([]Artist, 1)
	albums0 := []Album{Album{Id:"1", Title:"Soleil"}, Album{Id:"2", Title:"Lune"}, Album{Id:"3", Title:"Pluto"}}

	artists[0] = Artist{"123", "Pierre", "Details", albums0, "France"}
	cnx.SaveArtists(artists)

	//try to retrieve them
	albumsDb := cnx.FindAlbums(artists[0])
	log.Println(albumsDb)
}

func TestSaveNewArtistN2(t *testing.T) {
	label := "TestSaveNewArtistN2"
	cnx := CreateNewConnection(label)
	defer cnx.ClearDbPerLabel(label)

	var artists[] Artist
	artists = make([]Artist, 3)
	artists[0] = Artist{"123", "Pierre", "Details", nil, "France"}
	artists[1] = Artist{"456", "Paul", "Details2", nil, "France"}
	artists[2] = Artist{"789", "Jack", "Details3", nil, "France"}
	cnx.SaveArtists(artists)
	log.Println("Search...")
	cnx.FindArtists("Pierre")
}

func TestSaveSongsNeo(t *testing.T) {
	label := "TestSaveNewArtistN2"
	cnx := CreateNewConnection(label)
	defer cnx.ClearDbPerLabel(label)

	var songs[] Song
	songs = make([]Song, 3)
	songs[0] = Song{"123", "Good good", "Pastore"}
	songs[1] = Song{"456", "Water", "Pastore"}
	songs[2] = Song{"789", "Sunny", "Paul"}
	cnx.SaveSong(songs[0])
	cnx.SaveSong(songs[1])
	//	cnx.SaveSong(songs[2])
	songsFromDb := cnx.FindSongsByArtist("Pastore")

	assert.EqualValues(t, 2, len(songsFromDb))
	assert.EqualValues(t, songs[1], (songsFromDb)[1])
	assert.EqualValues(t, songs[0], (songsFromDb)[0])
}

// Check db cleaning from labels
func TestClearDbByLabel(t *testing.T) {
	label := "TestDeleteLabels"
	cnx := CreateNewConnection(label)

	s1 := Song{"678", "678SongTitle", "678Artist"}
	cnx.SaveSong(s1)
	s2 := Song{"789", "789SongTitle", "789Artist"}
	cnx.SaveSong(s2)
	s1s := cnx.FindSongsByArtist("678Artist")
	s2s := cnx.FindSongsByArtist("789Artist")

	assert.EqualValues(t, 1, len(s1s))
	assert.EqualValues(t, s1, s1s[0])

	assert.EqualValues(t, 1, len(s2s))
	assert.EqualValues(t, s2, s2s[0])

	cnx.ClearDbPerLabel(label)

	s1s = cnx.FindSongsByArtist("678Artist")
	s2s = cnx.FindSongsByArtist("789Artist")

	assert.EqualValues(t, 0, len(s1s))
	assert.EqualValues(t, 0, len(s2s))

}