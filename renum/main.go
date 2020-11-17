package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func maximumNumberForPadding(n int) int {
	n--
	result := 9
	for i := 0; i < n; i++ {
		result *= 10
		result += 9
	}
	return result
}

func main() {

	padding := flag.Int("p", 5, "Padding (number of leading zeros).")
	flag.Parse()

	if *padding < 1 {
		fmt.Println("Padding must be greater than 0!")
		return
	}

	currentDirectory, e := os.Getwd()
	if e != nil {
		fmt.Println("Could not get current directory!")
		return
	}

	fileDescriptions, e := ioutil.ReadDir(currentDirectory)
	if e != nil {
		fmt.Println("Could not list files in the current directory!")
		return
	}

	nFileDescriptions := len(fileDescriptions)
	if nFileDescriptions > maximumNumberForPadding(*padding) {
		fmt.Printf("Number of files in the current directory (= %d) is too big for this padding!\n", nFileDescriptions)
		return
	}

	i := 1

	for _, fileDescription := range fileDescriptions {

		if fileDescription.IsDir() {
			continue
		}

		template := "%0" + strconv.Itoa(*padding) + "d"
		newFileName := fmt.Sprintf(template, i) + filepath.Ext(fileDescription.Name())
		i++

		oldFilePath := filepath.Join(currentDirectory, fileDescription.Name())
		newFilePath := filepath.Join(currentDirectory, newFileName)

		e := os.Rename(oldFilePath, newFilePath)
		if e != nil {
			fmt.Printf("Could not rename \"%s\" into \"%s\"!\n", oldFilePath, newFilePath)
		}
	}
}
