package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/SampritiMitra/golang_apis/models"
	guuid "github.com/google/uuid"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"time"
)
var M map[string]models.Response
var Stat map[string]string
var count int

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
	fmt.Println(start_time)
	Stat=make(map[string]string)
	s:=guuid.New().String()
	M=make(map[string]models.Response)
	status:="Queued"
	M[s]=models.Response{Id:s,Start_time:start_time,End_time:start_time,Status:status,Download_type:newLink.Types,Files:Stat}
	id:=&models.Response{Id:s,Start_time:start_time,End_time:start_time,Status:status,Download_type:newLink.Types,Files:Stat}
	by,_:=json.Marshal(id)
	w.Write(by)
	if newLink.Types=="Serial"{
		Serial(newLink)
	}else{
		Concurrent(newLink)
	}
	end_time:=time.Now()


	M[s] = models.Response{Id: s, Start_time: start_time, End_time: end_time, Status: "Successful", Download_type: newLink.Types, Files: Stat}
	id = &models.Response{Id: s, Start_time: start_time, End_time: end_time, Status: "Successful", Download_type: newLink.Types, Files: Stat}

	by,_=json.Marshal(id)
	w.Write(by)
	fmt.Println(M[s].Start_time)
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
		count++
		path:=fmt.Sprintf("/users/sampritimitra/Desktop/file%d.jpg",index)
		Stat[Url]=path
		if err := DownloadFile(path, Url); err != nil {
			panic(err)
		}
	}
}
func DF(filepath string, url string, count *int) error {
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
	//if *count==length{
	//	*status="Successful"
	//}
	return err
}
func Concurrent(newLink models.Links){
	count=0
	simul:=5
	fmt.Println(len(newLink.Urls))
	for index:=0;index<len(newLink.Urls);index+=simul{
		for i:=0;i<int(math.Min(float64(simul),float64(len(newLink.Urls)-index)));i++{
			path:=fmt.Sprintf("/users/sampritimitra/Desktop/file%d.jpg",index+i)
			fmt.Println(path)
			Url:=newLink.Urls[index+i]
			Stat[Url]=path
			go DF(path, Url,&count)
		}
	}
}
func Status(w http.ResponseWriter, r *http.Request){
	var Id models.Id
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter proper data")
	}
	json.Unmarshal(reqBody, &Id)
	res:=M[Id.Id]
	id:=&models.Response{Id:res.Id,Start_time:res.Start_time,End_time:res.End_time,Status:res.Status,Download_type:res.Download_type,Files:res.Files}
	by,_:=json.Marshal(id)
	w.Write(by)
	fmt.Println("status",res.Start_time)
}
