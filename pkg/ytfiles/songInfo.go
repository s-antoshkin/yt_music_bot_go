package ytfiles

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	youtube "github.com/kkdai/youtube/v2"
)

type SongsInfo []SongInfo

type SongInfo struct {
	ID       string
	Name     string
	FullPath string
	Uploaded bool
}

func (s *SongInfo) setID(id string) {
	s.ID = id
}

func createAuthor(authorOld string) string {
	match, err := regexp.Match(" - Topic", []byte(authorOld))
	if err != nil {
		log.Fatal(err)
	}
	if match {
		return strings.Split(authorOld, " - Topic")[0]
	}

	return authorOld
}

func checkTitle(author, title string) bool {
	match, err := regexp.Match(author, []byte(title))
	if err != nil {
		log.Fatal(err)
	}
	if match {
		return true
	}
	return false
}

func (s *SongInfo) add(video *youtube.Video) {
	s.ID = video.ID
	vidAuthor := createAuthor(video.Author)
	isTitle := checkTitle(vidAuthor, video.Title)
	if isTitle {
		s.Name = video.Title
	} else {
		s.Name = fmt.Sprintf("%s - %s", vidAuthor, video.Title)
	}
	s.FullPath = ""
	s.Uploaded = false
}

func (s *SongInfo) setPath(path string) {
	s.FullPath = path
}
