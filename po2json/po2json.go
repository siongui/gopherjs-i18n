// Package po2json converts PO files to JSON file. The JSON file is used in
// frontend code to translate the HTML document on demand.
package po2json

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"regexp"
)

// LocalesMsg is the data structrue to store translations of PO file.
type LocalesMsg map[string]MsgIdStrPairs

// MsgIdStrPairs is the data structrue to store translations of PO file.
type MsgIdStrPairs map[string]string

const pattern = `msgid "(.+)"\nmsgstr "(.+)"`

func getPOPath(locale, domain, localedir string) string {
	filename := domain + ".po"
	return path.Join(localedir, locale, "LC_MESSAGES", filename)
}

func extractFromPOFile(popath string) (pairs MsgIdStrPairs, err error) {
	buf, err := ioutil.ReadFile(popath)
	if err != nil {
		return
	}

	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(string(buf), -1)

	pairs = MsgIdStrPairs{}
	for _, array := range matches {
		pairs[array[1]] = array[2]
	}
	return
}

// PO2JSON converts PO files to JSON file. The JSON file is used in frontend
// code to translate the HTML document on demand.
func PO2JSON(domain, localedir, jsonPath string) (err error) {
	dirs, err := ioutil.ReadDir(localedir)
	if err != nil {
		return
	}

	// create PO-like json data for i18n
	obj := LocalesMsg{}
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}
		locale := dir.Name()
		// English is default language
		if locale == "en_US" {
			continue
		}

		obj[locale], err = extractFromPOFile(getPOPath(locale, domain, localedir))
		if err != nil {
			return
		}
	}

	b, err := json.Marshal(obj)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(jsonPath, b, 0644)
	return
}
