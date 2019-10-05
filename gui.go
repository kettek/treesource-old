package main

import (
	"fmt"
	"html/template"
	"net/url"

	"github.com/zserge/webview"
)

func runGUI() {
	const myHTML = `<!doctype html><html><head></head><body>yo</body></html>`

	const myStylesCSS = `
body {
  background: black;
  color: white;
}`

	w := webview.New(webview.Settings{
		Title:     "treesource",
		URL:       `data:text/html,` + url.PathEscape(myHTML),
		Resizable: true,
		Debug:     true,
	})

	w.Dispatch(func() {
		w.Bind("teststate", &TestState{})
		// Inject CSS
		w.Eval(fmt.Sprintf(`(function(css){
      var style = document.createElement('style');
      var head = document.head || document.getElementsByTagName('head')[0];
      style.setAttribute('type', 'text/css');
      if (style.styleSheet) {
      	style.styleSheet.cssText = css;
      } else {
      	style.appendChild(document.createTextNode(css));
      }
      head.appendChild(style);
    })("%s")`, template.JSEscapeString(myStylesCSS)))

		// Inject JS
		w.Eval(fmt.Sprintf(`
      //alert(teststate.data.value)
    `))
		//    w.Eval(myJSFramework)
		//    w.Eval(myAppJS)
	})

	w.Dispatch(func() {
		r := w.Dialog(webview.DialogTypeOpen, webview.DialogFlagDirectory, "Wat", "butts")
		fmt.Printf("%+v\n", r)
	})

	w.Run()

}
