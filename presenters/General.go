package presenters

import (
	"time"
)

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
