package services

import (
	"context"
	"court-service/domain"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
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
func (ss *SubjectServiceImpl) GetSubjectByID(id primitive.ObjectID) (*domain.Subject, error) {
	var subject domain.Subject
	err := ss.collection.FindOne(ss.ctx, bson.M{"_id": id}).Decode(&subject)
	if err != nil {
		return nil, err
	}
	return &subject, nil
}

func (ss *SubjectServiceImpl) GetAllSubjects() ([]domain.Subject, error) {
	cursor, err := ss.collection.Find(ss.ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ss.ctx)

	var subjects []domain.Subject
	if err = cursor.All(ss.ctx, &subjects); err != nil {
		return nil, err
	}

	return subjects, nil
}
func (ss *SubjectServiceImpl) UpdateSubjectStatus(id primitive.ObjectID, status domain.Status) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": status}}

	_, err := ss.collection.UpdateOne(ss.ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
func (ss *SubjectServiceImpl) UpdateSubjectJudgment(id primitive.ObjectID, judgment string) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"judgment": judgment}}

	_, err := ss.collection.UpdateOne(ss.ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (ss *SubjectServiceImpl) UpdateSubjectCompromis(id primitive.ObjectID, compromis string) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"compromis": compromis}}

	_, err := ss.collection.UpdateOne(ss.ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
