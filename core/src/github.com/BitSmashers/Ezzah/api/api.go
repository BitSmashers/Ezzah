package api

import (
	"net/http"
	"os"
	"github.com/daaku/go.httpgzip"
	"log"
	"encoding/json"
	"net/url"
	"github.com/BitSmashers/Ezzah/musicbrainz"
	"github.com/gorilla/mux"
	"github.com/BitSmashers/Ezzah/utils"

	"github.com/BitSmashers/Ezzah/persistence"
)

var connection persistence.Connection = persistence.CreateNewConnection()

func StartServer() {
	port := os.Getenv("PORT")
	if (port == "") {
		port = "9000"
	}
	log.Println("Launching server on port : " + port)
	panic(http.ListenAndServe(":" + port, httpgzip.NewHandler(Handlers())))
}

// Relative paths to UI files
var uiPath string
var bowerPath string

func Handlers() (*mux.Router) {
	uiPath = "../ui/app/"
	bowerPath = uiPath + "../bower_components/"

	router := mux.NewRouter()

	router.HandleFunc("/song/{id}", songHandler)
	router.HandleFunc("/youtube/{id}", youtubeHandler)
	router.HandleFunc("/albums/{id}", artistHandler)
	router.HandleFunc("/search/{query}", searchHandler)
	router.HandleFunc("/bower_components/{path:..*}", bowerHandler)
	router.HandleFunc("/{path:..*}", uiHandler)
	router.HandleFunc("/", uiHandler)

	return router
}

func uiHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, uiPath + r.URL.Path)
}
func bowerHandler(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["path"]
	http.ServeFile(w, r, bowerPath + path)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)["query"]
	artists := musicbrainz.ArtistSearch(query)
	if len(artists) > 0 {
		connection.SaveArtists(artists)
	}
	jsonHandler(w, artists)
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)["id"]
	jsonHandler(w, musicbrainz.GetArtistAlbums(query))
}

func songHandler(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)["id"]
	jsonHandler(w, musicbrainz.GetSongDetails(query))
}

func youtubeHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	song := musicbrainz.GetSongDetails(id)

	//TODO: move this in a `youtube` package
	query := song.Artist + " - " + song.Title
	query = url.QueryEscape(query)

	url := "https://www.googleapis.com/youtube/v3/search?q=" + query + "&part=snippet&key=AIzaSyCjHL3fQcfHvny-XEnLyGJ8rrxeCtnqOew"

	var ytres YoutubeResults
	jsontext, err := utils.GetJson(url)

	err = json.Unmarshal(jsontext, &ytres)

	if err != nil {
		panic(err)
	}

	ytId := ""
	if (len(ytres.Items) > 0) {
		ytId = ytres.Items[0].Id.Id
	}
	//TODO END

	jsonHandler(w, ytId)

}


func jsonHandler(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}


//TODO: move this in a `youtube` package
type YoutubeId struct {
	Id   string `json:"videoId"`
	Kind string `json:"kind"`
}

type YoutubeResultItem struct {
	Id YoutubeId `json:"id"`
}

type YoutubeResults struct {
	Items []YoutubeResultItem `json:"items"`
}

type YoutubeLink struct {
	Id string `json:"id"`
}
//TODO END
