package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{ //creating a slice named LINKS, of type STRINGS, containing the strings below  (separted by commas)
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	c := make(chan string)

	for _, link := range links { //creating a for loop that's going through the range of slice LINKS
		checkLink(link, c)
	}
	for i := 0; 1 < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) { //function named CHECKLINK, will take a LINK of type string and see if the link is responding to traffic
	_, err := http.Get(link) //we get two values back from calling this function.  The first is a struct and the second is an error message if one occured.
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}
	fmt.Println(link, "is working properly!")
	c <- "Yep it's up"

}
