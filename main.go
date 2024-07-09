package main

import (
	"flag"
	"github.com/MetaerMarket/tg-bot-master/config"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

var (
	version = "aio version: aio/1.25.18"
)

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	log.SetOutput(os.Stdout)
	path := flag.String("c", "one-bot.yml", "项目的配置文件地址(使用绝对路径) 例: -c /etc/one-bot.yml")
	v := flag.Bool("v", false, "返回当前版本")
	flag.Parse()
	if *v {
		log.Println(version)
		return
	}
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Runtime panic caught: %v\n", err)
		}
		log.Println("end")
	}()

	go func() {
		log.Println("ListenAndServe start")
		http.Handle("/", &helloHandler{})
		http.ListenAndServe(":6060", nil)
	}()

	config.Load(*path)

}
