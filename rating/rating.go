package rating

type RatingType string

var Rating = struct {
	GENERAL_AUDIENCES           RatingType
	PARENTAL_GUIDANCE_SUGGESTED RatingType
	RESTRICTED                  RatingType
	XPLICIT                     RatingType
}{
	"g",
	"pg",
	"r",
	"x",
}
