// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/tri125/gwentapi/design
// --out=$(GOPATH)\src\github.com\tri125\gwentapi
// --version=v1.1.0-dirty
//
// API "gwentapi": Application Resource Href Factories
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"fmt"
	"strings"
)

// CardHref returns the resource href.
func CardHref(cardID interface{}) string {
	paramcardID := strings.TrimLeftFunc(fmt.Sprintf("%v", cardID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/v0/cards/%v", paramcardID)
}

// CategoryHref returns the resource href.
func CategoryHref(categoryID interface{}) string {
	paramcategoryID := strings.TrimLeftFunc(fmt.Sprintf("%v", categoryID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/v0/categories/%v", paramcategoryID)
}

// FactionHref returns the resource href.
func FactionHref(factionID interface{}) string {
	paramfactionID := strings.TrimLeftFunc(fmt.Sprintf("%v", factionID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/v0/factions/%v", paramfactionID)
}

// GroupHref returns the resource href.
func GroupHref(groupID interface{}) string {
	paramgroupID := strings.TrimLeftFunc(fmt.Sprintf("%v", groupID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/v0/groups/%v", paramgroupID)
}

// IndexHref returns the resource href.
func IndexHref() string {
	return "/v0"
}

// RarityHref returns the resource href.
func RarityHref(rarityID interface{}) string {
	paramrarityID := strings.TrimLeftFunc(fmt.Sprintf("%v", rarityID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/v0/rarities/%v", paramrarityID)
}
