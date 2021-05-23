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
	var conf app.Root
	er1 := config.Load(&conf, "configs/config")
	if er1 != nil {
		panic(er1)
	}

	r := mux.NewRouter()

	log.Initialize(conf.Log)
	r.Use(mid.BuildContext)
	logger := mid.NewStructuredLogger()
	if log.IsInfoEnable() {
		r.Use(mid.Logger(conf.MiddleWare, log.InfoFields, logger))
	}
	r.Use(mid.Recover(log.ErrorMsg))

	er2 := app.Route(r, context.Background(), conf.DB)
	if er2 != nil {
		panic(er2)
	}

	fmt.Println("Start server")
	server := ""
	if conf.Server.Port > 0 {
		server = ":" + strconv.FormatInt(conf.Server.Port, 10)
	}
	if er3 := http.ListenAndServe(server, r); er3 != nil {
		fmt.Println(er3.Error())
	}
}
