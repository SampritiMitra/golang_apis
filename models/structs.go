package models

import "time"

type Event struct {
	ID          int    `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
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
