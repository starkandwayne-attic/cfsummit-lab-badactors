package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	for {
		go readFile()
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("listening on port " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func readFile() {
	fmt.Println("attempting to read file")
	file, err := os.Open("./file.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	time.Sleep(time.Minute * 8)
	defer file.Close()

	io.Copy(os.Stdout, file)
}
