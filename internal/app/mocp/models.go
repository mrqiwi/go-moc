package mocp

import (
	"path"
	"path/filepath"
	"strings"
)

type TrackInfo struct {
	State       string
	FileName    string
	Title       string
	Artist      string
	SongTitle   string
	Album       string
	TotalTime   string
	TimeLeft    string
	TotalSec    string
	CurrentTime string
	CurrentSec  string
	Bitrate     string
	AvgBitrate  string
	Rate        string
}

func fillTrackInfo(data string) TrackInfo {
	var trackInfo = TrackInfo{}

	info := strings.Split(data, "\n")
	if len(info) == 0 {
		return trackInfo
	}

	if strings.Contains(info[0], "State: ") {
		state := info[0]
		trackInfo.State = state[len("State: "):]
	}

	if strings.Contains(info[1], "File: ") {
		file := info[1]
		basePath := filepath.Base(file[len("File: "):])
		trackInfo.FileName = strings.TrimSuffix(basePath, path.Ext(basePath))
	}

	if strings.Contains(info[2], "Title: ") {
		file := info[2]
		trackInfo.Title = file[len("Title: "):]
	}

	if strings.Contains(info[3], "Artist: ") {
		file := info[3]
		trackInfo.Artist = file[len("Artist: "):]
	}

	if strings.Contains(info[4], "SongTitle: ") {
		file := info[4]
		trackInfo.SongTitle = file[len("SongTitle: "):]
	}

	if strings.Contains(info[5], "Album: ") {
		file := info[5]
		trackInfo.Album = file[len("Album: "):]
	}

	if strings.Contains(info[6], "TotalTime: ") {
		file := info[6]
		trackInfo.TotalTime = file[len("TotalTime: "):]
	}

	if strings.Contains(info[7], "TimeLeft: ") {
		file := info[7]
		trackInfo.TimeLeft = file[len("TimeLeft: "):]
	}

	if strings.Contains(info[8], "TotalSec: ") {
		file := info[8]
		trackInfo.TotalSec = file[len("TotalSec: "):]
	}

	if strings.Contains(info[9], "CurrentTime: ") {
		file := info[9]
		trackInfo.CurrentTime = file[len("CurrentTime: "):]
	}

	if strings.Contains(info[10], "CurrentSec: ") {
		file := info[10]
		trackInfo.CurrentSec = file[len("CurrentSec: "):]
	}

	if strings.Contains(info[11], "Bitrate: ") {
		file := info[11]
		trackInfo.Bitrate = file[len("Bitrate: "):]
	}

	if strings.Contains(info[12], "AvgBitrate: ") {
		file := info[12]
		trackInfo.AvgBitrate = file[len("AvgBitrate: "):]
	}

	if strings.Contains(info[13], "Rate: ") {
		file := info[13]
		trackInfo.Rate = file[len("Rate: "):]
	}

	return trackInfo
}
