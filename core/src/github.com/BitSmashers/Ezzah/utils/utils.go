package utils

import "net/http"
import "io/ioutil"

func GetJson(url string) ([]byte) {
  res, err := http.Get(url)

  if err != nil {
    panic(err)
  }

  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    panic(err)
  }

  return body
}
