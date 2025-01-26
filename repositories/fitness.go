package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"server/db"
	"server/dto"
)

func CreateFitness(ctx context.Context, fitnessDTO dto.FitnessDTO) error {
	collection := db.GetCollection("fitness")

	fitness := bson.M{
		"title":         fitnessDTO.Title,
		"image":         fitnessDTO.Image,
		"description":   fitnessDTO.Description,
		"category":      fitnessDTO.Category,
		"body_category": fitnessDTO.BodyCategory,
		"trait":         fitnessDTO.Trait,
		"video":         fitnessDTO.Video,
		"workout":       fitnessDTO.Workout,
		"deleted":       fitnessDTO.Deleted,
	}

	_, err := collection.InsertOne(ctx, fitness)
	return err
}

func GetFitnessByID(ctx context.Context, id string) (dto.FitnessDTO, error) {
	collection := db.GetCollection("fitness")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dto.FitnessDTO{}, err
	}

	var fitness dto.FitnessDTO
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&fitness)
	if err != nil {
		return dto.FitnessDTO{}, err
	}

	return fitness, nil
}

func GetAllFitness(ctx context.Context) ([]dto.FitnessDTO, error) {
	collection := db.GetCollection("fitness")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var fitnessList []dto.FitnessDTO
	for cursor.Next(ctx) {
		var fitness dto.FitnessDTO
		if err := cursor.Decode(&fitness); err != nil {
			return nil, err
		}
		fitnessList = append(fitnessList, fitness)
	}

	return fitnessList, nil
}

func UpdateFitness(ctx context.Context, id string, fitnessDTO dto.FitnessDTO) error {
	collection := db.GetCollection("fitness")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"title":         fitnessDTO.Title,
			"image":         fitnessDTO.Image,
			"description":   fitnessDTO.Description,
			"category":      fitnessDTO.Category,
			"body_category": fitnessDTO.BodyCategory,
			"trait":         fitnessDTO.Trait,
			"video":         fitnessDTO.Video,
			"workout":       fitnessDTO.Workout,
			"deleted":       fitnessDTO.Deleted,
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}

func DeleteFitness(ctx context.Context, id string) error {
	collection := db.GetCollection("fitness")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
