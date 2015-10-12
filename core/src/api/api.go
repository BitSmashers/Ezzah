package api

import (
	//"io"
	"net/http"
  "os"
	"github.com/daaku/go.httpgzip"
	"log"
	//"strings"
	//"io/ioutil"
  "encoding/json"
  . "music"
)

func StartServer() {
	log.Println("Setting up server handlers")

	panic(http.ListenAndServe(":" + os.Getenv("PORT"), httpgzip.NewHandler(Handlers())))
}

var uiPath string
var bowerPath string

func Handlers() (*http.ServeMux) {
  uiPath = "ui/app/"
  bowerPath = uiPath+"../"

	serveMux := http.NewServeMux()

  serveMux.HandleFunc("/search/", searchHandler)
	//log.Println(http.Dir("../ui").getAbsolutePath())

  serveMux.HandleFunc("/bower_components/", bowerHandler)
  serveMux.HandleFunc("/", uiHandler)

	//serveMux.HandleFunc("/", ezzahHandler)

  return serveMux
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
  //params := w.Vars(r)
  //name := strings.Replace(r.URL.Path, "/search/", "", 1)

	w.Header().Set("Content-Type", "application/json")
  results := Results{
    Artist{
      Name: "Artist A",
      Albums: []Album{
        Album{
          Title: "Album 1",
          Songs: []Song { Song{ Title: "song 1" } } } } } }

  //results := Album { Title: "Album 1", Songs: Songs { Song{ Title: "song 1" } } }
  //results := Song{ Title: "song 1" }

  //body := "[{\"name\":\""+name+"\"}"

	//log.Println(results)

  //res, e := json.Marshal(results)

  //songs := []Song{Song{title: "song 1"}, Song{title: "another song"}}

  //song := Song{title: "song 1"}

  json.NewEncoder(w).Encode(results)
  //str, _ := json.Marshal(results)

  //rr, ee := json.Marshal()

  //log.Println(string(str))

	//log.Println(string(res))
	//log.Println(e)

  //io.WriteString(w, string(str))
}

