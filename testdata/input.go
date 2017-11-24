package main

import "github.com/matthewmueller/deep-interface-bug/window"

func Element() window.Element {
	return nil
}

func main() {
	html := Element().(window.HTMLElement)
	html.DispatchEvent()
}
