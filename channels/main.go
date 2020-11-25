package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://logan-wood-science.netlify.app",
	}

	// create a new channel that passes type string
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	// for i := 0; i < len(links); i++ {
	// 	go checkLink(<-c, c)
	// }

	// infinite loop, will still only progress through the loop as requests are complete
	// for {
	// 	go checkLink(<-c, c)
	// }

	// alternative, more clear loop using channel as the range
	// NOTE: you don't want to share the same variable, l that gets passed into the goroutine
	// because the child routine and main routine would end up pointing at the same memory location
	// this is why a function literal (like anonymous function) was created and we had to pass the link
	// value argument when invoking the function, and we had to define the parameter of type string
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down.")
		c <- link
		return
	}

	fmt.Println(link, " is up.")
	c <- link
}
