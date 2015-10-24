package musicbrainz

import "time"

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