package main

import (
	"github.com/gopherjs/gopherjs/js"
	jsgettext "github.com/siongui/gopherjs-i18n"
)

const jsonData = `{"vi_VN":{"About":"Giới thiệu","Canon":"Kinh điển","Home":"Trang chính","Setting":"Thiết lập","Translation":"Dịch"},"zh_TW":{"About":"關於","Canon":"經典","Home":"首頁","Setting":"設定","Translation":"翻譯"}}`

func main() {
	jsgettext.SetupTranslationMapping(jsonData)

	d := js.Global.Get("document")
	nodeList := d.Call("querySelectorAll", "button")
	length := nodeList.Get("length").Int()
	for i := 0; i < length; i++ {
		btn := nodeList.Call("item", i)
		btn.Call("addEventListener", "click", func() {
			jsgettext.TranslateDocument(btn.Get("value").String())
		})
	}
}
