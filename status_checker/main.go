package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	linksStr := os.Getenv("LINKS")
	linksStr = strings.TrimSpace(linksStr)
	links := strings.Split(linksStr, ",")

	sleep, err := strconv.Atoi(os.Getenv("STATUS_CHECKER_INTERVAL"))
	if err != nil {
		fmt.Println("ERROR LOADING INTERVAL")
		return
	}

	c := make(chan string)

	for _, link := range links {
		go getWebsite(link, c)
	}
	for l := range c {
		go func(l string) {
			time.Sleep(time.Duration(sleep) * time.Second)
			getWebsite(l, c)
		}(l)
	}
}

func getWebsite(l string, c chan string) {
	url := formatURL(l)
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(l + " is down!")
		c <- l
		return
	} else {
		fmt.Println(l + " is up!")
		c <- l
	}
}

func formatURL(link string) string {
	link = strings.TrimSpace(link)
	if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		link = "http://" + link
	}
	return link
}
