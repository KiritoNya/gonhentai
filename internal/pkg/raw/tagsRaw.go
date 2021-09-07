package raw

import (
	"encoding/json"
	"strconv"
)

// TagRaw contains the json returned by the API for tags
type TagRaw struct {
	Data map[string]json.RawMessage
}

// Id is a function that returns the id of tags raw
func (tr *TagRaw) Id() (id int, err error) {
	err = json.Unmarshal(tr.Data["id"], &id)
	if err != nil {
		// Check if id is a string
		var idString string
		err2 := json.Unmarshal(tr.Data["id"], &idString)
		if err2 != nil {
			// It isn't a string, it's a real error
			return 0, err2
		}

		// Convert string to int
		id, err = strconv.Atoi(idString)
		if err2 != nil {
			return 0, err
		}
	}

	return id, nil
}

// Type is a function that returns the type of tags raw
func (tr *TagRaw) Type() (tagsType string, err error) {
	err = json.Unmarshal(tr.Data["type"], &tagsType)
	if err != nil {
		return "", err
	}
	return tagsType, nil
}

// Name is a function that returns the name of tags raw
func (tr *TagRaw) Name() (name string, err error) {
	err = json.Unmarshal(tr.Data["name"], &name)
	if err != nil {
		return "", err
	}
	return name, nil
}

// Url is a function that returns the url of tags raw
func (tr *TagRaw) Url() (url string, err error) {
	err = json.Unmarshal(tr.Data["url"], &url)
	if err != nil {
		return "", err
	}
	return url, nil
}

// Count is a function that returns the count of tags raw
func (tr *TagRaw) Count() (count int, err error) {
	err = json.Unmarshal(tr.Data["count"], &count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// All is a function that returns all info of tags in a map
func (tr *TagRaw) All() (tagMap map[string]interface{}, err error) {
	tagMap = make(map[string]interface{})

	// Get Id
	id, err := tr.Id()
	if err != nil {
		return nil, err
	}

	// Get type
	tagType, err := tr.Type()
	if err != nil {
		return nil, err
	}

	// Get Name
	name, err := tr.Name()
	if err != nil {
		return nil, err
	}

	// Get Url
	url, err := tr.Url()
	if err != nil {
		return nil, err
	}

	// Get Count
	count, err := tr.Count()
	if err != nil {
		return nil, err
	}

	// Fill map
	tagMap["Id"] = id
	tagMap["Type"] = tagType
	tagMap["Name"] = name
	tagMap["Url"] = url
	tagMap["Count"] = count

	return tagMap, nil
}
