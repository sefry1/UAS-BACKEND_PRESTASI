package repository

import (
	"context"
	"prestasi_backend/app/database"
	"prestasi_backend/app/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AchievementMongoRepository struct {
	Coll *mongo.Collection
}

func NewAchievementMongoRepository() *AchievementMongoRepository {
	return &AchievementMongoRepository{
		Coll: database.MongoDB.Collection("achievements"),
	}
}

func (r *AchievementMongoRepository) FindByID(id string) (*model.AchievementMongo, error) {
	objID, _ := primitive.ObjectIDFromHex(id)

	var data model.AchievementMongo
	err := r.Coll.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *AchievementMongoRepository) Create(data model.AchievementMongo) (string, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	res, err := r.Coll.InsertOne(context.Background(), data)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *AchievementMongoRepository) Update(id string, updateData model.AchievementMongo) error {
	objID, _ := primitive.ObjectIDFromHex(id)

	update := bson.M{
		"$set": bson.M{
			"title":           updateData.Title,
			"description":     updateData.Description,
			"achievementType": updateData.AchievementType,
			"details":         updateData.Details,
			"tags":            updateData.Tags,
			"points":          updateData.Points,
			"updatedAt":       time.Now(),
		},
	}

	_, err := r.Coll.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	return err
}

func (r *AchievementMongoRepository) AddAttachment(id string, file model.Attachment) error {
	objID, _ := primitive.ObjectIDFromHex(id)

	update := bson.M{
		"$push": bson.M{"attachments": file},
		"$set":  bson.M{"updatedAt": time.Now()},
	}

	_, err := r.Coll.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	return err
}

func (r *AchievementMongoRepository) Delete(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := r.Coll.DeleteOne(context.Background(), bson.M{"_id": objID})
	return err
}
