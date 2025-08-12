package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aprdec/rjgl/pkg/setting"
	"github.com/aprdec/rjgl/routers"
)

func main() {
	r := routers.InitRouter()

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", setting.HTTPPort),
		ReadTimeout:  setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		Handler:      r,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("ListenAndServe: %v\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
