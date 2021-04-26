package lib

type Config struct {
  Addr string
  ShortcutPrefix string
  DefaultRedirect string
  ShortcutSeparator string
  Shortcuts map[string]string
}

func loadConfig() Config {
  return Config{"unix:/tmp/sock", "!", "http://localhost", " ",
    map[string]string{"w": "https://en.wikipedia.org/wiki/%s",
      "wde": "https://de.wikipedia.org/wiki/%s"}}
}
