package lib

type Config struct {
  shortcutPrefix string
  defaultRedirect string
  shortcutSeparator string
  shortcuts map[string]string
}

func GetConfig(shortcutPrefix, defaultRedirect, shortcutSeparator string) Config {
  return Config{shortcutPrefix, defaultRedirect, shortcutSeparator, map[string]string{}}
}
