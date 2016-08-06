package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var ArtworkType = Type("artworkType", func() {
	Description("Type of card artwork")
	Attribute("full", String, "Href to full size artwork")
	Attribute("artist", String, "Name of the artist")
	Attribute("small", String, "Href to thumbnail")

	Required("full", "small")
})