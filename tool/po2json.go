package po2json

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"regexp"
)

type msgIdStrPairs map[string]string
type localesMsg map[string]msgIdStrPairs

const pattern = `msgid "(.+)"\nmsgstr "(.+)"`

func getPOPath(locale, domain, localedir string) string {
	filename := domain + ".po"
	return path.Join(localedir, locale, "LC_MESSAGES", filename)
}

func extractFromPOFile(popath string) msgIdStrPairs {
	buf, err := ioutil.ReadFile(popath)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(string(buf), -1)

	pairs := msgIdStrPairs{}
	for _, array := range matches {
		pairs[array[1]] = array[2]
	}
	return pairs
}

func PO2JSON(domain, localedir, jsonPath string) {
	dirs, err := ioutil.ReadDir(localedir)
	if err != nil {
		panic(err)
	}

	// create PO-like json data for i18n
	obj := localesMsg{}
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}
		locale := dir.Name()
		// English is default language
		if locale == "en_US" {
			continue
		}

		obj[locale] = extractFromPOFile(getPOPath(locale, domain, localedir))
	}

	b, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(jsonPath, b, 0644)
	if err != nil {
		panic(err)
	}
}
