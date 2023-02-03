package main

import (
	gvalidator2 "github.com/whilesun/go-admin/pkg/utils/gvalidator"
	"github.com/whilesun/go-admin/routers"
	"net/http"
	"time"
)

func main() {
	//models.NewSysInit().Run()
	gvalidator2.InitGinValidate("zh")
	router := routers.InitRouter()
	s := &http.Server{
		Addr:         ":3000",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	s.ListenAndServe()
}
