package musicbrainz

import "time"
import . "github.com/BitSmashers/Ezzah/music"
import "encoding/json"
import "log"
import "net/url"
import . "github.com/BitSmashers/Ezzah/utils"

var mb_url string = "http://musicbrainz-mirror.eu:5000/ws/2/" //artist/?query=artist:swift&fmt=json"
var mb_url_format string = "&fmt=json"

func artists_query_url(query string) (string) {
	return mb_url + "artist/?query=artist:" + query + mb_url_format
}

func artist_releases_url(q string) (string) {
	return mb_url + "release?artist=" + q + mb_url_format + "&inc=recordings"
}

func GetArtistAlbums(query string) ([]Album) {
	query = url.QueryEscape(query)
	log.Println("Ask for : " + artist_releases_url(query))

	jsonbytes := GetJson(artist_releases_url(query))

	var data MBReleases
	err := json.Unmarshal(jsonbytes, &data)

	if err != nil {
		panic(err)
	}

	return releasesToEzzahModels(data)
}

func GetSongDetails(id string) (Song) {
	jsonbytes := GetJson(mb_url + "recording/" + id + "?inc=artist-credits+isrcs+releases" + mb_url_format)
	var data MBRecording
	err := json.Unmarshal(jsonbytes, &data)

	if err != nil {
		panic(err)
	}

	return Song{Id: data.Id, Title: data.Title, Artist: data.ArtistCredit[0].Name }
}

func ArtistSearch(query string) ([]Artist) {
	query = url.QueryEscape(query)
	log.Println("Ask for : " + artists_query_url(query))

	jsonbytes := GetJson(artists_query_url(query))

	var data MBArtistResults
	err := json.Unmarshal(jsonbytes, &data)

	if err != nil {
		panic(err)
	}

	return toEzzahModels(data)
}

func releasesToEzzahModels(res MBReleases) ([]Album) {
	var arr = make([]Album, len(res.Releases))

	for i, el := range res.Releases {

		var tracks = make([]Song, len(el.Media[0].Tracks))

		for ii, elel := range el.Media[0].Tracks {

			tracks[ii] = Song{Id: elel.NestedId.Id, Title: elel.Title }
		}

		arr[i] = Album{
			Id: el.Id,
			Title: el.Title,
			Tracks: tracks }


	}

	return arr
}

func toEzzahModels(res MBArtistResults) ([]Artist) {
	var arr = make([]Artist, len(res.Artists))

	for i, el := range res.Artists {
		arr[i] = Artist{
			Id: el.Id,
			Name: el.Name,
			Details: el.Disambiguation,
			Albums: []Album{},
			Country: el.Country }
	}

	return arr
}

type MBReleases struct {
	Releases      []MBRelease `json:"releases"`
	ReleaseOffset int `json:"release-offset"`
	ReleaseCount  int `json:"release-count"`
}

type MBRelease struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	//Date *time.Time `json:"date"`
	Media []MBReleaseMedia `json:"media"`
}

type MBReleaseMedia struct {
	Tracks []MBTrack `json:"tracks"`
}

type NestedId struct {
	Id string `json:"id"`
}

type MBTrack struct {
	NestedId NestedId `json:"recording"`
	Title    string `json:"title"`
	Number   string `json:"number"`
}

type MBArtistCredit struct {
	Name string `json:"name"`
}

type MBRecording struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	ArtistCredit []MBArtistCredit `json:"artist-credit"`
}

type MBArtistResults struct {
	Created *time.Time `json:"created,omitempty"`
	Count   int `json:"count"`
	Offset  int `json:"offset"`
	Artists []MBArtist `json:"artists"`
}

type MBArtist struct {
	Id             string `json:"id"`
	Kind           string `json:"type"`
	Score          string `json:"score"`
	Name           string `json:"name"`
	SortName       string `json:"sort-name"`
	Country        string `json:"country"`
	Area           MBArea `json:"area"`
	Disambiguation string `json:"disambiguation"`
	Life_span      MBLifeSpan `json:"life-span"`
}
type MBLifeSpan struct {
	Ended bool `json:"ended,omitempty"`
}

type MBArea struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	SortName string `json:"sort-name"`
}
