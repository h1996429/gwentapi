// Code generated by goagen v1.2.0, DO NOT EDIT.
//
// API "gwentapi": Application User Types
//
// Command:
// $ goagen
// --design=github.com/GwentAPI/gwentapi/design
// --out=$(GOPATH)\src\github.com\GwentAPI\gwentapi
// --version=v1.2.0

package app

import (
	"github.com/goadesign/goa"
)

// Type of card art
type artType struct {
	// Name of the artist
	Artist *string `form:"artist,omitempty" json:"artist,omitempty" xml:"artist,omitempty"`
	// Href to full size artwork
	FullsizeImage *string `form:"fullsizeImage,omitempty" json:"fullsizeImage,omitempty" xml:"fullsizeImage,omitempty"`
	// Href to thumbnail size artwork
	ThumbnailImage *string `form:"thumbnailImage,omitempty" json:"thumbnailImage,omitempty" xml:"thumbnailImage,omitempty"`
}

// Validate validates the artType type instance.
func (ut *artType) Validate() (err error) {
	if ut.ThumbnailImage == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "thumbnailImage"))
	}
	return
}

// Publicize creates ArtType from artType
func (ut *artType) Publicize() *ArtType {
	var pub ArtType
	if ut.Artist != nil {
		pub.Artist = ut.Artist
	}
	if ut.FullsizeImage != nil {
		pub.FullsizeImage = ut.FullsizeImage
	}
	if ut.ThumbnailImage != nil {
		pub.ThumbnailImage = *ut.ThumbnailImage
	}
	return &pub
}

// Type of card art
type ArtType struct {
	// Name of the artist
	Artist *string `form:"artist,omitempty" json:"artist,omitempty" xml:"artist,omitempty"`
	// Href to full size artwork
	FullsizeImage *string `form:"fullsizeImage,omitempty" json:"fullsizeImage,omitempty" xml:"fullsizeImage,omitempty"`
	// Href to thumbnail size artwork
	ThumbnailImage string `form:"thumbnailImage" json:"thumbnailImage" xml:"thumbnailImage"`
}

// Validate validates the ArtType type instance.
func (ut *ArtType) Validate() (err error) {
	if ut.ThumbnailImage == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "thumbnailImage"))
	}
	return
}

// Type used to define the associated crafting/milling cost cost
type costType struct {
	// Normal cost
	Normal *int `form:"normal,omitempty" json:"normal,omitempty" xml:"normal,omitempty"`
	// Premium cost
	Premium *int `form:"premium,omitempty" json:"premium,omitempty" xml:"premium,omitempty"`
}

// Validate validates the costType type instance.
func (ut *costType) Validate() (err error) {
	if ut.Premium == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "premium"))
	}
	if ut.Normal == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "normal"))
	}
	return
}

// Publicize creates CostType from costType
func (ut *costType) Publicize() *CostType {
	var pub CostType
	if ut.Normal != nil {
		pub.Normal = *ut.Normal
	}
	if ut.Premium != nil {
		pub.Premium = *ut.Premium
	}
	return &pub
}

// Type used to define the associated crafting/milling cost cost
type CostType struct {
	// Normal cost
	Normal int `form:"normal" json:"normal" xml:"normal"`
	// Premium cost
	Premium int `form:"premium" json:"premium" xml:"premium"`
}

// Validate validates the CostType type instance.
func (ut *CostType) Validate() (err error) {

	return
}
