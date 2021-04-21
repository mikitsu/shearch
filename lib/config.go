package lib

type Config struct {
  Addr string
  ShortcutPrefix string
  DefaultRedirect string
  Shortcuts map[string]string
}

func loadConfig() Config {
  return Config{"unix:/tmp/sock", "!", "http://localhost", map[string]string{}}
}
