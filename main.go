package main

import (
	"embed"
	"log"
	"os/exec"

	"go-moc/internal/app/mocp"
	"go-moc/internal/app/startup"
	"go-moc/internal/app/transport/http"
)

//go:embed web
var static embed.FS

func main() {
	_, err := exec.LookPath("mocp")
	if err != nil {
		log.Fatal("installing mocp is in your future")
	}

	moc := mocp.NewMOCp()
	handler := http.NewHTTPHandler(static, moc)
	router := http.NewGinRouter(handler)
	startup.StartServer(router)
}
