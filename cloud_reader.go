package main

import "github.com/zserge/webview"

func main() {
	webview.Open("Amazon Kindle",
		"https://read.amazon.com/",  800, 600, true)
}
