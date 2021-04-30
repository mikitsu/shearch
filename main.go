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
  conf := getConfig()
  listener := getListener("unix:/tmp/sock")
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

func getConfig() lib.Config {
  var prefix, redirect, separator string
  flag.StringVar(&prefix, "prefix", "!", "default shortcut prefix")
  flag.StringVar(&redirect, "redirect", "https://duckduckgo.com/?q=",
                 "default redirect location if no shortcut matches")
  flag.StringVar(&separator, "separator", " ",
                 "default separator between shortcut and query")
  flag.Parse()
  return lib.GetConfig(prefix, redirect, separator)
}
