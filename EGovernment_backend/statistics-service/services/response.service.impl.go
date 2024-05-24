package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"statistics-service/domain"
)

type ResponseServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewResponseServiceImpl(collection *mongo.Collection, ctx context.Context) ResponseService {
	return &ResponseServiceImpl{collection, ctx}
}

func (r *ResponseServiceImpl) Create(response *domain.Response) (error, bool) {
	_, err := r.collection.InsertOne(r.ctx, response)
	if err != nil {
		return err, false
	}
	return nil, true
}

func (r *ResponseServiceImpl) GetAll() ([]*domain.Response, error) {
	var responses []*domain.Response
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
		var response domain.Response
		if err := cursor.Decode(&response); err != nil {
			return nil, err
		}
		responses = append(responses, &response)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return responses, nil
}

func (r *ResponseServiceImpl) GetById(id string) (*domain.Response, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var response domain.Response
	err = r.collection.FindOne(r.ctx, bson.M{"_id": objID}).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
