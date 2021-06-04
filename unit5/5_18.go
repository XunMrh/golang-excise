package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func main5_18() {
	url := "https://www.baidu.com"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		log.Fatal(err)
	}
	n, err := io.Copy(f, resp.Body)
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
	fmt.Println(n)
}
