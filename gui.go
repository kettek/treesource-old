// +build !nogui

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/url"

	"github.com/zserge/webview"
)

var w webview.WebView

func (a *App) Dispatch(e string, v interface{}) {
	w.Dispatch(func() {
		js, err := json.Marshal(v)
		if err != nil {
			fmt.Println(err)
		}
		w.Eval(fmt.Sprintf("jsApp.handleEvent('%s', %s)", template.JSEscapeString(e), string(js)))
	})
}

func runGUI() error {
	myHTML, err := Asset("assets/app.html")
	if err != nil {
		return err
	}

	myCSS, err := Asset("assets/app.css")
	if err != nil {
		return err
	}

	myJS, err := Asset("assets/app.js")
	if err != nil {
		return err
	}

	w = webview.New(webview.Settings{
		Title:     "treesource",
		URL:       `data:text/html,` + url.PathEscape(string(myHTML)),
		Width:     512,
		Height:    640,
		Resizable: true,
		Debug:     true,
	})

	w.Dispatch(func() {
		w.Bind("app", &app)
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
    })("%s")`, template.JSEscapeString(string(myCSS))))

		// Inject JS
		w.Eval(string(myJS))
		//w.Eval(template.JSEscapeString(string(myJS)))
	})

	/*	w.Dispatch(func() {
		r := w.Dialog(webview.DialogTypeOpen, webview.DialogFlagDirectory, "Wat", "butts")
		fmt.Printf("%+v\n", r)
	})*/

	w.Run()

	return nil
}
