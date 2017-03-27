package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/tri125/gwentapi/app"
	"github.com/tri125/gwentapi/dataLayer/dal"
	"github.com/tri125/gwentapi/dataLayer/factory"
	"github.com/tri125/gwentapi/helpers"
)

// CardController implements the card resource.
type CardController struct {
	*goa.Controller
}

// NewCardController creates a card controller.
func NewCardController(service *goa.Service) *CardController {
	return &CardController{Controller: service.NewController("CardController")}
}

// CardFaction runs the cardFaction action.
func (c *CardController) CardFaction(ctx *app.CardFactionCardContext) error {
	// CardController_CardFaction: start_implement
	dataStore := &dal.DataStore{}
	dataStore.GetSession()
	// Close the session
	defer dataStore.Close()
	dc := dal.NewDalCard(dataStore)
	df := dal.NewDalFaction(dataStore)
	factionUUID, errFactionUUID := helpers.Base64ToUUID(ctx.FactionID)

	if errFactionUUID != nil {
		return ctx.NotFound()
	}
	faction, errFaction := df.Fetch(factionUUID)

	collectionCount, err := dc.CountFromFaction(faction.ID)

	if err != nil {
		return ctx.NotFound()
	}
	limit, offset := helpers.ValidateLimitOffset(collectionCount, ctx.Limit, ctx.Offset)
	cards, err := dc.FetchFromFactionPaging(faction.ID, limit, offset)

	if err != nil || errFaction != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Card", "action", "CardFaction", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError", err.Error())
		return ctx.InternalServerError()
	}
	// CardController_CardFaction: end_implement
	res, _ := factory.CreatePageCard(cards, collectionCount, limit, offset)
	return ctx.OK(res)
}

// CardLeader runs the cardLeader action.
func (c *CardController) CardLeader(ctx *app.CardLeaderCardContext) error {
	// CardController_CardLeader: start_implement

	dataStore := &dal.DataStore{}
	dataStore.GetSession()
	// Close the session
	defer dataStore.Close()
	dc := dal.NewDalCard(dataStore)
	dg := dal.NewDalGroup(dataStore)

	group, errGroup := dg.FetchWithName("Leader")
	collectionCount, err := dc.CountLeader(group.ID)

	if err != nil {
		return ctx.NotFound()
	}
	limit, offset := helpers.ValidateLimitOffset(collectionCount, ctx.Limit, ctx.Offset)
	cards, err := dc.FetchLeaderPaging(group.ID, limit, offset)
	if err != nil || errGroup != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Card", "action", "CardLeader", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError", err.Error())
		return ctx.InternalServerError()
	}
	// CardController_CardLeader: end_implement
	res, _ := factory.CreatePageCard(cards, collectionCount, limit, offset)
	return ctx.OK(res)
}

// CardVariation runs the cardVariation action.
func (c *CardController) CardVariation(ctx *app.CardVariationCardContext) error {
	// CardController_CardVariation: start_implement
	dataStore := &dal.DataStore{}
	dataStore.GetSession()
	// Close the session
	defer dataStore.Close()
	dc := dal.NewDalVariation(dataStore)
	uuid, err := helpers.Base64ToUUID(ctx.CardID)
	variationUUID, errVariation := helpers.Base64ToUUID(ctx.VariationID)

	if err != nil || errVariation != nil {
		return ctx.NotFound()
	}

	variation, err := dc.Fetch(variationUUID)

	if err != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Card", "action", "CardVariation", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError", err.Error())
		return ctx.InternalServerError()
	}

	// CardController_CardVariation: end_implement
	res, _ := factory.CreateVariation(variation, uuid)
	return ctx.OK(res)
}

// CardVariations runs the cardVariations action.
func (c *CardController) CardVariations(ctx *app.CardVariationsCardContext) error {
	// CardController_CardVariations: start_implement
	dataStore := &dal.DataStore{}
	dataStore.GetSession()
	// Close the session
	defer dataStore.Close()
	dc := dal.NewDalCard(dataStore)
	dv := dal.NewDalVariation(dataStore)
	uuid, err := helpers.Base64ToUUID(ctx.CardID)

	if err != nil {
		return ctx.NotFound()
	}

	card, err := dc.Fetch(uuid)
	variations, errVariation := dv.FetchFromCardID(card.ID)
	if err != nil || errVariation != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Card", "action", "CardVariations", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError", err.Error())
		return ctx.InternalServerError()
	}

	// CardController_CardVariations: end_implement
	res, _ := factory.CreateLinkVariationCollection(variations, card.UUID)
	return ctx.OK(res)
}

// List runs the list action.
func (c *CardController) List(ctx *app.ListCardContext) error {
	// CardController_List: start_implement
	dataStore := &dal.DataStore{}
	dataStore.GetSession()
	// Close the session
	defer dataStore.Close()
	dc := dal.NewDalCard(dataStore)
	collectionCount, err := dc.Count()

	if err != nil {
		return ctx.NotFound()
	}
	limit, offset := helpers.ValidateLimitOffset(collectionCount, ctx.Limit, ctx.Offset)
	cards, err := dc.FetchAllPaging(limit, offset)
	if err != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Card", "action", "List", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError", err.Error())
		return ctx.InternalServerError()
	}
	// CardController_List: end_implement
	res, _ := factory.CreatePageCard(cards, collectionCount, limit, offset)
	return ctx.OK(res)
}

// Show runs the show action.
func (c *CardController) Show(ctx *app.ShowCardContext) error {
	// CardController_Show: start_implement
	dataStore := &dal.DataStore{}
	dataStore.GetSession()
	// Close the session
	defer dataStore.Close()
	dc := dal.NewDalCard(dataStore)
	dv := dal.NewDalVariation(dataStore)
	uuid, err := helpers.Base64ToUUID(ctx.CardID)

	if err != nil {
		return ctx.NotFound()
	}

	card, err := dc.Fetch(uuid)
	variations, errVariation := dv.FetchFromCardID(card.ID)
	if err != nil || errVariation != nil {
		ctx.ResponseData.Service.LogError("InternalServerError", "req_id", middleware.ContextRequestID(ctx), "ctrl", "Card", "action", "Show", ctx.RequestData.Request.Method, ctx.RequestData.Request.URL, "databaseError", err.Error())
		return ctx.InternalServerError()
	}

	// CardController_Show: end_implement
	res, _ := factory.CreateCard(card, variations)
	return ctx.OK(res)
}
