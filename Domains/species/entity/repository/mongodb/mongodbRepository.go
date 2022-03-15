package mongodb

import (
	"iskayPetMicro/Domains/species/entity/repository"
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

const CollectionName = model.DBCOLLECTION_SPECIES

func (r *repo) Create(objectToCreate *model.Species) error {
	collection := setCollection(r.session, CollectionName)
	errInsert := collection.Insert(objectToCreate)

	return errInsert
}

func (r *repo) Delete(queryFilter *model.QueryFilters) error {

	collection := setCollection(r.session, CollectionName)
	errDelete := collection.Remove(queryFilter.Filter)

	return errDelete
}

func (r *repo) GetOne(queryFilter *model.QueryFilters) (objectToReturn model.Species, errFind error) {
	collection := setCollection(r.session, CollectionName)

	var aggregate []bson.M

	//match
	matchFile := bson.M{"$match": queryFilter.Filter}
	aggregate = append(aggregate, matchFile)

	//skip
	if queryFilter.Skip > 0 {
		skipPipeLine := bson.M{"$skip": queryFilter.Skip}
		aggregate = append(aggregate, skipPipeLine)
	}

	//limit
	if queryFilter.Limit > 0 {
		limitPipeLine := bson.M{"$limit": queryFilter.Limit}
		aggregate = append(aggregate, limitPipeLine)
	}

	//Find
	pipe := collection.Pipe(aggregate)
	errFind = pipe.One(&objectToReturn)

	return objectToReturn, errFind

}

func (r *repo) GetAll(filter model.QueryFilters) (objectToReturn []model.Species, errFind error) {
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
	errFind = pipe.All(&objectToReturn)

	return objectToReturn, errFind
}

func (r *repo) Update(filter *model.QueryFilters, objectToUpdate model.Species) error {
	collection := setCollection(r.session, CollectionName)

	errUpdate := collection.Update(filter.Filter, bson.M{"$set": objectToUpdate})

	return errUpdate
}
