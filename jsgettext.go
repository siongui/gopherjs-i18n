// Package jsgettext provides frontend methods to translate html document in the
// browser on demand, similar to GNU gettext function. This package is written
// in Go and compiled to JavaScript via GopherJS.
package jsgettext

import (
	"encoding/json"

	"github.com/gopherjs/gopherjs/js"
)

// LocalesMsg is the data structrue to store translations of PO file.
type LocalesMsg map[string]MsgIdStrPairs

// MsgIdStrPairs is the data structrue to store translations of PO file.
type MsgIdStrPairs map[string]string

var msg = LocalesMsg{}
var document = js.Global.Get("document")

func SetupTranslationMapping(jsonBlob string) (err error) {
	err = json.Unmarshal([]byte(jsonBlob), &msg)
	return
}

// Gettext translates the given string to the language specified by locale.
func Gettext(locale, str string) string {
	if val, ok := msg[locale]; ok {
		if val2, ok2 := val[str]; ok2 {
			return val2
		}
	}
	return str
}

// TranslateDocument searches all element with [data-default-string] attribute
// in the document, translate the string with given locale, and set the
// textContent of the element to the translated string.
func TranslateDocument(locale string) {
	nodeList := document.Call("querySelectorAll", "[data-default-string]")
	length := nodeList.Get("length").Int()
	for i := 0; i < length; i++ {
		element := nodeList.Call("item", i)
		str := element.Get("dataset").Get("defaultString").String()
		element.Set("textContent", Gettext(locale, str))
	}
}
