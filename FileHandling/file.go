package file

import (
	"bufio"
	"fmt"
	"os"
)

func FileCreation(filename string) error {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
			fmt.Print("Would you like to create? -> yes or no :")
			var createInput string
			fmt.Scanln(&createInput)
			if createInput[0] == 'y' {
				file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, 0600)
				if err != nil {
					return err
				}
				defer file.Close()
			} else {
				return nil
			}
		} else {
			return err
		}
	}

	file, err := os.OpenFile(filename, os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	fmt.Fprintln(writer, "first fill we have writed successfully 123")
	return nil
}
