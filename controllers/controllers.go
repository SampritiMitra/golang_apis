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
var M =make(map[string]models.Response)
var Stat map[string]string
var St =make(map[string]string)
var t =make(map[string]time.Time)
var counter=make(map[string]int)

func HomeLink(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "OK")
}
func Download(w http.ResponseWriter, r *http.Request){
	var newLink models.Links
	//if r.Response.StatusCode!=201{
	//	fmt.Fprint(w,"error in url")
	//	//return
	//}
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
	St[s]="Queued"
	M[s]=models.Response{Id:s,Start_time:start_time,End_time:start_time,Status:St[s],Download_type:newLink.Types,Files:Stat}
	counter[s]=0
	if newLink.Types=="Serial"{
		Serial(newLink)
		t[s]=time.Now()
	}else{
		Concurrent(newLink,St,s)
	}
	end_time:=t[s]
	M[s]=models.Response{Id:s,Start_time:start_time,End_time:end_time,Status:St[s],Download_type:newLink.Types,Files:Stat}
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
func DF(filepath string, url string, count *int, ch chan string, s string) error {
	// Get the data
	resp, err := http.Get(url)
	*count++
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
func Concurrent(newLink models.Links, St map[string]string, s string){
	count:=0
	simul:=10
	fmt.Println(len(newLink.Urls))
	for index:=0;index<len(newLink.Urls);index+=simul{
		ch:=make(chan string,2*simul)
		for i:=0;i<int(math.Min(float64(simul),float64(len(newLink.Urls)-index)));i++{
			path:=fmt.Sprintf("/users/sampritimitra/Desktop/file%d.jpg",index+i)
			Url:=newLink.Urls[index+i]
			Stat[Url]=path
			go DF(path, Url,&count,ch,s)
		}
		go func() {
			for{
				y:=0
				select{
				case <-ch:
					if(counter[s]==int(math.Min(float64(index+simul),float64(len(newLink.Urls))))){
						if(counter[s]==len(newLink.Urls)){
							St[s]="Successful"
							t[s]=time.Now()
							fmt.Println("returning from concurr",St[s],s)
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
	temp:=M[id]
	temp.Status=St[id]
	temp.End_time=t[id]
	fmt.Println("stat",St[id],t[id])
	M[id]=temp
	//fmt.Fprint(w,M[id])
	resp:=&models.Response{Id:id,Start_time:M[id].Start_time,End_time:M[id].End_time,Status:M[id].Status,Download_type:M[id].Download_type,Files:M[id].Files}
	by,_:=json.Marshal(resp)
	w.Write(by)
}