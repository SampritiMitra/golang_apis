package models

import "time"

type Downloading interface{
	Download()
}
type Links struct {
	Types          string    `json:"types"`
	Urls       	[]string `json:"urls"`
}
type Response struct {
	Id          string    `json:"id"`
	Start_time  time.Time     `json:"start_time"`
	End_time 	time.Time		`json:"end_time"`
	Status		string		`json:"status"`
	Download_type string	`json:"download_type"`
	Files		map[string]string	`json:"files"`
}
type Id struct{
	Id string `json:"id"`
}
type AllLinks []Links
var Request = AllLinks{}
