package minitask6

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(path string) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic:", err)
			fmt.Println("continue...")
		}
	}()

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		file.Close()
	}()

	info, err := file.Stat()
	if err == nil && info.IsDir() {
		panic("path mengarah ke directory")
	}

	bacaFile := bufio.NewScanner(file)
	for bacaFile.Scan() {
		fmt.Println(bacaFile.Text())
	}
	if err := bacaFile.Err(); err != nil {
		panic(err)
	}

	return nil
}
