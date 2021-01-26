package main

import (
	"fmt"
	"log"
	"modifyStep/controller"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

}
func main() {
	r := gin.Default()
	r.Use(controller.Cors())
	r.POST("/modifyStep", controller.ModifyStep)
	//初始化http配置
	s := &http.Server{
		Addr:           "0.0.0.0:9000",
		Handler:        r,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println(`Welcome to labortool-go 
	 默认运行地址: http://0.0.0.0:9000`)

	err := s.ListenAndServe()
	if err != nil {
		panic("bootstrap error")
	}
}
