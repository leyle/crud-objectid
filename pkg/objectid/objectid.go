package objectid

import "github.com/leyle/crud-objectid/internal/primitive"

func GetObjectId() string {
	return primitive.NewObjectID().Hex()
}
