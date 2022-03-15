package mongodb

import (
	"iskayPetMicro/Domains/pets/entity/repository"
	"iskayPetMicro/model"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type repo struct {
	session *mgo.Session
}

func NewMongoDBRepository(session *mgo.Session) repository.RepositoryInterface {
	return &repo{
		session: session,
	}
}

func setCollection(session *mgo.Session, colName string) *mgo.Collection {
	collection := session.DB("challengeDB").C(colName)
	return collection
}

const CollectionName = model.DBCOLLECTION_PETS

func (r *repo) Create(objectToCreate *model.Pet) error {
	collection := setCollection(r.session, CollectionName)
	errInsert := collection.Insert(objectToCreate)

	return errInsert
}

func (r *repo) GetAll(filter model.QueryFilters) (objectToReturn []model.Pet, errFind error) {
	collection := setCollection(r.session, CollectionName)
	var aggregate []bson.M

	//match
	matchFile := bson.M{"$match": filter.Filter}
	aggregate = append(aggregate, matchFile)

	//Lookup speciesId
	lookupSpeciesId := bson.M{"$lookup": bson.M{"from": "species", "localField": "speciesId", "foreignField": "_id", "as": "speciesInfo"}}
	aggregate = append(aggregate, lookupSpeciesId)

	//unwind speciesId
	unwindSpeciesId := bson.M{"$unwind": bson.M{"path": "$speciesInfo", "preserveNullAndEmptyArrays": true}}
	aggregate = append(aggregate, unwindSpeciesId)

	//skip
	if filter.Skip > 0 {
		skipPipeLine := bson.M{"$skip": filter.Skip}
		aggregate = append(aggregate, skipPipeLine)
	}

	//limit
	if filter.Limit > 0 {
		limitPipeLine := bson.M{"$limit": filter.Limit}
		aggregate = append(aggregate, limitPipeLine)
	}

	//Find
	pipe := collection.Pipe(aggregate)
	errFind = pipe.All(&objectToReturn)

	return objectToReturn, errFind
}

func (r *repo) Count(filter model.QueryFilters) (int, error) {
	var objectToReturn []model.Pet
	collection := setCollection(r.session, CollectionName)
	var aggregate []bson.M

	//match
	matchFile := bson.M{"$match": filter.Filter}
	aggregate = append(aggregate, matchFile)

	//skip
	if filter.Skip > 0 {
		skipPipeLine := bson.M{"$skip": filter.Skip}
		aggregate = append(aggregate, skipPipeLine)
	}

	//limit
	if filter.Limit > 0 {
		limitPipeLine := bson.M{"$limit": filter.Limit}
		aggregate = append(aggregate, limitPipeLine)
	}

	//Find
	pipe := collection.Pipe(aggregate)
	errFind := pipe.All(&objectToReturn)

	return len(objectToReturn), errFind
}
