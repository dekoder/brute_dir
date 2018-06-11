package main

import (
    "fmt"
	"net/http"
	"time"
	"os"
	"strings"
)
import "./file"

func httpGet(url string) {
    client := &http.Client{
		Timeout: time.Second * 20,
	  }
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Referer", url) // referer set to domain + random_path
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	// request.Header.Set("Accept-Encoding", "gzip, deflate")
	request.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3396.79 Safari/537.36")
	
    resp, err := client.Do(request)
    if err == nil {
		defer resp.Body.Close()

		if (resp.StatusCode < 400) {
			fmt.Println(url)
		}
	} else {
		//fmt.Println(err)
	}
 	
	//done <- 1
}

func scanPath(prefix string) {
	bytes, _ := file.Asset("all.txt")
	paths :=  strings.Split(string(bytes), "\n")

	for _, tail := range paths {
		httpGet(fmt.Sprintf("%s%s", strings.Trim(prefix, "/ "), tail))
	}

	fmt.Println("end")
}

func main() {
	if (len(os.Args) > 1) {
		scanPath(os.Args[1])
	} else {
		fmt.Printf("usage: ./exe http://example.com/some_dir\n")
	}

}