package services

import (
	"context"
	"court-service/domain"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CitizenServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewCitizenServiceImpl(collection *mongo.Collection, ctx context.Context) CitizenService {
	return &CitizenServiceImpl{collection, ctx}
}

func (s *CitizenServiceImpl) InsertCitizen(citizen *domain.Citizen) (*domain.Citizen, string, error) {
	result, err := s.collection.InsertOne(context.Background(), citizen)
	if err != nil {
		return nil, "", err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, "", errors.New("failed to get inserted ID")
	}

	return citizen, insertedID.Hex(), nil
}

func (s *CitizenServiceImpl) GetAllCitizens() ([]*domain.Citizen, error) {
	var citizens []*domain.Citizen
	filter := bson.D{}

	cursor, err := s.collection.Find(s.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(s.ctx)

	for cursor.Next(s.ctx) {
		var citizen domain.Citizen
		if err := cursor.Decode(&citizen); err != nil {
			return nil, err
		}
		citizens = append(citizens, &citizen)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return citizens, nil
}

func (s *CitizenServiceImpl) GetCitizenByID(jmbg string) (*domain.Citizen, error) {
	var citizen domain.Citizen
	filter := bson.M{"jmbg": jmbg}

	err := s.collection.FindOne(s.ctx, filter).Decode(&citizen)
	if err != nil {
		return nil, err
	}

	return &citizen, nil
}
