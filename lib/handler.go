package lib

import (
  "fmt"
  "strings"
  "net/http"
  "net/url"
)


func (conf *variableConfig) handleQuery(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query().Get("q")
  redirect := conf.getShortcutRedirect(query)
  w.Header().Add("Location", redirect)
  w.WriteHeader(http.StatusTemporaryRedirect)
}

func (conf variableConfig) getShortcutRedirect(query string) string {
  default_response := conf.defaultRedirect + url.QueryEscape(query)
  if ! strings.HasPrefix(query, conf.shortcutPrefix){
    return default_response
  }
  query = query[len(conf.shortcutPrefix):]
  split_query := strings.SplitN(query, conf.shortcutSeparator, 2)
  cut := split_query[0]
  query = split_query[1]
  for key, value := range conf.shortcuts {
    if cut == key {
      return fmt.Sprintf(value, url.QueryEscape(query))
    }
  }
  return default_response
}
