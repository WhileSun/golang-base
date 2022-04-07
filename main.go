package main

import (
	"github.com/whilesun/go-admin/app/models"
	"github.com/whilesun/go-admin/pkg/gvalidator"
	"github.com/whilesun/go-admin/routers"
	"net/http"
	"time"
)

func main() {
	models.NewSysInit().Run()
	gvalidator.InitGinValidate("zh")
	router := routers.InitRouter()
	s := &http.Server{
		Addr:         ":3000",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	s.ListenAndServe()
}
