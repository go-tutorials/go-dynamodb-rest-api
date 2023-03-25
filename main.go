package main

import (
	"context"
	"fmt"
	"github.com/core-go/config"
	"github.com/core-go/log"
	mid "github.com/core-go/log/middleware"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"go-service/internal/app"
)

func main() {
	var conf app.Config
	err := config.Load(&conf, "configs/config")
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	log.Initialize(conf.Log)
	r.Use(mid.BuildContext)
	logger := mid.NewLogger()
	if log.IsInfoEnable() {
		r.Use(mid.Logger(conf.MiddleWare, log.InfoFields, logger))
	}
	r.Use(mid.Recover(log.ErrorMsg))

	err = app.Route(context.Background(), r, conf.DB)
	if err != nil {
		panic(err)
	}

	fmt.Println("Start server")
	server := ""
	if conf.Server.Port > 0 {
		server = ":" + strconv.FormatInt(conf.Server.Port, 10)
	}
	if err = http.ListenAndServe(server, r); err != nil {
		fmt.Println(err.Error())
	}
}
