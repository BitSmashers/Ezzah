package api

import "testing"
import "net/http/httptest"
import "fmt"
import "net/http"
import "io/ioutil"
import . "github.com/BitSmashers/Ezzah/music"
import "encoding/json"

var (
	server *httptest.Server
)

func init() {
	server = httptest.NewServer(Handlers())
}

func TestJson(t *testing.T) {
	url := fmt.Sprintf("%s/search/x", server.URL)
	resp, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	var data Results
	jerr := json.Unmarshal(body, &data)

	if jerr != nil {
		t.Error(jerr)
	}

}
