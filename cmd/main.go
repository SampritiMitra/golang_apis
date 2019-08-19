package main

import (
	"encoding/json"
	"fmt"
	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type event struct {
	ID          int    `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}
type links struct {
	Types          string    `json:"types"`
	Urls       	[]string `json:"urls"`
}
type allLinks []links
var events = allLinks{}
var m map[string]links

func homeLink(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "OK")
}
func download(w http.ResponseWriter, r *http.Request){
	var newLink links
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter proper data")
	}
	json.Unmarshal(reqBody, &newLink)
	events = append(events, newLink)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newLink)

	for index,Url:=range newLink.Urls{
		path:=fmt.Sprintf("/users/sampritimitra/Desktop/file%d.jpg",index)
		if err := DownloadFile(path, Url); err != nil {
			panic(err)
		}
	}
	s:=guuid.New().String()
	m=make(map[string]links)
	m[s]=newLink
	fmt.Fprint(w,m[s])

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
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", homeLink)
	router.HandleFunc("/downloads", download).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))
}
