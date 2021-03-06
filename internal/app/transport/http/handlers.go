package http

import (
	"embed"
	"net/http"

	"go-moc/internal/app/mocp"

	"github.com/gin-gonic/gin"
)

type MusicOnConsolePlayer interface {
	Play() error
	Stop() error
	TogglePause() error
	Next() error
	Previous() error
	Seek(bool) error
	Exit() error
	Volume(bool) error
	Mute() error
	TrackInfo() (mocp.TrackInfo, error)
}

type HTTPHandler struct {
	static embed.FS
	mocp MusicOnConsolePlayer
}

func NewHTTPHandler(static embed.FS, mocp MusicOnConsolePlayer) *HTTPHandler {
	return &HTTPHandler{
		static: static,
		mocp: mocp,
	}
}

func (h *HTTPHandler) Index(ctx *gin.Context) {
	ctx.FileFromFS("web/index.htm", http.FS(h.static))
}

func (h *HTTPHandler) Play(ctx *gin.Context) {
	err := h.mocp.Play()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.Status(http.StatusOK)
}

func (h *HTTPHandler) Stop(ctx *gin.Context) {
	err := h.mocp.Stop()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.Status(http.StatusOK)
}

func (h *HTTPHandler) TogglePause(ctx *gin.Context) {
	err := h.mocp.TogglePause()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.Status(http.StatusOK)
}

func (h *HTTPHandler) Next(ctx *gin.Context) {
	err := h.mocp.Next()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.Status(http.StatusOK)
}

func (h *HTTPHandler) Previous(ctx *gin.Context) {
	err := h.mocp.Previous()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.Status(http.StatusOK)
}

func (h *HTTPHandler) Forward(ctx *gin.Context) {
	err := h.mocp.Seek(true)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.Status(http.StatusOK)
}

func (h *HTTPHandler) Backward(ctx *gin.Context) {
	err := h.mocp.Seek(false)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.Status(http.StatusOK)
}


func (h *HTTPHandler) Exit(ctx *gin.Context) {
	err := h.mocp.Exit()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.Status(http.StatusOK)
}

func (h *HTTPHandler) VolumeIncrease(ctx *gin.Context) {
	err := h.mocp.Volume(true)
	if err != nil {
		ctx.Status(http.StatusBadRequest)

		return
	}

	ctx.Status(http.StatusOK)
}

func (h *HTTPHandler) VolumeDecrease(ctx *gin.Context) {
	err := h.mocp.Volume(false)
	if err != nil {
		ctx.Status(http.StatusBadRequest)

		return
	}

	ctx.Status(http.StatusOK)
}

func (h *HTTPHandler) Mute(ctx *gin.Context) {
	err := h.mocp.Mute()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.Status(http.StatusOK)
}

func (h *HTTPHandler) TrackInfo(ctx *gin.Context) {
	trackInfo, err := h.mocp.TrackInfo()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, mapTrackInfoToTrackInfoResponse(trackInfo))
}
