package main

import (
	"composeapp/db"
	"flag"
	"fmt"

	"composeapp/internal/config"
	"composeapp/internal/handler"
	"composeapp/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/zerodb-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	connection := db.GetConnection()

	ctx := svc.NewServiceContext(c, connection)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
