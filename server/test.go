package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

const emailFolderPath = "./enron_mail_20110402/maildir"

func GetMailUsers() ([]string, error) {
	var users []string

	dirs, err := ioutil.ReadDir(emailFolderPath)
	if err != nil {
		return nil, err
	}

	for _, dir := range dirs {
		//fmt.Println(dir.Name())
		users = append(users, dir.Name())
	}
	return users, nil

}

func ProcessMailsByUser(user string, wg *sync.WaitGroup) {
	defer wg.Done()

	var countmails int
	path := emailFolderPath + "/" + user

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("Error in visitAndProcessEmailFiles: "+path, err)
		}
		if !info.IsDir() {
			//log.Println("Sucess open file: " + path)
			countmails += 1
		}
		return nil
	})

	if err != nil {
		return
	}
	fmt.Println(countmails)
}

func main() {
	usersId, err := GetMailUsers()
	var count int
	// a := 6
	if err != nil {
		return
	}

	var wg sync.WaitGroup

	for _, us := range usersId {
		wg.Add(1)
		go ProcessMailsByUser(us, &wg)
	}
	wg.Wait()

	fmt.Print("total: " + strconv.Itoa(count))
}
