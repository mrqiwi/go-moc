package startup

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func StartServer(router *gin.Engine) {
	const addr = "localhost:9999"

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("listening on %s", addr)
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Printf("failed to listen: %s", err)
			}
		}
	}()
	<-stop

	log.Println("shutting down ...")

	if err := server.Shutdown(context.TODO()); err != nil {
		log.Printf("failed to shutown %s: ", err)
	}
}
