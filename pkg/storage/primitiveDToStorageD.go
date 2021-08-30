package storage

import "go.mongodb.org/mongo-driver/bson/primitive"

func PrimitiveDToStorageD(primitiveD primitive.D) (d D) {
	d = D{}

	for _, value := range primitiveD {
		d = append(d, E{
			Key:   value.Key,
			Value: value.Value,
		})
	}

	return
}
