package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	url := flag.Arg(0)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}
	filename := strings.Split(strings.Split(url, "//")[1], "/")[0]
	if err != nil {
		log.Fatal(err.Error())
	}
	err = os.Mkdir(filename, os.ModePerm)
	if err != nil {
		log.Fatal(err.Error())
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	f, err := os.Create(filename + "/index.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()
	_, err = f.WriteString(string(bytes))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
