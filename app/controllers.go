// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/tri125/gwentapi/design
// --out=$(GOPATH)\src\github.com\tri125\gwentapi
// --version=v1.1.0-dirty
//
// API "gwentapi": Application Controllers
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// CardController is the controller interface for the Card actions.
type CardController interface {
	goa.Muxer
	CardFaction(*CardFactionCardContext) error
	CardLeader(*CardLeaderCardContext) error
	CardVariation(*CardVariationCardContext) error
	CardVariations(*CardVariationsCardContext) error
	List(*ListCardContext) error
	Show(*ShowCardContext) error
}

// MountCardController "mounts" a Card resource controller on the given service.
func MountCardController(service *goa.Service, ctrl CardController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCardFactionCardContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.CardFaction(rctx)
	}
	service.Mux.Handle("GET", "/v0/cards/factions/:factionID", ctrl.MuxHandler("CardFaction", h, nil))
	service.LogInfo("mount", "ctrl", "Card", "action", "CardFaction", "route", "GET /v0/cards/factions/:factionID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCardLeaderCardContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.CardLeader(rctx)
	}
	service.Mux.Handle("GET", "/v0/cards/leaders", ctrl.MuxHandler("CardLeader", h, nil))
	service.LogInfo("mount", "ctrl", "Card", "action", "CardLeader", "route", "GET /v0/cards/leaders")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCardVariationCardContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.CardVariation(rctx)
	}
	service.Mux.Handle("GET", "/v0/cards/:cardID/variations/:variationID", ctrl.MuxHandler("CardVariation", h, nil))
	service.LogInfo("mount", "ctrl", "Card", "action", "CardVariation", "route", "GET /v0/cards/:cardID/variations/:variationID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCardVariationsCardContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.CardVariations(rctx)
	}
	service.Mux.Handle("GET", "/v0/cards/:cardID/variations", ctrl.MuxHandler("CardVariations", h, nil))
	service.LogInfo("mount", "ctrl", "Card", "action", "CardVariations", "route", "GET /v0/cards/:cardID/variations")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListCardContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/v0/cards", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Card", "action", "List", "route", "GET /v0/cards")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowCardContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/v0/cards/:cardID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Card", "action", "Show", "route", "GET /v0/cards/:cardID")
}

// CategoryController is the controller interface for the Category actions.
type CategoryController interface {
	goa.Muxer
	List(*ListCategoryContext) error
	Show(*ShowCategoryContext) error
}

// MountCategoryController "mounts" a Category resource controller on the given service.
func MountCategoryController(service *goa.Service, ctrl CategoryController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListCategoryContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/v0/categories", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Category", "action", "List", "route", "GET /v0/categories")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowCategoryContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/v0/categories/:categoryID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Category", "action", "Show", "route", "GET /v0/categories/:categoryID")
}

// FactionController is the controller interface for the Faction actions.
type FactionController interface {
	goa.Muxer
	List(*ListFactionContext) error
	Show(*ShowFactionContext) error
}

// MountFactionController "mounts" a Faction resource controller on the given service.
func MountFactionController(service *goa.Service, ctrl FactionController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListFactionContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/v0/factions", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Faction", "action", "List", "route", "GET /v0/factions")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowFactionContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/v0/factions/:factionID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Faction", "action", "Show", "route", "GET /v0/factions/:factionID")
}

// GroupController is the controller interface for the Group actions.
type GroupController interface {
	goa.Muxer
	List(*ListGroupContext) error
	Show(*ShowGroupContext) error
}

// MountGroupController "mounts" a Group resource controller on the given service.
func MountGroupController(service *goa.Service, ctrl GroupController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListGroupContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/v0/groups", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Group", "action", "List", "route", "GET /v0/groups")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowGroupContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/v0/groups/:groupID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Group", "action", "Show", "route", "GET /v0/groups/:groupID")
}

// IndexController is the controller interface for the Index actions.
type IndexController interface {
	goa.Muxer
	Show(*ShowIndexContext) error
}

// MountIndexController "mounts" a Index resource controller on the given service.
func MountIndexController(service *goa.Service, ctrl IndexController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowIndexContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/v0", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Index", "action", "Show", "route", "GET /v0")
}

// RarityController is the controller interface for the Rarity actions.
type RarityController interface {
	goa.Muxer
	List(*ListRarityContext) error
	Show(*ShowRarityContext) error
}

// MountRarityController "mounts" a Rarity resource controller on the given service.
func MountRarityController(service *goa.Service, ctrl RarityController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListRarityContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/v0/rarities", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Rarity", "action", "List", "route", "GET /v0/rarities")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowRarityContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/v0/rarities/:rarityID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Rarity", "action", "Show", "route", "GET /v0/rarities/:rarityID")
}
