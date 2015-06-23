package default_img

type DefaultImageType string

var DefaultImage = struct {
	GRAVATAR_ICON DefaultImageType
	IDENTICON     DefaultImageType
	MONSTERID     DefaultImageType
	WAVATAR       DefaultImageType
	HTTP_404      DefaultImageType
}{
	"",
	"identicon",
	"monsterid",
	"wavatar",
	"404",
}
