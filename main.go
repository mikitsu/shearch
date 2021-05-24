package main

import (
  "log"
  "flag"
  "net"
  "net/http"
  "strings"

  "github.com/mik2k2/shearch/lib"
)

func main() {
  laddr, conf := getConfig()
  listener := getListener(laddr)
  log.Println("listening on", listener.Addr())
  http.Serve(listener, &conf)
}

func getListener(laddr string) net.Listener {
  var lnet string
  if strings.HasPrefix(laddr, "unix:") {
    laddr = laddr[5:]
    lnet = "unix"
  } else {
    lnet = "tcp"
  }
  listener, err := net.Listen(lnet, laddr)
  if err != nil {
    log.Fatal(err)
  }
  return listener
}

func getConfig() (string, lib.Config) {
  var laddr, prefix, redirect, separator, opensearch string
  flag.StringVar(&laddr, "listen", "127.0.0.1:8080",
                 "listen on; may also be a unix:/domain.socket")
  flag.StringVar(&prefix, "prefix", "!", "default shortcut prefix")
  flag.StringVar(&redirect, "redirect", "https://duckduckgo.com/?q=",
                 "default redirect location if no shortcut matches")
  flag.StringVar(&separator, "separator", " ",
                 "default separator between shortcut and query")
  flag.StringVar(&opensearch, "url", "", "URL shearch is available on (for opensearch)")
  flag.Parse()
  return laddr, lib.GetConfig(prefix, redirect, separator, opensearch)
}
