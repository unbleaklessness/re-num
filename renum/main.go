package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
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

	rand.Seed(time.Now().UnixNano())

	for _, fileDescription := range fileDescriptions {

		if fileDescription.IsDir() {
			continue
		}

		newFileName := fileDescription.Name() + " (" + strconv.Itoa(rand.Int()) + ")" + filepath.Ext(fileDescription.Name())

		oldFilePath := filepath.Join(currentDirectory, fileDescription.Name())
		newFilePath := filepath.Join(currentDirectory, newFileName)

		e := os.Rename(oldFilePath, newFilePath)
		if e != nil {
			fmt.Printf("Could not rename \"%s\" into \"%s\"!\n", oldFilePath, newFilePath)
			return
		}
	}

	fileDescriptions, e = ioutil.ReadDir(currentDirectory)
	if e != nil {
		fmt.Println("Could not list files in the current directory!")
		return
	}

	nFileDescriptions = len(fileDescriptions)
	if nFileDescriptions > maximumNumberForPadding(*padding) {
		fmt.Printf("Number of files in the current directory (= %d) is too big for this padding!\n", nFileDescriptions)
		return
	}

	for i, fileDescription := range fileDescriptions {

		if fileDescription.IsDir() {
			continue
		}

		template := "%0" + strconv.Itoa(*padding) + "d"
		index := i + 1
		newFileName := fmt.Sprintf(template, index) + filepath.Ext(fileDescription.Name())

		oldFilePath := filepath.Join(currentDirectory, fileDescription.Name())
		newFilePath := filepath.Join(currentDirectory, newFileName)

		e := os.Rename(oldFilePath, newFilePath)
		if e != nil {
			fmt.Printf("Could not rename \"%s\" into \"%s\"!\n", oldFilePath, newFilePath)
			return
		}
	}
}
