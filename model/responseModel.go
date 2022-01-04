package model

import (
	"time"
)

type ResponseStruct struct {
	Info struct {
		Count int         `json:"count"`
		Pages int         `json:"pages"`
		Next  string      `json:"next"`
		Prev  interface{} `json:"prev"`
	} `json:"info"`
	Results []struct {
		Id      int    `json:"id"`
		Name    string `json:"name"`
		Status  string `json:"status"`
		Species string `json:"species"`
		Type    string `json:"type"`
		Gender  string `json:"gender"`
		Origin  struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"origin"`
		Location struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"location"`
		Image   string    `json:"image"`
		Episode []string  `json:"episode"`
		Url     string    `json:"url"`
		Created time.Time `json:"created"`
	} `json:"results"`
}

//func (results ResponseStruct) TextOutput() string {
//	p :=fmt.Sprintf(
//		"Id: %d\nName: %s\nStatus: %s\nSpecies: %s\nType: %s\nGender: %s\nOrigin: {\n\tName: %s\n\tUrl: %s\n}\nLocation: {\n\tName: %s\n\tUrl: %s\n}\nImage: %s\nEpisode: %s\nUrl: %s\nCreated: %s\n",
//		results.Results[0].Id, results.Results[0].Name, results.Results[0].Status,
//		results.Results[0].Species, results.Results[0].Type,
//		results.Results[0].Gender, results.Results[0].Origin.Name, results.Results[0].Origin.Url,
//		results.Results[0].Location.Name, results.Results[0].Location.Url, results.Results[0].Image,
//		results.Results[0].Episode, results.Results[0].Url, results.Results[0].Created)
//
//	return p
//}
