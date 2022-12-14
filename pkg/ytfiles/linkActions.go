package ytfiles

import (
	"io"
	"log"
	"os"

	"github.com/kkdai/youtube/v2"
)

func PrepareForSending(userId int, ytLink string) (*SongsInfo, error) {
	var songsList SongsInfo

	DeleteDir(userId)
	dir := createDir(userId)

	completeSongLins(dir, ytLink, &songsList)

	return &songsList, nil
}

func completeSongLins(dir, ytLink string, songsList *SongsInfo) {
	var newSong SongInfo
	client := youtube.Client{}

	switch setLinkType(ytLink) {
	case "playlist":
		playlistId, err := extractPlaylistID(ytLink)
		if err != nil {
			log.Fatal(err)
		}
		playlist, err := client.GetPlaylist(playlistId)
		if err != nil {
			panic(err)
		}
		for _, video := range playlist.Videos {
			newSong.setID(video.ID)
			*songsList = append(*songsList, newSong)
		}
	default:
		videoId, err := youtube.ExtractVideoID(ytLink)
		if err != nil {
			log.Fatal(err)
		}
		newSong.setID(videoId)
		*songsList = append(*songsList, newSong)
	}
	getStreamAndDownload(dir, client, songsList)
}

func getStreamAndDownload(dir string, client youtube.Client, songsList *SongsInfo) {

	for i := 0; i < len(*songsList); i++ {
		video, err := client.GetVideo((*songsList)[i].ID)
		if err != nil {
			log.Fatal(err)
		}
		(*songsList)[i].add(video)
		(*songsList)[i].setPath(dir + "/" + (*songsList)[i].Name + ".mp3")
		formats := video.Formats.Type("audio")
		stream, _, err := client.GetStream(video, &formats[0])
		if err != nil {
			log.Fatal(err)
		}

		file, err := os.Create((*songsList)[i].FullPath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		_, err = io.Copy(file, stream)
		if err != nil {
			log.Fatal(err)
		}

		if err != nil {
			log.Fatal(err)
		}
	}
}
