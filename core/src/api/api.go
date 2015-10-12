package api

import (
	//"io"
	"net/http"
  "os"
	"github.com/daaku/go.httpgzip"
	"log"
	"strings"
	//"io/ioutil"
  "encoding/json"
  . "music"
)

func StartServer() {
	log.Println("Setting up server handlers")

	log.Println(":" + os.Getenv("PORT"))
	panic(http.ListenAndServe(":" + os.Getenv("PORT"), httpgzip.NewHandler(Handlers())))
}

var uiPath string
var bowerPath string

func Handlers() (*http.ServeMux) {
  uiPath = "../ui/app/"
  bowerPath = uiPath+"../"

	serveMux := http.NewServeMux()

  serveMux.HandleFunc("/search/", searchHandler)
  serveMux.HandleFunc("/youtube/", youtubeHandler)
	//log.Println(http.Dir("../ui").getAbsolutePath())

  serveMux.HandleFunc("/bower_components/", bowerHandler)
  serveMux.HandleFunc("/", uiHandler)

	//serveMux.HandleFunc("/", ezzahHandler)

  return serveMux
}

type YoutubeLink struct {
  Id string `json:"id"`
}

func youtubeHandler(w http.ResponseWriter, r *http.Request) {
  //log.Println("path:", r.URL.Path)
  parts := strings.Split(r.URL.Path, "/")
  id := parts[len(parts) -1]
  log.Println(id)

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(YoutubeLink { "7tKVKG4jdQk" } )

}

func uiHandler(w http.ResponseWriter, r *http.Request) {
  //log.Println("path:", r.URL.Path)
  http.ServeFile(w, r, uiPath+r.URL.Path)
}

func bowerHandler(w http.ResponseWriter, r *http.Request) {
  //log.Println("BOWER:", r.URL.Path)
  http.ServeFile(w, r, bowerPath+r.URL.Path)
}


func ezzahHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("form ", r.Form, " body", r.Body, " url ", r.URL, " postVal(\"user\") ", r.PostFormValue("user"))

	w.Header().Set("Content-Type", "text/html")

	//body, _ := ioutil.ReadFile(uiPath + "index.html")
	//log.Println(body)
	//w.Write(body)
	w.Write([]byte("<h1>Ezzah</h1>"))
}

func searchHandler(w http.ResponseWriter, r *http.Request) {

  w.Header().Set("Content-Type", "application/json")

  results := Results{
    Artist{
      Id: "artista",
      Name: "Vald",
      Albums: []Album{
        Album{
          Title: "NQNT",
          Tracks: []Song { Song { Id: "nqnt-0", Title: "C'est Pour" },
          Song { Id: "nqnt-1", Title: "Par Toutatis" },
          Song { Id: "nqnt-2", Title: "Shoote un Ministre" } } },
        Album{
          Title: "NQNT 2",
          Tracks: []Song { Song { Id: "nqnt2-0", Title: "Bonjour" },
          Song { Id: "nqnt2-1",  Title: "Cartes sous l'coude" },
          Song { Id: "nqnt2-2", Title: "Selfie" } } } } } }

  json.NewEncoder(w).Encode(results)
}

