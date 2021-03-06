package factory

import (
	"github.com/GwentAPI/gwentapi/app"
	"github.com/GwentAPI/gwentapi/dataLayer/dal"
	"github.com/GwentAPI/gwentapi/dataLayer/models"
	"github.com/GwentAPI/gwentapi/helpers"
	"time"
)

func CreateCard(c *models.Card, ds *dal.DataStore, locale string) (*app.GwentapiCard, error) {
	uuid := helpers.EncodeUUID(c.UUID)
	dalCat := dal.NewDalCategory(ds)
	dalV := dal.NewDalVariation(ds)
	dalG := dal.NewDalGroup(ds)
	dalF := dal.NewDalFaction(ds)

	group, errG := dalG.FetchWithName(c.Group)
	if errG != nil {
		return nil, errG
	}
	groupMedia, _ := CreateGroupLink(group)

	faction, errF := dalF.FetchWithName(c.Faction)
	if errF != nil {
		return nil, errF
	}
	factionMedia, _ := CreateFactionLink(faction)

	categories, errC := dalCat.FetchFromArrayID(c.Categories_id)
	if errC != nil {
		return nil, errC
	}
	categoriesLinkMedia := make(app.GwentapiCategoryLinkCollection, len(*categories))
	for i, category := range *categories {
		cl, _ := CreateCategoryLink(&category)
		categoriesLinkMedia[i] = cl
	}

	variations, errV := dalV.FetchFromCardID(c.ID)
	if errV != nil {
		return nil, errV
	}
	variationsLinkMedia := make(app.GwentapiVariationLinkCollection, len(*variations))
	for i, variation := range *variations {
		cv, _ := CreateVariationLink(&variation, c.UUID, ds)
		variationsLinkMedia[i] = cv
	}

	name := c.Name[locale]
	flavor := c.Flavor[locale]
	info := c.Info[locale]

	result := &app.GwentapiCard{
		Name:       name,
		Categories: categoriesLinkMedia,
		Flavor:     &flavor,
		Info:       &info,
		Strength:   c.Strength,
		Positions:  c.Positions,
		Faction:    factionMedia,
		Group:      groupMedia,
		Variations: variationsLinkMedia,
		Href:       helpers.CardURL(uuid),
		UUID:       uuid,
	}

	return result, nil
}

func CreatePageCard(c *[]models.Card, url string, resultCount int, limit int, offset int, locale string) (*app.GwentapiPagecard, time.Time, error) {
	results := make(app.GwentapiCardLinkCollection, len(*c))
	lastModified := time.Time{}

	for i, result := range *c {
		uuid := helpers.EncodeUUID(result.UUID)
		cl := &app.GwentapiCardLink{
			Href: helpers.CardURL(uuid),
			Name: result.Name[locale],
		}

		if lastModified.Before(result.Last_Modified) {
			lastModified = result.Last_Modified
		}

		results[i] = cl
	}

	prev, next := helpers.GeneratePrevNextPageHref(resultCount, limit, offset, helpers.CardURL(url))

	page := &app.GwentapiPagecard{
		Count:    resultCount,
		Next:     next,
		Previous: prev,
		Results:  results,
	}
	return page, lastModified, nil
}
