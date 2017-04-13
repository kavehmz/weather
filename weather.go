package weather

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// type Weather

// Read will fetch content of API server.
func Read(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if content == nil {
		return content, errors.New("emtpy content")
	}
	return content, nil
}
