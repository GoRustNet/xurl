package main

import (
	"flag"
	"log"

	"github.com/GoRustNet/xurl/conf"
	"github.com/GoRustNet/xurl/db"
	v1 "github.com/GoRustNet/xurl/httpapi/v1"
	"github.com/gin-gonic/gin"
)

func init() {
	cfgFilename := flag.String("c", "./config.json", "配置文件")
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := conf.InitFrom(*cfgFilename); err != nil {
		log.Fatalf("Init config %q failed: %v\n", *cfgFilename, err)
	}
	if err := db.Init(conf.Cfg.Pg); err != nil {
		log.Fatal("Init db failed:", err)
	}
}

func main() {
	app := gin.New()
	app.Use(gin.Logger(), gin.Recovery())
	v1.RegisterRouter(app, "/")
	log.Fatal(app.Run(conf.Cfg.Web.Addr))
}
