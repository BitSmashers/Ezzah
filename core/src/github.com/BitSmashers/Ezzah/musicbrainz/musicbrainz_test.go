package musicbrainz

import (
	"testing"
	"log"
)

func TestSearchArtist(t *testing.T) {
	artists := ArtistSearch("Eminem")

	for _, a := range artists {
		log.Println("Artists : ", a.Name)
	}

}
