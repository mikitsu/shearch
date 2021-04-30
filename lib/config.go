package lib

import "strings"

type Config struct {
  shortcutPrefix string
  defaultRedirect string
  shortcutSeparator string
  shortcuts map[string]string
}

func GetConfig(shortcutPrefix, defaultRedirect, shortcutSeparator string) Config {
  if ! strings.Contains(defaultRedirect, "%s") {
    defaultRedirect += "%s"
  }
  return Config{shortcutPrefix, defaultRedirect, shortcutSeparator, map[string]string{}}
}
