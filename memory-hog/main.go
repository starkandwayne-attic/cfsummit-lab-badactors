package main

import (
	"fmt"
	"net/http"
	"regexp"
	"os"
	"strconv"
)

func main() {

	size := "5M"
	if v := os.Getenv("LEAK_SIZE"); v != "" {
		size = v
	}

	n := 5 * 1024 * 1024
	re := regexp.MustCompile(`^(\d+)([KMG])$`)
	if m := re.FindStringSubmatch(size); m != nil {
		magnitude, err := strconv.ParseUint(m[1], 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid LEAK_SIZE '%s', aborting...\n", size)
			os.Exit(1)
		}
		n = int(magnitude)
		switch m[2] { // "K", "M", etc.
		case "K":
			n *= 1024
		case "M":
			n *= 1024 * 1024
		case "G":
			n *= 1024 * 1024 * 1024
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		go func() {
			memory := make([]uint8, n)

			for i := 0; i < len(memory); i++ {
				memory[i] = uint8(i % 255)
			}
			<-make(chan int, 0)
		}()
	})
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
