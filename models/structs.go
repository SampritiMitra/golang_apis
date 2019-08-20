//package models
//
//import (
//	"fmt"
//	"github.com/SampritiMitra/golang_apis/controllers"
//	"math"
//	"time"
//)
//
//type Downloading interface{
//	Downloader()
//}
//type Links struct {
//	Types          string    `json:"types"`
//	Urls       	[]string `json:"urls"`
//}
//type Response struct {
//	Id          string    `json:"id"`
//	Start_time  time.Time     `json:"start_time"`
//	End_time 	time.Time		`json:"end_time"`
//	Status		string		`json:"status"`
//	Download_type string	`json:"download_type"`
//	Files		map[string]string	`json:"files"`
//}
//type Id struct{
//	Id string `json:"id"`
//}
//type Serial struct{
//	Urls[] string `json:"urls"`
//}
//type Concurrent struct{
//	Urls[] string `json:"urls"`
//}
//
//type AllLinks []Links
//var Request = AllLinks{}
//
//func (newLink Serial) Downloader(){
//	for index,Url:=range newLink.Urls{
//		controllers.Count++
//		path:=fmt.Sprintf("/users/sampritimitra/Desktop/file%d.jpg",index)
//		controllers.Stat[Url]=path
//		if err := controllers.DownloadFile(path, Url); err != nil {
//			panic(err)
//		}
//	}
//}
//
//func (newLink Concurrent) Downloader(){
//	controllers.Count=0
//	simul:=5
//	fmt.Println(len(newLink.Urls))
//	for index:=0;index<len(newLink.Urls);index+=simul{
//		for i:=0;i<int(math.Min(float64(simul),float64(len(newLink.Urls)-index)));i++{
//			path:=fmt.Sprintf("/users/sampritimitra/Desktop/file%d.jpg",index+i)
//			fmt.Println(path)
//			Url:=newLink.Urls[index+i]
//			controllers.Stat[Url]=path
//			go controllers.DF(path, Url,&controllers.Count)
//		}
//	}
//}
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