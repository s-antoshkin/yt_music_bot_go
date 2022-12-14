package ytfiles

import (
	"log"
	"os"
	"strconv"
)

func DeleteDir(userId int) {
	err := os.RemoveAll("files/" + strconv.Itoa(userId))
	if err != nil {
		log.Fatal(err)
	}
}

func createDir(userId int) string {
	dir := "files/" + strconv.Itoa(userId)
	err := os.MkdirAll(dir, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
