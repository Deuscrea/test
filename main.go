package main

import "os"

func main() {
	file, err := os.Create("myfile.txt")

	if err != nil {
		panic(err)
	}
	data := []byte("123456 xxx \n")
	file.Write(data)
}
