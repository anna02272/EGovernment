package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"statistics-service/domain"
)

type RequestServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewRequestServiceImpl(collection *mongo.Collection, ctx context.Context) RequestService {
	return &RequestServiceImpl{collection, ctx}
}

func (r *RequestServiceImpl) Create(request *domain.Request) (error, bool) {
	_, err := r.collection.InsertOne(r.ctx, request)
	if err != nil {
		return err, false
	}
	return nil, true
}

func (r *RequestServiceImpl) GetAll() ([]*domain.Request, error) {
	var requests []*domain.Request
	filter := bson.D{}

	cursor, err := r.collection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, r.ctx)

	for cursor.Next(r.ctx) {
		var request domain.Request
		if err := cursor.Decode(&request); err != nil {
			return nil, err
		}
		requests = append(requests, &request)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return requests, nil
}

func (r *RequestServiceImpl) GetById(id string) (*domain.Request, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var request domain.Request
	err = r.collection.FindOne(r.ctx, bson.M{"_id": objID}).Decode(&request)
	if err != nil {
		return nil, err
	}
	return &request, nil
}

func (r *RequestServiceImpl) Delete(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(r.ctx, bson.M{"_id": objID})
	return err
}
