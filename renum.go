package main

import (
	"flag"
	"fmt"
	"io/fs"
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

func readDirectory(currentDirectory string, renameDirectories bool) []fs.FileInfo {

	fileDescriptions0, e := ioutil.ReadDir(currentDirectory)
	if e != nil {
		panic("Could not list files in the current directory!")
	}

	fileDescriptions := make([]fs.FileInfo, 0)
	for _, fileDescription := range fileDescriptions0 {
		isDirectory := fileDescription.IsDir()
		if (renameDirectories && isDirectory) || (!renameDirectories && !isDirectory) {
			fileDescriptions = append(fileDescriptions, fileDescription)
		}
	}

	return fileDescriptions
}

func paddingCheck(fileDescriptions []fs.FileInfo, padding int) {
	nFileDescriptions := len(fileDescriptions)
	if nFileDescriptions > maximumNumberForPadding(padding) {
		errorMessage := fmt.Sprintf("Number of files or directories in the current directory (= %d) is too big for this padding!\n", nFileDescriptions)
		panic(errorMessage)
	}
}

func main() {

	padding := flag.Int("p", 5, "Padding (number of leading zeros).")
	renameDirectories := flag.Bool("d", false, "Rename directories instead of files.")
	directoriesPostfix := flag.String("dp", "_", "Postfix for directories.")
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

	fileDescriptions := readDirectory(currentDirectory, *renameDirectories)
	paddingCheck(fileDescriptions, *padding)

	rand.Seed(time.Now().UnixNano())

	for _, fileDescription := range fileDescriptions {

		newFileName := strconv.Itoa(rand.Int()) + filepath.Ext(fileDescription.Name())

		oldFilePath := filepath.Join(currentDirectory, fileDescription.Name())
		newFilePath := filepath.Join(currentDirectory, newFileName)

		e := os.Rename(oldFilePath, newFilePath)
		if e != nil {
			fmt.Printf("Could not rename \"%s\" into \"%s\"!\n", oldFilePath, newFilePath)
			return
		}
	}

	fileDescriptions = readDirectory(currentDirectory, *renameDirectories)
	paddingCheck(fileDescriptions, *padding)

	for i, fileDescription := range fileDescriptions {

		template := "%0" + strconv.Itoa(*padding) + "d"
		index := i + 1
		newFileName := fmt.Sprintf(template, index) + filepath.Ext(fileDescription.Name())

		if fileDescription.IsDir() {
			newFileName += *directoriesPostfix
		}

		oldFilePath := filepath.Join(currentDirectory, fileDescription.Name())
		newFilePath := filepath.Join(currentDirectory, newFileName)

		e := os.Rename(oldFilePath, newFilePath)
		if e != nil {
			fmt.Println(e.Error())
			fmt.Printf("Could not rename \"%s\" into \"%s\"!\n", oldFilePath, newFilePath)
			return
		}
	}
}
