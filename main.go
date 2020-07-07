package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
type HomePageSize struct {
	URL string
	Size int
}
func main() {
	urls := []string{
		"http://www.apple.com",
		"http://www.amazon.com",
		"http://www.google.com",
		"http://www.microsoft.com",
		"https://github.com",
		"https://www.linkedin.com",
	}
	results := make(chan HomePageSize)
	
	for _, url := range urls {
		go func (url string){
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()
			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}
			results <- HomePageSize{
				URL: url,
				Size: len(bs),
			}
		}(url)
	}
	var big HomePageSize
	for range urls {
		result := <- results
		fmt.Println(result.URL,result.Size)
		if result.Size > big.Size{
			big = result
			
		}
	}
	fmt.Println("winner ---> ",big.URL)
}



