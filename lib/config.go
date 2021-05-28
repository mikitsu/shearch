package lib

type Config struct {
	shortcutPrefix    string
	defaultRedirect   string
	shortcutSeparator string
	opensearchUrl     string
	shortcuts         map[string]string
}

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

func GetConfig(shortcutPrefix, defaultRedirect, shortcutSeparator, opensearchUrl string) Config {
	return Config{shortcutPrefix, defaultRedirect, shortcutSeparator, opensearchUrl, map[string]string{}}
}
