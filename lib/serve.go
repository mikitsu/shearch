package lib

import (
  "log"
  "net"
  "net/http"
  "strings"
)

func Serve() {
  conf := loadConfig()
  listener := getListener(conf.Addr)
  log.Println("listening on", listener.Addr())
  mux := conf.getServeMux()
  http.Serve(listener, mux)
}

func getListener (laddr string) net.Listener {
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

func (conf Config) getServeMux() *http.ServeMux {
  mux := http.NewServeMux()
  mux.HandleFunc("/", conf.handleQuery)
  return mux
}
