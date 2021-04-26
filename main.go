package main

import (
  "log"
  "net/http"
  
  "github.com/mik2k2/shearch/lib"
)

func main() {
  conf := lib.FixedConfig{"unix:/tmp/sock", "", ""}
  listener := conf.GetListener()
  log.Println("listening on", listener.Addr())
  http.Serve(listener, conf.GetServeMux())
}
