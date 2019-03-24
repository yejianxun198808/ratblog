package main

import (
	"net/http"
	"ratblog/routers"
)

//gin 服务启动
func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Handler: router,
	}
	s.ListenAndServe()
}                    
