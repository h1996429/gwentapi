package dal

import (
	"github.com/GwentAPI/gwentapi/dataLayer/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"regexp"
)

type (
	DalCard struct {
		collection *mgo.Collection
	}

	CardQuery struct {
		Name string
		Lang string
	}
)

func NewDalCard(ds *DataStore) *DalCard {
	return &DalCard{ds.Collection("cards")}
}

func (dc DalCard) Fetch(uuid []byte) (*models.Card, error) {
	result := models.Card{}
	err := dc.collection.Find(bson.M{"uuid": uuid}).One(&result)
	return &result, err
}

func (dc DalCard) FetchAll() (*[]models.Card, error) {
	results := []models.Card{}
	err := dc.collection.Find(nil).Sort("name.en-US").All(&results)
	return &results, err
}

func (dc DalCard) FetchAllPaging(limit int, offset int) (*[]models.Card, int, error) {
	results := []models.Card{}
	query := dc.collection.Find(nil).Limit(limit).Sort("name.en-US").Skip(offset)
	err := query.All(&results)
	count, _ := dc.collection.Find(nil).Count()
	// db driver is bugged
	//count, _ := query.Count()
	return &results, count, err
}

func (dc DalCard) FetchQueryPaging(limit int, offset int, cardQuery CardQuery) (*[]models.Card, int, error) {
	results := []models.Card{}
	pattern := `^` + regexp.QuoteMeta(cardQuery.Name)
	fieldToSearch := "name." + cardQuery.Lang

	query := dc.collection.Find(bson.M{fieldToSearch: bson.RegEx{pattern, "i"}}).Limit(limit).Sort(fieldToSearch).Skip(offset)
	err := query.All(&results)
	count, _ := dc.collection.Find(bson.M{fieldToSearch: bson.RegEx{pattern, "i"}}).Count()
	// db driver is bugged
	//count, _ := query.Count()
	return &results, count, err
}

func (dc DalCard) FetchLeaderPaging(groupID bson.ObjectId, limit int, offset int) (*[]models.Card, error) {
	results := []models.Card{}
	err := dc.collection.Find(bson.M{"group_id": groupID}).Limit(limit).Sort("name.en-US").Skip(offset).All(&results)
	return &results, err
}

func (dc DalCard) FetchFromFactionPaging(factionID bson.ObjectId, limit int, offset int) (*[]models.Card, error) {
	results := []models.Card{}
	err := dc.collection.Find(bson.M{"faction_id": factionID}).Limit(limit).Sort("name.en-US").Skip(offset).All(&results)
	return &results, err
}

func (dc DalCard) CountLeader(groupID bson.ObjectId) (int, error) {
	result, err := dc.collection.Find(bson.M{"group_id": groupID}).Count()
	return result, err
}

func (dc DalCard) CountFromFaction(factionID bson.ObjectId) (int, error) {
	result, err := dc.collection.Find(bson.M{"faction_id": factionID}).Count()
	return result, err
}

func (dc DalCard) FetchFromArray(cardIDs []bson.ObjectId) (*[]models.Card, error) {
	results := []models.Card{}
	err := dc.collection.Find(bson.M{"_id": bson.M{"$in": cardIDs}}).Sort("name.en-US").All(&results)

	return &results, err
}

func (dc DalCard) Count() (int, error) {
	result, err := dc.collection.Count()
	return result, err
}
