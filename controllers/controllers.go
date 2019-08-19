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

	if newLink.Types=="Serial"{
		Serial(newLink)
	}else{
		Concurrent(newLink)
	}
	end_time:=time.Now()
	s:=guuid.New().String()
	M=make(map[string]models.Response)

	M[s]=models.Response{Id:s,Start_time:start_time,End_time:end_time,Status:"Successful",Download_type:newLink.Types,Files:Stat}
	id:=&models.Response{Id:s,Start_time:start_time,End_time:end_time,Status:"Successful",Download_type:newLink.Types,Files:Stat}
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
	fmt.Println("Concurrent")
	for index:=0;index<len(newLink.Urls);index+=simul{
		ch:=make(chan string)
		for i:=0;i<int(math.Min(float64(simul),float64(len(newLink.Urls))));i++{
			fmt.Println("here",i)
			path:=fmt.Sprintf("/users/sampritimitra/Desktop/file%d.jpg",index+i)
			Url:=newLink.Urls[index+i]
			go DF(path, Url,&count,ch)
		}
		for{
			y:=0
			select{
			case <-ch:
				fmt.Println("This is case ch")
				if(count==int(math.Min(float64(index+simul),float64(len(newLink.Urls))))){
					if(count==len(newLink.Urls)){
						return
					}
					y=1
					break
				}
			}
			if y==1{
				break
			}
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
	fmt.Fprint(w,M[Id.Id])
}
