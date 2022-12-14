package ytfiles

import (
	"log"
	"regexp"

	"github.com/kkdai/youtube/v2"
)

var (
	playlistIDRegex    = regexp.MustCompile("^[A-Za-z0-9_-]{13,42}$")
	playlistInURLRegex = regexp.MustCompile("[&?]list=([A-Za-z0-9_-]{13,42})(&.*)?$")
)

func setLinkType(link string) string {
	match, err := regexp.Match("/watch", []byte(link))
	if err != nil {
		log.Fatal(err)
	}
	if match {
		return "song"
	} else {
		return "playlist"
	}
}

func extractPlaylistID(url string) (string, error) {
	if playlistIDRegex.Match([]byte(url)) {
		return url, nil
	}

	matches := playlistInURLRegex.FindStringSubmatch(url)

	if matches != nil {
		return matches[1], nil
	}

	return "", youtube.ErrInvalidPlaylist
}
