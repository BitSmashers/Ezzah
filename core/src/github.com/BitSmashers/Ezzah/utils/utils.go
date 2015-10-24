package utils

import "net/http"
import (
	"io/ioutil"
	"log"
)

func GetJson(url string) ([]byte, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
func GetJsonWithRetry(url string, retry int) ([]byte, error) {
	val, err :=	GetJson(url)
	if err != nil {
		if(retry<0){
			return nil,err
		}
	return GetJsonWithRetry(url,retry-1)
	}
	log.Println("Done, retry left ",retry)
	return val,nil

}
