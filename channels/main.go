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
	}

	c := make(chan string)

	for _, link := range links {
		// running checkLinks for each link serially
		// checkLinkWithoutConcurrency(link)

		// running checkLinks for each link via new goroutines
		// go checkLinkWithConcurrency(link, c)

		go checkLinkWithConcurrencyWithRepeatedGoRoutines(link, c)
	}

	// use a loop to get the number of messages equal to the number of requests
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c) // blocking channel -> main routine is paused until a message comes
	// }

	// infinite loop for regular status checker
	// for {
	// 	go checkLinkWithConcurrencyWithRepeatedGoRoutines(<-c, c)
	// }

	//  regular status checker for values received form channel (alternative to above)
	// for l := range c {
	// 	go checkLinkWithConcurrencyWithRepeatedGoRoutines(l, c)
	// }

	//  regular status checker after some sleep time using function liertals (anonymous function)
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second) // sleep the current routine for a period of time
			go checkLinkWithConcurrencyWithRepeatedGoRoutines(link, c)
		}(l)
	}

}

// checkLink without concurrency
func checkLinkWithoutConcurrency(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%v might be down\n", link)
		return
	}

	fmt.Printf("%v is up\n", link)
}

// checkLink with concurrency
func checkLinkWithConcurrency(link string, c chan string) {
	_, err := http.Get(link) // blocking call -> spins up another goroutine
	if err != nil {
		fmt.Printf("%v might be down\n", link)
		c <- "Might be down I think"
		return
	}

	fmt.Printf("%v is up\n", link)
	c <- "Yup it's up"
}

// checkLink with concurrency for repeated goroutines
func checkLinkWithConcurrencyWithRepeatedGoRoutines(link string, c chan string) {
	_, err := http.Get(link) // blocking call -> spins up another goroutine
	if err != nil {
		fmt.Printf("%v might be down\n", link)
		c <- link // send back the link to the channel to continously check them
		return
	}

	fmt.Printf("%v is up\n", link)
	c <- link // send back the link to the channel to continously check them
}
