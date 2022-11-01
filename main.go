package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const path = "/Users/sabrinabilmak/Downloads"

func main() {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		// isDir - czy to folder
		if f.IsDir() {
			continue
		}

		type_ := getType(f.Name())
		folder := getFolder(type_)
		pathForSearching := path + "/" + folder

		if _, err := os.Stat(pathForSearching); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(pathForSearching, os.ModePerm)
			if err != nil {
				log.Println(err)
			}
		}

		oldLocation := path + "/" + f.Name()
		newLocation := path + "/" + folder + "/" + f.Name()
		err := os.Rename(oldLocation, newLocation)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(pathForSearching)
	}

}

func getType(filename string) string {
	result := strings.Split(filename, ".")
	return result[len(result)-1]
}

func getFolder(type_ string) string {
	if type_ == "pdf" || type_ == "doc" || type_ == "epub" || type_ == "docx" {
		return "books"
	}
	if type_ == "png" || type_ == "jpg" || type_ == "jpeg" {
		return "photo"
	}
	if type_ == "zip" || type_ == "dmg" || type_ == "pkg" {
		return "garbage"
	}
	return "unknown"
}
