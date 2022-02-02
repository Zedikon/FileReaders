package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	for true {

		var path string

		fmt.Println("")
		fmt.Println("Input path to file below.")
		fmt.Print(">")
		fmt.Scanf("%s\n", &path)

		file, err := os.Open(path)

		if err != nil {
			fmt.Print("Error! File not found, or not unavailable for reading!")
			main()
		}
		defer file.Close()

		data := make([]byte, 64)

		for {
			n, err := file.Read(data)
			if err == io.EOF {
				break
			}

			fmt.Print(string(data[:n]))

		}
	}
}
