package main

import "fmt"

func main() {
	//  var colors map[string]string

	colors := make(map[string]string) //this will create a new map as well

	// colors := map[string]string{}
	// 	"red" : :"ff0000",    	//don't forget the comma!!
	// 	"green" : "#4bf745",	//don't forget the comma!!

	colors["white"] = "#fffff"

	delete(colors, "white")

	fmt.Println(colors)
}
