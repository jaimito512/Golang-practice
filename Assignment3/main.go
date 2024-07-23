package main

import (
	"fmt"
	"io"
	"os"
)

// type thefile struct{}

func main() {
	fmt.Println(os.Args[1])
	mf := os.Args[1]
	myfile, err := os.Open(mf)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)

	}

	io.Copy(os.Stdout, myfile)

	// fmt.Println(myfile)

	// reader := io.Reader(file)
	// buffer := make([]byte, 20)
	// n := []int{}
	// n, _ = reader.Read(buffer)

	// fmt.Printf("Read n={%v}, err={%v}, buffer={%v}\n", n, err, string(buffer))

	// 	data, err := os.ReadFile(myfile.txt)
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 		os.Exit(1)
	// 	}

	// 	lw := File{}

	// 	io.Copy(lw, data.Body)

}

// func (f *File) Read(b []byte) (n int, err error) {
// 	fmt.Println(string(b))
// 	fmt.Println("Just wrote this many bytes:", len(b))
// 	return len(b), nil

// }
