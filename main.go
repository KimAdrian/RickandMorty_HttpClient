package main

import (
	"fmt"
	"github.com/KimAdrian/RickandMorty_HttpClient/client"
	"log"
)

func main() {

	characters, err := client.FetchData()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(characters)
}
