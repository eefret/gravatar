# Gravatar for Golang

Gravatar is a tiny go library for accessing gravatar.com avatar images.

## Usage example

	import "github.com/eefret/gravatar"

	func main(){
		g, err := gravatar.New() //Creating our object to hold preferences: size, defaultImg, rating
		
		g.URLParse("eefretsoul@gmail.com") 
		// prints https://www.gravatar.com/avatar/c82739de14cf64affaf30856ca95b851.jpg?d=&r=g&s=80
		
		g.setSize(uint(256))
		data, err := g.Download("eefretsoul@gmail.com")
		// Returns []byte containing the Image with modified size to 256px
		
		g.setSize(uint(512))
		g.DownloadToDisk("eefretsoul@gmail.com", "/tmp/avatar.jpg")
		// Saves to /tmp our avatar image file with a size of 512
		
		if err != nil {
			log.Fatal(err)
		}
	}

## License

You can use gravatar under the terms of the Apache v2 license: