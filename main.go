package main

import (
	"flag"
	"log"

	"github.com/GoRustNet/xurl/conf"
)

func init() {
	cfgFilename := flag.String("c", "./config.json", "配置文件")
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := conf.InitFrom(*cfgFilename); err != nil {
		log.Fatalf("Init config %q failed: %v\n", *cfgFilename, err)
	}
}

func main() {

}
