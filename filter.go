package nhentai

// Filter is the data struct that describes a filter
type Filter struct {
	Id   int
	Name string
	Type TagsType
}

// toBeFilter is a function that determines whether the doujinshi should be filtered or not. Returns true if the doujinshi should be kept in the collection or false if it should be removed.
func toBeFilter(doujin *Doujinshi, filters []Filter) (bool, error) {

	// Check filters validity
	err := validateFilters(filters)
	if err != nil {
		return false, err
	}

	// Extract all tags from doujinshi object
	tags := doujin.extractTags()

	// Apply all filters
	foundTags := 0
	for _, filter := range filters {
		if findTag(tags, filter) {
			foundTags++
		}
	}

	// Check if doujinshi satisfies all filters
	if foundTags == len(filters) {
		// Doujinshi musn't be removed
		return true, nil
	}

	// Doujinshi must be removed
	return false, nil
}

// toBeDelete is a function that determines whether the doujinshi should be deleted or not. Returns true if the doujinshi should be deteleted or false if it shouldn't be deleted.
func toBeDelete(doujin *Doujinshi, filters []Filter) (bool, error) {
	// Check filters validity
	err := validateFilters(filters)
	if err != nil {
		return false, err
	}

	// Extract all tags from doujinshi object
	tags := doujin.extractTags()

	// Apply all filters
	for _, filter := range filters {
		if findTag(tags, filter) {
			return true, nil
		}
	}

	return false, nil
}

// findTag is a function that find if there is a tag that corresponds to the filter
func findTag(tags []*TagInfo, filter Filter) bool {
	// Declare a function that verify if the single tag is equal at filter
	tagCompare := func(tag *TagInfo, f Filter) bool {
		// Check id
		if f.Id != 0 && f.Id != tag.Id {
			//NotFound
			return false
		}

		// Check name
		if f.Name != "" && f.Name != tag.Name {
			return false
		}

		// Check tagsType
		if f.Type != "" && f.Type != tag.Type {
			return false
		}

		return true
	}

	// Foreach tag
	for _, tag := range tags {
		if tagCompare(tag, filter) == true {
			return true
		}
	}
	return false
}
