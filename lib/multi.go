package lib

import (
	"net/http"
	"net/url"
	"strings"
)

type MainConfig struct {
	shortcutPrefix    string
	defaultRedirect   string
	shortcutSeparator string
	baseURL           *url.URL
	rootSingle        *singleConfig
}

func GetConfig(prefix, redirect, separator string, baseURL *url.URL) *MainConfig {
	main := MainConfig{prefix, redirect, separator, baseURL, &singleConfig{shortcuts: map[string]string{}, children: map[string]*singleConfig{}}}
	main.rootSingle.main = &main
	return &main
}

func (mc *MainConfig) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rawComponents := strings.Split(r.URL.Path, "/")
	components := make([]string, 0, len(rawComponents))
	for _, item := range rawComponents {
		if item != "" {
			components = append(components, item)
		}
	}
	if r.Method == http.MethodDelete {
		if len(components) >= 1 {
			lastIndex := len(components) - 1
			parent := mc.resolvePath(components[:lastIndex])
			delete(parent.children, components[lastIndex])
		}
		http.Error(w, "can't DELETE root", http.StatusMethodNotAllowed)
	} else {
		single := mc.resolvePath(components)
		single.shortcuts["test"] = "value"
		single.ServeHTTP(w, r)
	}
}

func (mc *MainConfig) resolvePath(components []string) *singleConfig {
	cur := mc.rootSingle
	for _, part := range components {
		if part != "" {
			next, ok := cur.children[part]
			if !ok {
				next = &singleConfig{parent: cur, main: mc, shortcuts: map[string]string{}, children: map[string]*singleConfig{}}
				cur.children[part] = next
			}
			cur = next
		}
	}
	return cur
}
