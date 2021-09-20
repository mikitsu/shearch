#!/bin/sh

url=http://localhost:1234

do_test() {
    "$@" && echo pass || echo fail "$@"
}

check_redirect() {
    curl --silent --head "$url/$1" | do_test grep --quiet "Location: $2"
}

go run main.go -listen localhost:1234 -url http://localhost:1234 &

echo waiting for startup
while ! curl --silent "$url";
do
    sleep 1
done

check_redirect '?q=123' 'https://duckduckgo.com/?q=123'
curl -X POST "$url/first?a=pref-a-&b=pref-b-"
check_redirect '?q=%21a+test' 'https://duckduckgo.com/?q=%21a+test'
check_redirect '/first?q=%21a+test' 'pref-a-test'
check_redirect '/first?q=%21b+test' 'pref-b-test'
curl -X POST "$url//second//?b=pref2-b-"
check_redirect '/first?q=%21b+test' 'pref-b-test'
check_redirect '/second?q=%21b+test' 'pref2-b-test'
check_redirect '/first/sub?q=%21b+test' 'pref-b-test'

kill $(netstat -tlp 2>/dev/null | grep -Po 'LISTEN\s+\K\d+(?=/main)')
