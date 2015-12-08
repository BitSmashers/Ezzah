package utils

import "net/http"
import (
	"io/ioutil"

	"time"
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
	val, err := GetJson(url)
	if err != nil {
		if (retry < 0) {
			return nil, err
		}
		//avoid cpu burning or
		time.Sleep(200 * time.Millisecond)
		return GetJsonWithRetry(url, retry - 1)
	}
	LOG.Info("Request [%s] done, retry left %d",url, retry)
	return val, nil
}

func HandleError(err error) {
	if err != nil {
		LOG.Error("Error : %s",err)
		panic(err)
	}
}

func Append(slice, data []string) []string {
	l := len(slice)
	if l + len(data) > cap(slice) {  // reallocate
		// Allocate double what's needed, for future growth.
		newSlice := make([]string, (l + len(data)) * 2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:l + len(data)]
	for i, c := range data {
		slice[l + i] = c
	}
	return slice
}
func AppendOne(elem string, slice []string) []string {
	l := len(slice)
	if l + 1 > cap(slice) {  // reallocate
		// Allocate double what's needed, for future growth.
		newSlice := make([]string, (l + 1))
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:l + 1]
	slice[l + 1] = elem
	return slice
}
