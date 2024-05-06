package services

import (
	"auth-service/domain"
	"auth-service/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthServiceImpl struct {
	collection  *mongo.Collection
	ctx         context.Context
	userService UserService
}

func NewAuthService(collection *mongo.Collection, ctx context.Context, userService UserService) AuthService {
	return &AuthServiceImpl{collection, ctx, userService}
}

func (uc *AuthServiceImpl) Login() (*domain.User, error) {
	return nil, nil
}

func (uc *AuthServiceImpl) Registration(user *domain.User, ctx context.Context) (*domain.UserResponse, error) {
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword

	credentials := &domain.User{
		ID:       primitive.NewObjectID(),
		Username: user.Username,
		Password: hashedPassword,
		Email:    user.Email,
		Name:     user.Name,
		Lastname: user.Lastname,
		UserRole: user.UserRole,
	}
	res, err := uc.collection.InsertOne(ctx, credentials)
	if err != nil {
		return nil, err
	}

	var newUser *domain.UserResponse
	query := bson.M{"_id": res.InsertedID}

	err = uc.collection.FindOne(ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil

}
