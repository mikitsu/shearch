package lib

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const opensearchTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<OpenSearchDescription xmlns="http://a9.com/-/spec/opensearch/1.1/">
  <ShortName>shearch</ShortName>
  <Description>search with shortcuts</Description>
  <Url type="text/html" template="%s?q={searchTerms}"/>
</OpenSearchDescription>`
const opensearchLinkHTML = `<!DOCTYPE html>
<html><head>
<link rel="search" title="shearch" type="application/opensearchdescription+xml" href="?opensearchxml">
</head></html>`

type singleConfig struct {
	shortcuts map[string]string
	children  map[string]*singleConfig
	parent    *singleConfig
	main      *MainConfig
}

func (conf *singleConfig) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	switch r.Method {
	case http.MethodGet, http.MethodHead:
		if _, ok := query["q"]; ok {
			conf.handleQuery(w, query)
		} else if conf.main.baseURL != nil {
			if _, ok := query["opensearch"]; ok {
				w.Header().Add("Content-Type", "text/html")
				w.Write([]byte(opensearchLinkHTML))
			} else if _, ok := query["opensearchxml"]; ok {
				completeURL := conf.main.baseURL.ResolveReference(r.URL)
				completeURL.RawQuery = ""
				w.Header().Add("Content-Type", "application/opensearchdescription+xml")
				w.Write([]byte(fmt.Sprintf(opensearchTemplate, completeURL.String())))
			}
		}
	case http.MethodPost:
		conf.handleShortcut(w, query)
	default:
		// DELETE is handled one level above
		http.Error(w, "GET/HEAD for redirects and opensearch, POST to update shortcuts, DELETE to, y'know, delete shortcuts", http.StatusNotImplemented)
	}
}

func (conf *singleConfig) handleQuery(w http.ResponseWriter, query url.Values) {
	q := query.Get("q")
	response := ""
	if strings.HasPrefix(q, conf.main.shortcutPrefix) {
		split_query := strings.SplitN(q[len(conf.main.shortcutPrefix):], conf.main.shortcutSeparator, 2)
		if len(split_query) == 2 {
			response = conf.resolveShortcut(split_query[0])
			if response != "" {
				q = split_query[1]
			}
		}
	}
	if response == "" {
		response = conf.main.defaultRedirect
	}
	if strings.Contains(response, "%s") {
		response = fmt.Sprintf(response, url.QueryEscape(q))
	} else {
		response = response + url.QueryEscape(q)
	}
	w.Header().Add("Location", response)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (conf *singleConfig) resolveShortcut(shortcut string) string {
	for key, value := range conf.shortcuts {
		if shortcut == key {
			return value
		}
	}
	if conf.parent == nil {
		return ""
	} else {
		return conf.parent.resolveShortcut(shortcut)
	}
}

func (conf *singleConfig) handleShortcut(w http.ResponseWriter, query url.Values) {
	for k, v := range query {
		newDest := v[0]
		conf.shortcuts[k] = newDest
	}
}
