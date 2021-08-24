package nhentai

import (
	"net/http"
)

const (

	// BaseUrl is the url base of site
	BaseUrl string = "https://nhentai.net"

	// DoujinBaseUrl is the prefix that comes before the id in the doujinshi url
	DoujinBaseUrl string = "https://nhentai.net/g/"

	// ImageBaseUrl is base url for the image link
	ImageBaseUrl string = "https://i.nhentai.net"

	// ThumbnailBaseUrl is the base url for the thumbnail link
	ThumbnailBaseUrl string = "https://t.nhentai.net"

	// ImageTemplate is the default template for generate the image name in the Doujinshi.Save method
	DefaultImageNameTemplate string = "{{.pageNum}}.{{.ext}}"

	// BaseUrlApi is the url base api of the site
	baseUrlApi string = "https://nhentai.net/api"

	// galleryApi is the endpoint for get the doujinshi info
	galleryApi string = "/gallery/{{.id}}"

	// imageCompleteUrl is the url for get the  thumbnail image or page image
	imageCompleteUrl string = "{{.baseImageUrl}}/galleries/{{.mediaId}}/{{.numPage}}.{{.ext}}"

	// galleries/search?query=${query}&page=${page}&sort=${sort}
	// searchApi is the endpoint for get the result of research
	searchApi string = "/galleries/search"

	// galleries/tagged?tag_id=${id}&page=${page}${sort ? `&sort=${sort}
	// searchTagIdApi is the endpoint for get the result of tagId research
	searchTagIdApi string = "/galleries/tagged?{{ $count := 0 }}{{ if .id }}tag_id={{.id}}{{$count := add $count 1}}{{end}}{{if .page}}{{if $count gt 0}}&{{end}}page={{.page}}{{$count := add $count 1}}{{end}}{{if .sort}}{{if $count gt 0}}&{{end}}sort={{.sort}}{{end}}"

	// searchRelatedApi is the endpoint for get the related doujinshi
	searchRelatedApi string = "/gallery/{{.id}}/related"

	// commentsApi is the endpoint for get the comments linked to a doujinshi
	commentsApi string = "/gallery/{{.id}}/comments"
)

// ClientHttp is the client used for http requests. The default value is http.DefaultClient.
var ClientHttp *http.Client

func init() {
	ClientHttp = http.DefaultClient
}
