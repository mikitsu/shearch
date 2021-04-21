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
  http.Serve(listener, conf)
}

func (conf Config) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query().Get("q")
  redirect := conf.getShortcutRedirect(query)
  w.Header().Add("Location", redirect)
  w.WriteHeader(http.StatusTemporaryRedirect)
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
