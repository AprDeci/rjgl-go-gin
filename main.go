package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	docs "github.com/aprdec/rjgl/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/aprdec/rjgl/pkg/setting"
	"github.com/aprdec/rjgl/routers"
)

func main() {
	//docs
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Title = "RJGL API"
	docs.SwaggerInfo.Description = "RJGL API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := routers.InitRouter()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

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
