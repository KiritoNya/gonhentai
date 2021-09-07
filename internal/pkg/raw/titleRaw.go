package raw

import (
	"encoding/json"
)

// TitleRaw is the json api result of the doujinshi title
type TitleRaw struct {
	Data map[string]json.RawMessage
}

// English is a function that returns the english title of the Title Raw
func (t *TitleRaw) English() (string, error) {
	return t.title("english")
}

// Japanese is a function that returns the japanese title of the Title Raw
func (t *TitleRaw) Japanese() (string, error) {
	return t.title("japanese")
}

// Pretty is a function that returns the pretty title of the Title Raw
func (t *TitleRaw) Pretty() (string, error) {
	return t.title("pretty")
}

// All is a function that returns a title map with title in all languages
func (t *TitleRaw) All() (returnMap map[string]string, err error) {
	returnMap = make(map[string]string)

	// Get english title
	english, err := t.English()
	if err != nil {
		return nil, err
	}

	// Get japanese title
	japanese, err := t.Japanese()
	if err != nil {
		return nil, err
	}

	// Get pretty title
	pretty, err := t.Pretty()
	if err != nil {
		return nil, err
	}

	// Fill map
	returnMap["English"] = english
	returnMap["Japanese"] = japanese
	returnMap["Pretty"] = pretty

	return returnMap, nil
}

// title is a function that returns the english, japanese or pretty title of the Title Raw
func (t *TitleRaw) title(lang string) (title string, err error) {
	// Unmarshal
	err = json.Unmarshal(t.Data[lang], &title)
	if err != nil {
		return "", nil
	}

	if lang != "english" {
		title = decodeUnicodeString(title)
	}

	return title, err
}
