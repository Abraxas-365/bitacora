package repository

import (
	"context"
	"errors"
	"github/Abraxas-365/bitacora/internal/leakerrs"
	"github/Abraxas-365/bitacora/pkg/user/domain/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repo) Create(new models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	_, err := collection.InsertOne(ctx, new)
	if err != nil {
		return err
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
		return err
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
		return models.Users{}, err
	}

	if err := cursor.All(ctx, &users); err != nil {
		return models.Users{}, err
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
		return err
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
			return models.User{}, false, nil
		}
		return models.User{}, false, errors.New(leakerrs.InternalError)
	}
	return *user, true, nil
}
