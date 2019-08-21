package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/SampritiMitra/golang_apis/models"
	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"time"
)
var ResponseMap =make(map[string]models.Response)
var UrlToPathMap map[string]string
var IdToStatusMap =make(map[string]string)
var TimerMap =make(map[string]time.Time)
var counter=make(map[string]int)

func HomeLink(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "OK")
}

func Download(w http.ResponseWriter, r *http.Request){
	var newLink models.Links
	w.Header().Set("Content-Type","application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter proper data")
	}
	json.Unmarshal(reqBody, &newLink)
	models.Request = append(models.Request, newLink)
	//w.WriteHeader(http.StatusCreated)

	start_time:=time.Now()
	UrlToPathMap =make(map[string]string)
	s:=guuid.New().String()
	IdToStatusMap[s]="Queued"
	ResponseMap[s]=models.Response{Id:s,Start_time:start_time,End_time:start_time,Status:IdToStatusMap[s],Download_type:newLink.Types,Files: UrlToPathMap}
	counter[s]=0

	if newLink.Types=="Serial"{
		Serial(newLink,s)
		TimerMap[s]=time.Now()
	}else if newLink.Types=="Concurrent"{
		Concurrent(newLink,IdToStatusMap,s)
	} else{
		w.WriteHeader(400)
		id:=&models.Error{4001,
			"unknown type of download"}
		by,_:=json.Marshal(id)
		w.Write(by)
		return
	}

	end_time:= TimerMap[s]
	ResponseMap[s]=models.Response{Id:s,Start_time:start_time,End_time:end_time,Status:IdToStatusMap[s],Download_type:newLink.Types,Files: UrlToPathMap}
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

func Serial(newLink models.Links,s string){
	for index,Url:=range newLink.Urls{
		path:=fmt.Sprintf("/users/sampritimitra/Desktop/file%d.jpg",index)
		UrlToPathMap[Url]=path
		if err := DownloadFile(path, Url); err != nil {
			panic(err)
		}
	}
	IdToStatusMap[s]="Successful"
}

func DF(filepath string, url string, ch chan string, s string) error {
	// Get the data
	resp, err := http.Get(url)
	counter[s]++
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

func Concurrent(newLink models.Links, IdToStatusMap map[string]string, s string){
	simul:=2
	fmt.Println(len(newLink.Urls))
	for index:=0;index<len(newLink.Urls);index+=simul{
		ch:=make(chan string)
		for i:=0;i<int(math.Min(float64(simul),float64(len(newLink.Urls)-index)));i++{
			path:=fmt.Sprintf("/users/sampritimitra/Desktop/file%d.jpg",index+i)
			fmt.Println(int(math.Min(float64(simul),float64(len(newLink.Urls)-index))))
			Url:=newLink.Urls[index+i]
			UrlToPathMap[Url]=path
			go DF(path, Url,ch,s)
		}
		go func() {
			for{
				y:=0
				select{
					case <-ch:
						if(counter[s]==int(math.Min(float64(index+simul),float64(len(newLink.Urls))))){
							if(counter[s]==len(newLink.Urls)){
								IdToStatusMap[s]="Successful"
								TimerMap[s]=time.Now()
								fmt.Println("returning from concurr",IdToStatusMap[s],s)
								//close(ch)
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
		//fmt.Println("count",count)
	}
}

func Status(w http.ResponseWriter, r *http.Request){
	id:=(mux.Vars(r)["id"])
	temp,ok:=ResponseMap[id]
	if !ok{
		w.WriteHeader(400)
		id:=&models.Error{4001,
			"unknown download id"}
		by,_:=json.Marshal(id)
		w.Write(by)
		return
	}
	temp.Status=IdToStatusMap[id]
	temp.End_time= TimerMap[id]
	fmt.Println("stat",IdToStatusMap[id], TimerMap[id])
	ResponseMap[id]=temp
	//fmt.Fprint(w,M[id])
	resp:=&models.Response{Id:id,Start_time:ResponseMap[id].Start_time,End_time:ResponseMap[id].End_time,Status:ResponseMap[id].Status,Download_type:ResponseMap[id].Download_type,Files:ResponseMap[id].Files}
	by,_:=json.Marshal(resp)
	w.Write(by)
}

func Files(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("browse.html")
	if err != nil {
		fmt.Fprintf(w, "Unable to load template")
	}

	// mapp_file := ResponseMap
	t.Execute(w, ResponseMap)
}