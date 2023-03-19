package repository

import (
	"context"
	"github/Abraxas-365/bitacora/pkg/report/domain/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) Create(new models.Report) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	_, err := collection.InsertOne(ctx, new)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Delete(reportId interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	filter := bson.M{"_id": reportId}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Get(key string, value interface{}) (models.Reports, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	var reports models.Reports
	filter := bson.D{{key, value}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return models.Reports{}, err
	}

	if err := cursor.All(ctx, &reports); err != nil {
		return models.Reports{}, err
	}

	return reports, nil
}

func (r *repo) Update(report models.Report) error {
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
