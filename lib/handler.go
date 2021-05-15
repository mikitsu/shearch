package lib

import (
  "fmt"
  "strings"
  "net/http"
  "net/url"
)

func (conf *Config) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
  case http.MethodGet:
    conf.handleQuery(w, r)
  case http.MethodPut:
    conf.handleConfig(w, r)
  case http.MethodPost:
    conf.handleShortcut(w, r)
  }
}


func (conf *Config) handleQuery(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query().Get("q")
  redirect := conf.getShortcutRedirect(query)
  w.Header().Add("Location", redirect)
  w.WriteHeader(http.StatusTemporaryRedirect)
}

func (conf Config) getShortcutRedirect(query string) string {
  response := conf.defaultRedirect
  if strings.HasPrefix(query, conf.shortcutPrefix){
    split_query := strings.SplitN(query[len(conf.shortcutPrefix):], conf.shortcutSeparator, 2)
    cut := split_query[0]
    part_query := split_query[1]
    for key, value := range conf.shortcuts {
      if cut == key {
        response = value
        query = part_query
        break
      }
    }
  }
  if strings.Contains(response, "%s") {
    return fmt.Sprintf(response, url.QueryEscape(query))
  } else {
    return response + url.QueryEscape(query)
  }
}

func (conf *Config) handleConfig(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()
  if newPrefix := query.Get("prefix"); newPrefix != "" {
    conf.shortcutPrefix = newPrefix
  }
  if newRedirect := query.Get("redirect"); newRedirect != "" {
    conf.defaultRedirect = newRedirect
  }
  if newSeparator := query.Get("separator"); newSeparator != "" {
    conf.shortcutSeparator = newSeparator
  }
}

func (conf *Config) handleShortcut(w http.ResponseWriter, r *http.Request) {
  for k, v := range r.URL.Query() {
    newDest := v[0]
    conf.shortcuts[k] = newDest
  }
}
