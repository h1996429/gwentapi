//************************************************************************//
// API "gwentapi": Application Contexts
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/tri125/gwentapi/design
// --out=$(GOPATH)\src\github.com\tri125\gwentapi
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
)

// CardFactionCardContext provides the card cardFaction action context.
type CardFactionCardContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	FactionID string
}

// NewCardFactionCardContext parses the incoming request URL and body, performs validations and creates the
// context used by the card controller cardFaction action.
func NewCardFactionCardContext(ctx context.Context, service *goa.Service) (*CardFactionCardContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CardFactionCardContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramFactionID := req.Params["factionID"]
	if len(paramFactionID) > 0 {
		rawFactionID := paramFactionID[0]
		rctx.FactionID = rawFactionID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CardFactionCardContext) OK(r GwentapiCardCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.card+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *CardFactionCardContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CardFactionCardContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// CardLeaderCardContext provides the card cardLeader action context.
type CardLeaderCardContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewCardLeaderCardContext parses the incoming request URL and body, performs validations and creates the
// context used by the card controller cardLeader action.
func NewCardLeaderCardContext(ctx context.Context, service *goa.Service) (*CardLeaderCardContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CardLeaderCardContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CardLeaderCardContext) OK(r GwentapiCardCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.card+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *CardLeaderCardContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CardLeaderCardContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// CardRarityCardContext provides the card cardRarity action context.
type CardRarityCardContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	RarityID string
}

// NewCardRarityCardContext parses the incoming request URL and body, performs validations and creates the
// context used by the card controller cardRarity action.
func NewCardRarityCardContext(ctx context.Context, service *goa.Service) (*CardRarityCardContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CardRarityCardContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramRarityID := req.Params["rarityID"]
	if len(paramRarityID) > 0 {
		rawRarityID := paramRarityID[0]
		rctx.RarityID = rawRarityID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CardRarityCardContext) OK(r GwentapiCardCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.card+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *CardRarityCardContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CardRarityCardContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ListCardContext provides the card list action context.
type ListCardContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListCardContext parses the incoming request URL and body, performs validations and creates the
// context used by the card controller list action.
func NewListCardContext(ctx context.Context, service *goa.Service) (*ListCardContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListCardContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListCardContext) OK(r GwentapiCardCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.card+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListCardContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListCardContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowCardContext provides the card show action context.
type ShowCardContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	CardID string
}

// NewShowCardContext parses the incoming request URL and body, performs validations and creates the
// context used by the card controller show action.
func NewShowCardContext(ctx context.Context, service *goa.Service) (*ShowCardContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowCardContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramCardID := req.Params["cardID"]
	if len(paramCardID) > 0 {
		rawCardID := paramCardID[0]
		rctx.CardID = rawCardID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowCardContext) OK(r *GwentapiCard) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.card+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowCardContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowCardContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ListFactionContext provides the faction list action context.
type ListFactionContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListFactionContext parses the incoming request URL and body, performs validations and creates the
// context used by the faction controller list action.
func NewListFactionContext(ctx context.Context, service *goa.Service) (*ListFactionContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListFactionContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListFactionContext) OK(r GwentapiFactionCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.faction+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListFactionContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListFactionContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowFactionContext provides the faction show action context.
type ShowFactionContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	FactionID string
}

// NewShowFactionContext parses the incoming request URL and body, performs validations and creates the
// context used by the faction controller show action.
func NewShowFactionContext(ctx context.Context, service *goa.Service) (*ShowFactionContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowFactionContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramFactionID := req.Params["factionID"]
	if len(paramFactionID) > 0 {
		rawFactionID := paramFactionID[0]
		rctx.FactionID = rawFactionID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowFactionContext) OK(r *GwentapiFaction) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.faction+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowFactionContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowFactionContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ListGlyphContext provides the glyph list action context.
type ListGlyphContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListGlyphContext parses the incoming request URL and body, performs validations and creates the
// context used by the glyph controller list action.
func NewListGlyphContext(ctx context.Context, service *goa.Service) (*ListGlyphContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListGlyphContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListGlyphContext) OK(r GwentapiGlyphCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.glyph+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListGlyphContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListGlyphContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowGlyphContext provides the glyph show action context.
type ShowGlyphContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	GlyphID string
}

// NewShowGlyphContext parses the incoming request URL and body, performs validations and creates the
// context used by the glyph controller show action.
func NewShowGlyphContext(ctx context.Context, service *goa.Service) (*ShowGlyphContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowGlyphContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramGlyphID := req.Params["glyphID"]
	if len(paramGlyphID) > 0 {
		rawGlyphID := paramGlyphID[0]
		rctx.GlyphID = rawGlyphID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowGlyphContext) OK(r *GwentapiGlyph) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.glyph+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowGlyphContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowGlyphContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// LatestPatchContext provides the patch latest action context.
type LatestPatchContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewLatestPatchContext parses the incoming request URL and body, performs validations and creates the
// context used by the patch controller latest action.
func NewLatestPatchContext(ctx context.Context, service *goa.Service) (*LatestPatchContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := LatestPatchContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *LatestPatchContext) OK(r *GwentapiPatch) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.patch+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKFull sends a HTTP response with status code 200.
func (ctx *LatestPatchContext) OKFull(r *GwentapiPatchFull) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.patch+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *LatestPatchContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *LatestPatchContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ListPatchContext provides the patch list action context.
type ListPatchContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListPatchContext parses the incoming request URL and body, performs validations and creates the
// context used by the patch controller list action.
func NewListPatchContext(ctx context.Context, service *goa.Service) (*ListPatchContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListPatchContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListPatchContext) OK(r GwentapiPatchCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.patch+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKFull sends a HTTP response with status code 200.
func (ctx *ListPatchContext) OKFull(r GwentapiPatchFullCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.patch+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListPatchContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListPatchContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowPatchContext provides the patch show action context.
type ShowPatchContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	PatchID string
}

// NewShowPatchContext parses the incoming request URL and body, performs validations and creates the
// context used by the patch controller show action.
func NewShowPatchContext(ctx context.Context, service *goa.Service) (*ShowPatchContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowPatchContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramPatchID := req.Params["patchID"]
	if len(paramPatchID) > 0 {
		rawPatchID := paramPatchID[0]
		rctx.PatchID = rawPatchID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowPatchContext) OK(r *GwentapiPatch) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.patch+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKFull sends a HTTP response with status code 200.
func (ctx *ShowPatchContext) OKFull(r *GwentapiPatchFull) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.patch+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowPatchContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowPatchContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowPhonebookContext provides the phonebook show action context.
type ShowPhonebookContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewShowPhonebookContext parses the incoming request URL and body, performs validations and creates the
// context used by the phonebook controller show action.
func NewShowPhonebookContext(ctx context.Context, service *goa.Service) (*ShowPhonebookContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowPhonebookContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowPhonebookContext) OK(r *GwentapiResource) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.resource+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowPhonebookContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListRarityContext provides the rarity list action context.
type ListRarityContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListRarityContext parses the incoming request URL and body, performs validations and creates the
// context used by the rarity controller list action.
func NewListRarityContext(ctx context.Context, service *goa.Service) (*ListRarityContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListRarityContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListRarityContext) OK(r GwentapiRarityCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.rarity+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListRarityContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListRarityContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowRarityContext provides the rarity show action context.
type ShowRarityContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	RarityID string
}

// NewShowRarityContext parses the incoming request URL and body, performs validations and creates the
// context used by the rarity controller show action.
func NewShowRarityContext(ctx context.Context, service *goa.Service) (*ShowRarityContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowRarityContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramRarityID := req.Params["rarityID"]
	if len(paramRarityID) > 0 {
		rawRarityID := paramRarityID[0]
		rctx.RarityID = rawRarityID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowRarityContext) OK(r *GwentapiRarity) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.rarity+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowRarityContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowRarityContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ListTypeContext provides the type list action context.
type ListTypeContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListTypeContext parses the incoming request URL and body, performs validations and creates the
// context used by the type controller list action.
func NewListTypeContext(ctx context.Context, service *goa.Service) (*ListTypeContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListTypeContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListTypeContext) OK(r GwentapiTypeCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.type+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListTypeContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListTypeContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowTypeContext provides the type show action context.
type ShowTypeContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	TypeID string
}

// NewShowTypeContext parses the incoming request URL and body, performs validations and creates the
// context used by the type controller show action.
func NewShowTypeContext(ctx context.Context, service *goa.Service) (*ShowTypeContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowTypeContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramTypeID := req.Params["typeID"]
	if len(paramTypeID) > 0 {
		rawTypeID := paramTypeID[0]
		rctx.TypeID = rawTypeID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowTypeContext) OK(r *GwentapiType) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.type+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowTypeContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowTypeContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}