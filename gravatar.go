package gravatar

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/eefret/gravatar/default_img"
	"github.com/eefret/gravatar/rating"
)

const (
	gravatarURI string = "https://www.gravatar.com/avatar/"
	defaultSIZE uint   = 80
)

//Gravatar object
type Gravatar struct {
	rating       rating.RatingType
	defaultImage default_img.DefaultImageType
	size         uint
}

//New instantiates a new Gravatar Object
func New() (*Gravatar, error) {
	//Placing defaults
	t := &Gravatar{
		rating:       rating.Rating.GENERAL_AUDIENCES,
		defaultImage: default_img.DefaultImage.GRAVATAR_ICON,
		size:         defaultSIZE,
	}
	return t, nil
}

//Download will get the image and return the bytes
func (g *Gravatar) Download(email string) (data []byte, err error) {
	resp, err := http.Get(g.parseURI(emailToHash(email)).String())
	check(err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(http.StatusText(resp.StatusCode))
	}

	data, err = ioutil.ReadAll(resp.Body)
	check(err)
	return data, nil
}

//DownloadToDisk will store the avatar in the provided path
func (g *Gravatar) DownloadToDisk(email string, path string) {
	data, err := g.Download(email)
	check(err)
	ioutil.WriteFile(path, data, 0666)
}

//URLParse will parse the Email with your preferences and return an URL to the avatar
func (g *Gravatar) URLParse(email string) string {
	return g.parseURI(emailToHash(email)).String()
}

func (g *Gravatar) SetSize(size uint) {

	g.size = size
}

func (g *Gravatar) SetDefaultImage(defaultImage default_img.DefaultImageType) {
	g.defaultImage = defaultImage
}

func (g *Gravatar) SetRating(rating rating.RatingType) {
	g.rating = rating
}

func emailToHash(email string) string {
	hasher := md5.New()
	hasher.Write([]byte(email))
	return fmt.Sprintf("%v", hex.EncodeToString(hasher.Sum(nil)))
}

func (g *Gravatar) parseURI(hash string) *url.URL {
	var URL *url.URL
	URL, err := url.Parse(gravatarURI)
	check(err)
	URL.Path += hash + ".jpg"
	parameters := url.Values{}
	parameters.Add("s", fmt.Sprintf("%v", g.size))
	parameters.Add("r", fmt.Sprintf("%v", g.rating))
	parameters.Add("d", fmt.Sprintf("%v", g.defaultImage))
	URL.RawQuery = parameters.Encode()
	return URL
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
