package main

import "github.com/webview/webview"

func main() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Amazon Kindle Reader")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("https://read.amazon.com/")
	w.Run()
}
