package jsgettext

import "github.com/gopherjs/gopherjs/js"
import "encoding/json"

type msgIdStrPairs map[string]string
type localesMsg map[string]msgIdStrPairs

var msg = localesMsg{}
var document = js.Global.Get("document")

func SetupTranslationMapping(jsonBlob string) {
	err := json.Unmarshal([]byte(jsonBlob), &msg)
	if err != nil {
		panic(err)
	}
}

func Gettext(locale, str string) string {
	if val, ok := msg[locale]; ok {
		if val2, ok2 := val[str]; ok2 {
			return val2
		}
	}
	return str
}

func Translate(locale string) {
	nodeList := document.Call("querySelectorAll", "[data-default-string]")
	length := nodeList.Get("length").Int()
	for i := 0; i < length; i++ {
		element := nodeList.Call("item", i)
		str := element.Get("dataset").Get("defaultString").String()
		element.Set("textContent", Gettext(locale, str))
	}
}
