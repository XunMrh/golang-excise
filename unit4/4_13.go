package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	mtitle = flag.String("t", "", "")
	apikey = flag.String("apikey", "", "")
)

const api = "https://www.omdbapi.com/"

type result struct {
	PosterURL string `json:"Poster"`
}

func main4_13() {
	flag.Parse()
	murl := api + "?t=" + url.QueryEscape(*title) + "&apikey=" + url.QueryEscape(*apikey)
	resp, err := http.Get(murl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var res result
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		log.Fatal(err)
	}
	poster, err := http.Get(res.PosterURL)
	if err != nil {
		log.Fatal(err)
	}
	defer poster.Body.Close()
	f, err := os.Create(*mtitle + ".jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err2 := io.Copy(f, poster.Body)
	if err2 != nil {
		log.Fatal(err2)
	}
}
