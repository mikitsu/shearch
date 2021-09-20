Shearch -- search with shortcuts
================================

I'm a big fan of DuckDuckGo !bangs, but there are some I miss
(which don't meet the "useful for more than 500 people")
and some I'd like to point somewhere else
(again, in a way probably not really useful to others).

This is a simple redirector with configurable shortcuts.

Usage
-----

Run the binary with the options you need.
You can then issue requests to wherever you are listening:

- GET requests will examine the "q" parameter and handle it according to redirect rules
- GET requests without "q" paramter with an "opensearch" parameter will return an HTML page with a ``<link rel="search">``,
  which should add a search engine to browsers
- POST requests will update the redirect rules: every query paramter (in the URL)
  will add/update the corresponding shortcut.
- DELETE requests will remove saved shortcuts for the given path and all subpaths

Every path saves shortcuts separately and looks up undefined shortcuts in parent components recursively.

### how a shortcut is matched

When a query ("q" parameter of GET request) is received, check whether

- it starts with the shortcut prefix and
- the part between shortcut prefix and first occurence of separator
  is a known shortcut

If yes, redirect to where the shortcut points to.
If no, redirect to the default redirect.

When redirecting, the passed query, without prefix+shortcut+separator it matched,
will be %-formatted into the configured location (i.e. shortcut target or default redirect)
*if* it contains "%s" (in this case, remmber to escape "%" as "%%").
Otherwise, it is appended to the end.

Examples
--------

configured shortcuts: ``wpage => https://en.wikipedia.org/wiki/, other => https://example.com/%s/%%3D``

- GET with ``q=!wpage Golang`` will redirect to https://en.wikipedia.org/wiki/Golang
- GET with ``q=!other middle`` will redirect to https://example.com/middle/%3D
- GET with ``q=!unknown shortcut`` will redirect to https://duckduckgo.com/?q=%21unknown+shortcut
- GET with ``q=not a shortcut`` will redirect to https://duckduckgo.com/?q=not+a+shortcut
- POST with ``other=https://example.net/`` will replace the shortcut
