package lib

import (
  "log"
  "net"
  "net/http"
  "strings"
)

type FixedConfig struct {
  Addr string
  ConfigPath string
  ShortcutPath string
}

type variableConfig struct {
  shortcutPrefix string
  defaultRedirect string
  shortcutSeparator string
  shortcuts map[string]string
}

func (fconf FixedConfig) GetListener() net.Listener {
  var lnet string
  laddr := fconf.Addr
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

func (fconf FixedConfig) GetServeMux() *http.ServeMux {
  vconf := variableConfig{"!", "http://localhost/", " ", map[string]string{}}
  mux := http.NewServeMux()
  mux.HandleFunc("/", vconf.handleQuery)
  return mux
}
