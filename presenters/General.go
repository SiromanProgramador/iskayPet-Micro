package presenters

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//method to parse strucs to bson
func StructToBson(structObject interface{}) (bson.M, error) {
	var bsonToReturn bson.M
	data, err := bson.Marshal(structObject)
	if err != nil {
		return bsonToReturn, err
	}

	errUnmarshaling := bson.Unmarshal(data, &bsonToReturn)
	if err != nil {
		return bsonToReturn, errUnmarshaling
	}
	return bsonToReturn, nil
}

//method to parse array strucs to bson array
func ArrayStructToBson(array, outArray interface{}) error {
	inStructArrData, err := bson.Marshal(array)
	if err != nil {
		return err
	}
	// kind 4 for array
	raw := bson.Raw{Kind: 4, Data: inStructArrData}

	return raw.Unmarshal(outArray)
}

//method to get timenow
func GetTimeNow() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

//class and methos to instance new objects into collections in MongoDB
const (
	InstanceStatusActive   string = "ACTIVE"
	InstanceStatusInactive string = "INACTIVE"
	InstanceStatusDeleted  string = "DELETED"
)

type Instance struct {
	Status     string `json:"status,omitempty" bson:"status,omitempty"`
	CreatedAt  int64  `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	ModifiedAt int64  `json:"modifiedAt,omitempty" bson:"modifiedAt,omitempty"`
}

func CreateInstance() Instance {
	var instance Instance
	instance.Status = InstanceStatusActive
	instance.ModifiedAt = GetTimeNow()
	instance.CreatedAt = GetTimeNow()

	return instance
}
