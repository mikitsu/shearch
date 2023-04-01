package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"git.sr.ht/~mikitsu/sdnet"

	"github.com/mik2k2/shearch/lib"
)

func main() {
	listener, conf := getConfig()
	log.Println(conf)
	log.Println("listening on", listener.Addr())
	http.Serve(listener, conf)
}

func getListener(laddr string, sock_mode uint) net.Listener {
	var lnet string
	if strings.HasPrefix(laddr, "unix:") {
		laddr = laddr[5:]
		lnet = "unix"
	} else {
		lnet = "tcp"
	}
	listener, err := sdnet.Listen(lnet, laddr)
	if err != nil {
		log.Fatal(err)
	}
	if lnet == "unix" {
		err := os.Chmod(laddr, os.FileMode(sock_mode))
		if err != nil {
			log.Println("error setting socket mode: ", err)
		}
	}
	return listener
}

func getConfig() (net.Listener, *lib.MainConfig) {
	var laddr, prefix, redirect, separator string
	opensearch := lib.URLValue{}
	var mode uint
	flag.StringVar(&laddr, "listen", "127.0.0.1:8080",
		"listen on; may also be a unix:/domain.socket")
	flag.UintVar(&mode, "sock-mode", 0777, "file permissions when listening on UDS")
	flag.StringVar(&prefix, "prefix", "!", "default shortcut prefix")
	flag.StringVar(&redirect, "redirect", "https://duckduckgo.com/?q=",
		"default redirect location if no shortcut matches")
	flag.StringVar(&separator, "separator", " ",
		"default separator between shortcut and query")
	flag.Var(&opensearch, "url", "URL shearch is available on (for opensearch)")
	flag.Parse()
	return getListener(laddr, mode), lib.GetConfig(prefix, redirect, separator, opensearch.URL)
}
