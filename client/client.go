package client

import (
	"encoding/json"
	"github.com/KimAdrian/RickandMorty_HttpClient/model"
	"io"
	"log"
	"net/http"
)

func FetchData() (string, error) {

	url := "https://rickandmortyapi.com/api/character"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(res)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	var response model.ResponseStruct
	//Decode data
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Fatal("An error has occurred: ", err)
	}

	return response.TextOutput(), err
}
