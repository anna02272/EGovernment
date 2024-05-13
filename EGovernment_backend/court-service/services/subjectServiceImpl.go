package services

import (
	"context"
	"court-service/domain"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SubjectServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewSubjectServiceImpl(collection *mongo.Collection, ctx context.Context) SubjectService {
	return &SubjectServiceImpl{collection, ctx}
}

func (ss *SubjectServiceImpl) CreateSubject(subject *domain.Subject) (*domain.Subject, error) {
	result, err := ss.collection.InsertOne(ss.ctx, subject)
	if err != nil {
		return nil, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to get inserted ID")
	}

	subject.ID = insertedID
	return subject, nil
}
