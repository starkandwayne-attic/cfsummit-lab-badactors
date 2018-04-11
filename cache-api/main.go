package main

import (
	"fmt"
	"net/http"
	"regexp"
	"os"
	"io/ioutil"
	"strconv"
)

func main() {

	size := "1M"
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

	buf := make([]byte, 8192)
	for i := 0; i < 8192; i++ {
		buf[i] = byte(i % 255)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := ioutil.TempFile("", "cache")
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to create cache file: %s\n", err)
			w.WriteHeader(500)
			fmt.Fprintf(w, "cache failed\n")
			return
		}

		for i := 0; i < n / 8192; i++ {
			f.Write(buf)
		}
		f.Close()

		w.WriteHeader(200)
		fmt.Fprintf(w, "cached!\n")
	})
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
