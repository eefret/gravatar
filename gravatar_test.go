package gravatar

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	myEmail string = "eefretsoul@gmail.com"
)

func TestGravatarNew(t *testing.T) {
	g, err := New()
	if err != nil {
		t.Error("An error has ocurred instantiating Gravatar")
	}
	assert.Equal(t, uint(80), g.size, "Size default should be 80")
}

func TestGravatarURLParse(t *testing.T) {
	g, err := New()
	if err != nil {
		t.Error("An error has ocurred instantiating Gravatar")
	}
	assert.Equal(t, "https://www.gravatar.com/avatar/c82739de14cf64affaf30856ca95b851.jpg?d=&r=g&s=80",
		g.URLParse(myEmail), "Should be equal")

	g.SetSize(uint(256))
	assert.Equal(t, "https://www.gravatar.com/avatar/c82739de14cf64affaf30856ca95b851.jpg?d=&r=g&s=256",
		g.URLParse(myEmail), "Should be equal")
}

func TestGravatarDownload(t *testing.T) {
	g, err := New()
	if err != nil {
		t.Error("An error has ocurred instantiating Gravatar")
	}
	data, err := g.Download(myEmail)
	assert.NotNil(t, data, "Should have something")
}

func TestGravatarDownloadToDisk(t *testing.T) {
	g, err := New()
	if err != nil {
		t.Error("An error has ocurred instantiating Gravatar")
	}
	g.SetSize(uint(512))
	g.DownloadToDisk(myEmail, "/tmp/avatar.jpg")
	file, err := ioutil.ReadFile("/tmp/avatar.jpg")
	assert.NoError(t, err, "should not be any errors")
	assert.NotNil(t, file, "should have something")
}
