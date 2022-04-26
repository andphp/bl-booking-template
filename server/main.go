package main

import (
	"context"
	"server/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"server/boot"

	"github.com/andphp/go-gin/goby"
)

func main() {
	// 初始化配置
	boot.InitConfig()

	boot.GormBoot()

	ainRouter := goby.MakeGin().RouterMount("")(router.Admin).RouterMount("")(router.Menu)

	s := &http.Server{
		Addr:           ":8081",
		Handler:        ainRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	time.Sleep(10 * time.Microsecond)

	go func() {

		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			boot.FatalErrorChan <- err
		}

	}()
	go func() {
		err := <-boot.FatalErrorChan
		if err != nil {
			shutdown(s, err)
		}
	}()
	signal.Notify(boot.Quit, os.Interrupt)
	errQ := <-boot.Quit
	if errQ != nil {
		shutdown(s, errQ)
	}
}

func shutdown(s *http.Server, err interface{}) {
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer close(boot.Quit)
	defer close(boot.FatalErrorChan)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		if goby.GOBY_DB != nil {
			// 程序结束前关闭数据库链接
			db, _ := goby.GOBY_DB.DB()
			defer db.Close()
		}
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
