package http

import "go-moc/internal/app/mocp"

type TrackInfoResponse struct {
	Title     string `json:"title"`
	FileName  string `json:"file_name"`
	Artist    string `json:"artist"`
	SongTitle string `json:"song_title"`
	Album     string `json:"album"`
}

func mapTrackInfoToTrackInfoResponse(info mocp.TrackInfo) TrackInfoResponse {
	return TrackInfoResponse{
		Title:     info.Title,
		FileName:  info.FileName,
		Artist:    info.Artist,
		SongTitle: info.SongTitle,
		Album:     info.Album,
	}
}
