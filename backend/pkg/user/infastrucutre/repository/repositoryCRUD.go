package repository

import (
	"context"
	"github/Abraxas-365/bitacora/internal/myerror"
	"github/Abraxas-365/bitacora/pkg/user/domain/models"

	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repo) Create(new models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	_, err := collection.InsertOne(ctx, new)
	if err != nil {
		return myerror.New("Failed to create user", http.StatusInternalServerError)
	}

	return nil
}

func (r *repo) Delete(userId interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	filter := bson.M{"_id": userId}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return myerror.New("Failed to delete user", http.StatusInternalServerError)
	}

	return nil
}

func (r *repo) Get(key string, value interface{}) (models.Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	var users models.Users
	filter := bson.D{{key, value}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return models.Users{}, myerror.New("Failed to get users", http.StatusInternalServerError)
	}

	if err := cursor.All(ctx, &users); err != nil {
		return models.Users{}, myerror.New("Failed to fetch users", http.StatusInternalServerError)
	}

	return users, nil
}

func (r *repo) Update(report models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	filter := bson.M{"_id": report.Id}
	_, err := collection.UpdateOne(ctx, filter, report)
	if err != nil {
		return myerror.New("Failed to update user", http.StatusInternalServerError)
	}

	return nil
}

func (r *repo) GetUser(key string, value interface{}) (models.User, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	user := new(models.User)
	if err := collection.FindOne(ctx, bson.M{key: value}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return models.User{}, false, myerror.New("User not found", http.StatusNotFound)
		}
		return models.User{}, false, myerror.New("Failed to get user", http.StatusInternalServerError)
	}
	return *user, true, nil
}
