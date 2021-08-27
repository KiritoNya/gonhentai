package nhentai

import (
	"fmt"
	"regexp"
)

// validateDoujinUrl is a function that checks if the url of doujinshi is valid.
func validateDoujinUrl(doujinUrl string) bool {
	// Check if it's a valid url
	match, _ := regexp.MatchString(
		`^https:\/\/(www.)?nhentai\.net\/g\/[0-9]{1,6}[\/]?$`,
		doujinUrl,
	)

	return match
}

// validateNhentaiId is a function that checks if the id of doujinshi is valid.
func validateNhentaiId(doujinId int) bool {
	doujinIdString := fmt.Sprintf("%d", doujinId)

	// Check if it's a valid nhentai id
	ok, _ := regexp.MatchString(`^[0-9]{1,6}$`, doujinIdString)

	return ok
}

// validateNhentaiImageUrl is a function that checks if the url of image is valid.
func validateNhentaiImageUrl(url string) bool {

	ok, _ := regexp.MatchString(
		`^https:\/\/(t|i)\.nhentai\.net\/(galleries|avatars)\/[0-9]+\/?.+\.(png|jpg|gif)?.+$`,
		url,
	)

	return ok
}

// validateNhentaiImageUrl is a function that checks if the url of comment is valid.
func validateCommentUrl(commentUrl string) bool {
	ok, _ := regexp.MatchString(
		`^https:\/\/(www.)?nhentai\.net\/g\/[0-9]{0,6}\/#comment-[0-9]+\/?$`,
		commentUrl,
	)

	return ok
}

// validateNhentaiImageUrl is a function that checks if the url of user is valid.
func validateUserUrl(userUrl string) bool {
	ok, _ := regexp.MatchString(
		`^https:\/\/(www.)?nhentai\.net\/users\/[0-9]+\/.+$`,
		userUrl,
	)

	return ok
}

// validateImageType is a function that checks if the image type is valid
func validateImageType(ext string) bool {
	if (ext != "jpg") && (ext != "png") && (ext != "gif") {
		return false
	}
	return true
}
