package lib

import (
  "strings"
  "net/url"
)

func (conf Config) getShortcutRedirect(query string) string {
  if ! strings.HasPrefix(query, conf.ShortcutPrefix){
    return conf.DefaultRedirect + url.QueryEscape(query)
  }
  return "TODO"
}
