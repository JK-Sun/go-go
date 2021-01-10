package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const fetchTimes int = 2

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// go fetch(url, ch)
		go mutilFetch(url, ch)
	}

	for i := 0; i < len(os.Args[1:])*fetchTimes; i++ {
		fmt.Println(<-ch)
	}
	// for range os.Args[1:] {
	// 	fmt.Println(<-ch)
	// }
	fmt.Printf("%.3fs elapsend\n", time.Since(start).Seconds())
}

func mutilFetch(url string, ch chan<- string) {
	for i := 0; i < fetchTimes; i++ {
		fetch(url, ch)
	}
}

func fetch(url string, ch chan<- string) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reding %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.3fs %7d %s", secs, nbytes, url)
}
