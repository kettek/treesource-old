// +build !nogui

package treesource

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/url"

	"github.com/zserge/webview"
	"treesource/assets"
)

var w webview.WebView

func (a *App) Dispatch(e string, v interface{}) {
	if !app.gui {
		return
	}
	fmt.Printf("what the\n")
	w.Dispatch(func() {
		js, err := json.Marshal(v)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("jsApp.handleEvent('%s', %s)", template.JSEscapeString(e), string(js))
		w.Eval(fmt.Sprintf("jsApp.handleEvent('%s', %s)", template.JSEscapeString(e), string(js)))
	})
}

func RunGUI() error {
	app.gui = true

	myHTML, err := assets.Asset("html/app.html")
	if err != nil {
		return err
	}

	myCSS, err := assets.Asset("css/app.css")
	if err != nil {
		return err
	}

	myJS, err := assets.Asset("js/app.js")
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
	defer w.Exit()

	w.Bind("app", &app)
	// Inject CSS
	w.Eval(fmt.Sprintf(`(function(css){
		function init() {
   	  var style = document.createElement('style');
   	  var head = document.head || document.getElementsByTagName('head')[0];
   	  style.setAttribute('type', 'text/css');
   	  if (style.styleSheet) {
   	  	style.styleSheet.cssText = css;
   	  } else {
   	  	style.appendChild(document.createTextNode(css));
			}
			head.appendChild(style);
		}
		if (/Linux|MSIE|Trident/.test(window.navigator.userAgent)) {
			init()
		} else {
			window.addEventListener('DOMContentLoaded', init)
		}
  })("%s")`, template.JSEscapeString(string(myCSS))))

	// Inject JS
	w.Eval(string(myJS))

	w.Run()

	return nil
}
