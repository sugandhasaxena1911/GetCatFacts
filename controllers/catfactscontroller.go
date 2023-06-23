package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func GetCatFacts() {
	log.Println("Inside GetCatFacts API ")
	var data []map[string]interface{}
	r, e := http.Get("https://cat-fact.herokuapp.com/facts")
	if e != nil {
		log.Fatalln(e)
		return
	}
	e = json.NewDecoder(r.Body).Decode(&data)
	if e != nil {
		log.Fatal(e)
		return
	}
	path := os.Getenv("APIFILEPATH")
	f, e := os.Create(fmt.Sprint(path, os.Getenv("APIFILENAME")))
	if e != nil {
		log.Println("Unable to create file : ", e)
		return
	}
	defer f.Close()

	b, e := json.Marshal(data)
	if e != nil {
		log.Println("Unable to marshal data into json : ", e)
		return
	}
	f.Write(b)
}
