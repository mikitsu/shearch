package lib

import (
  "fmt"
  "strings"
  "net/http"
  "net/url"
)

func (conf *Config) handleQuery(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query().Get("q")
  redirect := conf.getShortcutRedirect(query)
  w.Header().Add("Location", redirect)
  w.WriteHeader(http.StatusTemporaryRedirect)
}

func (conf Config) getShortcutRedirect(query string) string {
  default_response := conf.DefaultRedirect + url.QueryEscape(query)
  if ! strings.HasPrefix(query, conf.ShortcutPrefix){
    return default_response
  }
  query = query[len(conf.ShortcutPrefix):]
  split_query := strings.SplitN(query, conf.ShortcutSeparator, 2)
  cut := split_query[0]
  query = split_query[1]
  for key, value := range conf.Shortcuts {
    if cut == key {
      return fmt.Sprintf(value, url.QueryEscape(query))
    }
  }
  return default_response
}
