//package controllers
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/SampritiMitra/golang_apis/models"
//	"math"
//	"time"
//	guuid "github.com/google/uuid"
//	"io"
//	"io/ioutil"
//	"net/http"
//	"os"
//	//"time"
//)
//var M map[string]models.Response
//var Stat map[string]string
//var Count int
//
//func HomeLink(w http.ResponseWriter, r *http.Request) {
//	w.WriteHeader(200)
//	fmt.Fprintf(w, "OK")
//}
////func Download(w http.ResponseWriter, r *http.Request){
////	var newLink models.Links
////	reqBody, err := ioutil.ReadAll(r.Body)
////	if err != nil {
////		fmt.Fprintf(w, "Kindly enter proper data")
////	}
////	json.Unmarshal(reqBody, &newLink)
////	models.Request = append(models.Request, newLink)
////	w.WriteHeader(http.StatusCreated)
////	start_time:=time.Now()
////	fmt.Println(start_time)
////	Stat=make(map[string]string)
////	s:=guuid.New().String()
////	M=make(map[string]models.Response)
////	status:="Queued"
////	M[s]=models.Response{Id:s,Start_time:start_time,End_time:start_time,Status:status,Download_type:newLink.Types,Files:Stat}
////	id:=&models.Response{Id:s,Start_time:start_time,End_time:start_time,Status:status,Download_type:newLink.Types,Files:Stat}
////	by,_:=json.Marshal(id)
////	w.Write(by)
////	if newLink.Types=="Serial"{
////		Serial(newLink)
////	}else{
////		Concurrent(newLink)
////	}
////	end_time:=time.Now()
////
////	M[s] = models.Response{Id: s, Start_time: start_time, End_time: end_time, Status: "Successful", Download_type: newLink.Types, Files: Stat}
////	id = &models.Response{Id: s, Start_time: start_time, End_time: end_time, Status: "Successful", Download_type: newLink.Types, Files: Stat}
////
////	by,_=json.Marshal(id)
////	w.Write(by)
////	fmt.Println(M[s].Start_time)
////}
//
//func Download(w http.ResponseWriter, r *http.Request){
//	var newLink models.Links
//	reqBody, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		fmt.Fprintf(w, "Kindly enter proper data")
//	}
//	json.Unmarshal(reqBody, &newLink)
//	models.Request = append(models.Request, newLink)
//	w.WriteHeader(http.StatusCreated)
//	start_time:=time.Now()
//	//fmt.Println(start_time)
//	Stat=make(map[string]string)
//	s:=guuid.New().String()
//	M=make(map[string]models.Response)
//	status:="Queued"
//	M[s]=models.Response{Id:s,Start_time:start_time,End_time:start_time,Status:status,Download_type:newLink.Types,Files:Stat}
//	id:=&models.Response{Id:s,Start_time:start_time,End_time:start_time,Status:status,Download_type:newLink.Types,Files:Stat}
//	by,_:=json.Marshal(id)
//	w.Write(by)
//	//var Serial_obj models.Serial
//	//var Concurrent_obj models.Concurrent
//
//	if newLink.Types=="Serial"{
//		//Serial_obj.Urls=newLink.Urls
//		//Serial_obj.Downloader()
//		Serial(newLink)
//
//	}else{
//		//Concurrent_obj.Urls=newLink.Urls
//		//Concurrent_obj.Downloader()
//		Concurrent(newLink)
//	}
//	end_time:=time.Now()
//
//	M[s] = models.Response{Id: s, Start_time: start_time, End_time: end_time, Status: "Successful", Download_type: newLink.Types, Files: Stat}
//	id = &models.Response{Id: s, Start_time: start_time, End_time: end_time, Status: "Successful", Download_type: newLink.Types, Files: Stat}
//
//	by,_=json.Marshal(id)
//	w.Write(by)
//	fmt.Println(M[s].Start_time)
//}
//func DownloadFile(filepath string, url string) error {
//	// Get the data
//	resp, err := http.Get(url)
//	if err != nil {
//		return err
//	}
//	defer resp.Body.Close()
//
//	// Create the file
//	out, err := os.Create(filepath)
//	if err != nil {
//		return err
//	}
//	defer out.Close()
//
//	// Write the body to file
//	_, err = io.Copy(out, resp.Body)
//	return err
//}
//
//func DF(filepath string, url string, Count *int) error {
//	// Get the data
//	resp, err := http.Get(url)
//	*Count++
//	if err != nil {
//		return err
//	}
//	defer resp.Body.Close()
//
//	// Create the file
//	out, err := os.Create(filepath)
//	if err != nil {
//		return err
//	}
//	defer out.Close()
//
//	// Write the body to file
//	_, err = io.Copy(out, resp.Body)
//	//if *Count==length{
//	//	*status="Successful"
//	//}
//	return err
//}
//
//func Status(w http.ResponseWriter, r *http.Request){
//	var Id models.Id
//	reqBody, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		fmt.Fprintf(w, "Kindly enter proper data")
//	}
//	json.Unmarshal(reqBody, &Id)
//	res:=M[Id.Id]
//	id:=&models.Response{Id:res.Id,Start_time:res.Start_time,End_time:res.End_time,Status:res.Status,Download_type:res.Download_type,Files:res.Files}
//	by,_:=json.Marshal(id)
//	w.Write(by)
//	fmt.Println("status",res.Start_time)
//}
//
//func Serial(newLink models.Links){
//	for index,Url:=range newLink.Urls{
//		Count++
//		path:=fmt.Sprintf("/users/sampritimitra/Desktop/file%d.jpg",index)
//		Stat[Url]=path
//		if err := DownloadFile(path, Url); err != nil {
//			panic(err)
//		}
//	}
//}
//
//func Concurrent(newLink models.Links){
//	Count=0
//	simul:=5
//	fmt.Println(len(newLink.Urls))
//	for index:=0;index<len(newLink.Urls);index+=simul{
//		for i:=0;i<int(math.Min(float64(simul),float64(len(newLink.Urls)-index)));i++{
//			path:=fmt.Sprintf("/users/sampritimitra/Desktop/file%d.jpg",index+i)
//			fmt.Println(path)
//			Url:=newLink.Urls[index+i]
//			Stat[Url]=path
//			go DF(path, Url,&Count)
//		}
//	}
//}

package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/SampritiMitra/golang_apis/models"
	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"time"
)
var M map[string]models.Response
var Stat map[string]string

func HomeLink(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "OK")
}
func Download(w http.ResponseWriter, r *http.Request){
	var newLink models.Links
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter proper data")
	}
	json.Unmarshal(reqBody, &newLink)
	models.Request = append(models.Request, newLink)
	w.WriteHeader(http.StatusCreated)
	start_time:=time.Now()
	Stat=make(map[string]string)
	s:=guuid.New().String()
	M=make(map[string]models.Response)
	M[s]=models.Response{Id:s,Start_time:start_time,End_time:start_time,Status:"Queued",Download_type:newLink.Types,Files:Stat}

	if newLink.Types=="Serial"{
		Serial(newLink)
	}else{
		Concurrent(newLink)
	}
	end_time:=time.Now()
	M[s]=models.Response{Id:s,Start_time:start_time,End_time:end_time,Status:"Successful",Download_type:newLink.Types,Files:Stat}
	id:=&models.Id{Id:s}
	by,_:=json.Marshal(id)
	w.Write(by)
}
func DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
func Serial(newLink models.Links){
	for index,Url:=range newLink.Urls{
		path:=fmt.Sprintf("/users/sampritimitra/Desktop/file%d.jpg",index)
		Stat[Url]=path
		if err := DownloadFile(path, Url); err != nil {
			panic(err)
		}
	}
}
func DF(filepath string, url string, count *int, ch chan string) error {
	// Get the data
	resp, err := http.Get(url)
	*count++
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	ch<-"done"
	return err
}
func Concurrent(newLink models.Links){
	count:=0
	simul:=5
	fmt.Println(len(newLink.Urls))
	for index:=0;index<len(newLink.Urls);index+=simul{
		ch:=make(chan string,2*simul)
		for i:=0;i<int(math.Min(float64(simul),float64(len(newLink.Urls)-index)));i++{
			path:=fmt.Sprintf("/users/sampritimitra/Desktop/file%d.jpg",index+i)
			Url:=newLink.Urls[index+i]
			Stat[Url]=path
			go DF(path, Url,&count,ch)
		}
		go func() {
			for{
				y:=0
				select{
				case <-ch:
					if(count==int(math.Min(float64(index+simul),float64(len(newLink.Urls))))){
						if(count==len(newLink.Urls)){
							return
						}
						// want to break outer if i has reached index+simul value
						//simul is number of go routines simultaneously spawning
						// like i goes from 0 to 5 and then break at 5
						// ch close was not working
						y=1
						break
					}
				}
				if y==1{
					break
				}
			}
		}()
		fmt.Println("count",count)
	}
}
func Status(w http.ResponseWriter, r *http.Request){
	id:=(mux.Vars(r)["id"])
	fmt.Fprint(w,M[id])
}