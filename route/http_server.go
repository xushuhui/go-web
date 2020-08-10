package route

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"goal/global"
	"log"
	"net/http"
	"time"
)

var (
	HttpSrvHandler *http.Server
)

func HttpServerRun() {
	fmt.Println("mode", global.ServerSetting.RunMode)
	gin.SetMode(global.ServerSetting.RunMode)
	r := InitRouter()
	port := global.ServerSetting.HttpPort
	HttpSrvHandler = &http.Server{
		Addr:           ":" + port,
		Handler:        r,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		log.Printf(" [INFO] HttpServerRun:%s\n", port)
		if err := HttpSrvHandler.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", port, err)
		}
	}()
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}
