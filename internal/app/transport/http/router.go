package http

import "github.com/gin-gonic/gin"

func NewGinRouter(h *HTTPHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/play", h.Play)
	r.POST("/stop", h.Stop)
	r.POST("/pause", h.TogglePause)
	r.POST("/next", h.Next)
	r.POST("/previous", h.Previous)
	r.POST("/exit", h.Exit)
	r.POST("/volume-increase", h.VolumeIncrease)
	r.POST("/volume-decrease", h.VolumeDecrease)
	r.POST("/mute", h.Mute)
	r.GET("/info", h.TrackInfo)

	return r
}
