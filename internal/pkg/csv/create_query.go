package csv

import (
	"fmt"
	"os"
)

func CreateQueryFile(key string) error {
	// Open the file in append mode, create it if it doesn't exist
	file, err := os.OpenFile("../../test_output/query.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// write the string to the file
	if _, err := file.WriteString(key); err != nil {
		return err
	}

	fmt.Println("Successful create query file")
	return nil
}
