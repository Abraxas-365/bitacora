package repository

import (
	"context"
	"github/Abraxas-365/bitacora/internal/myerror"
	"github/Abraxas-365/bitacora/pkg/report/domain/models"

	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) Create(new models.Report) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	_, err := collection.InsertOne(ctx, new)
	if err != nil {
		return myerror.New("Failed to create report", http.StatusInternalServerError)
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
		return myerror.New("Failed to delete report", http.StatusInternalServerError)
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
		return models.Reports{}, myerror.New("Failed to get reports", http.StatusInternalServerError)
	}

	if err := cursor.All(ctx, &reports); err != nil {
		return models.Reports{}, myerror.New("Failed to fetch reports", http.StatusInternalServerError)
	}

	return reports, nil
}

func (r *repo) Update(report models.Report) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	filter := bson.M{"_id": report.Id}
	update := bson.M{"$set": report}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return myerror.New("Failed to update report", http.StatusInternalServerError)
	}

	return nil
}
