package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "currently testing stuff")
}

func leak(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "currently leaking\n")

	go func() {
		//allocated mem
		var size uint64 = 5 * 1024 * 1024
		memory := make([]uint8, size)
		fmt.Println("starting")

		//initialize memory
		for i := uint64(0); i < size; i++ {
			memory[i] = uint8(rand.Intn(100))
			time.Sleep(1 * time.Second)
		}
		//return
		time.Sleep(999 * time.Hour)
		fmt.Println("end of routine")
	}()
	fmt.Println("returning")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/leak", leak)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func main() {
	handleRequests()
}
