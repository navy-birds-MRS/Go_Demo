// package filewrite
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	FileName := "example.txt"
	_, err := os.Stat("example.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
			fmt.Print("Would you like to create? -> yes or no :")
			var File_creation_input string
			fmt.Scanln(&File_creation_input)
			if File_creation_input[0] == 'y' {
				Fc, err := os.OpenFile(FileName, os.O_CREATE, 0600)
				if err != nil {
					log.Fatal(err)
				}
				Fc.Close()
				//defer file.Close()
			} else {
				return
			}
		} else {
			log.Fatal(err)
		}
	}
	file, err := os.OpenFile(FileName, os.O_APPEND, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	fmt.Fprintln(writer, "first fill we have writed succesfully 123")
}
