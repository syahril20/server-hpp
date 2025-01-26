package repositories

import (
	"context"
	"server/db"
	dtoAssesment "server/dto/assesment"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Assesment
func CreateAssesment(ctx context.Context, assessment dtoAssesment.CreateAssesmentRequest) (dtoAssesment.CreateAssesmentRequest, error) {
	collection := db.GetCollection("assessment")
	_, err := collection.InsertOne(ctx, assessment)
	return assessment, err
}

func GetAssesmentByName(ctx context.Context, name string) (*dtoAssesment.CreateAssesmentRequest, error) {
	collection := db.GetCollection("assessment")
	var product dtoAssesment.CreateAssesmentRequest
	err := collection.FindOne(ctx, bson.M{"name": name}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}

func GetAssesmentById(ctx context.Context, id primitive.ObjectID) (*dtoAssesment.CreateAssesmentRequest, error) {
	collection := db.GetCollection("assesment")
	var assesment dtoAssesment.CreateAssesmentRequest
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&assesment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &assesment, nil
}

func GetActiveAssesmentById(ctx context.Context, id primitive.ObjectID) (*dtoAssesment.CreateAssesmentRequest, error) {
	collection := db.GetCollection("assesment")
	var assesment dtoAssesment.CreateAssesmentRequest

	filter := bson.M{"_id": id, "deleted_at": nil}
	err := collection.FindOne(ctx, filter).Decode(&assesment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &assesment, nil
}

func GetAllAssesments(ctx context.Context) ([]dtoAssesment.CreateAssesmentRequest, error) {
	collection := db.GetCollection("assesment")
	var assesments []dtoAssesment.CreateAssesmentRequest

	filter := bson.M{"deleted_at": bson.M{"$eq": nil}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var assesment dtoAssesment.CreateAssesmentRequest
		if err := cursor.Decode(&assesment); err != nil {
			return nil, err
		}
		assesments = append(assesments, assesment)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return assesments, nil
}

func GetAllNonActiveAssesments(ctx context.Context) ([]dtoAssesment.CreateAssesmentRequest, error) {
	collection := db.GetCollection("assesment")
	var assesments []dtoAssesment.CreateAssesmentRequest

	filter := bson.M{"deleted_at": bson.M{"$ne": nil}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var assesment dtoAssesment.CreateAssesmentRequest
		if err := cursor.Decode(&assesment); err != nil {
			return nil, err
		}
		assesments = append(assesments, assesment)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return assesments, nil
}

func UpdateAssesmentById(ctx context.Context, id primitive.ObjectID, updateData dtoAssesment.UpdateAssesmentRequest) (*dtoAssesment.UpdateAssesmentRequest, error) {
	collection := db.GetCollection("assesment")
	var updated dtoAssesment.UpdateAssesmentRequest

	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": updateData,
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(ctx, filter).Decode(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func ActiveAssesmentById(ctx context.Context, id primitive.ObjectID, updateData dtoAssesment.ActiveAssesmentRequest) (*dtoAssesment.ActiveAssesmentRequest, error) {
	collection := db.GetCollection("assesment")
	var updated dtoAssesment.ActiveAssesmentRequest

	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": updateData,
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(ctx, filter).Decode(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

// Questionnaire

func CreateQuestionnaire(ctx context.Context, assesmentId primitive.ObjectID, questionnaire dtoAssesment.CreateQuestionnaireRequest) (dtoAssesment.CreateQuestionnaireRequest, error) {
	collection := db.GetCollection("assesment")
	filter := bson.M{"_id": assesmentId}
	update := bson.M{
		"$push": bson.M{"questionnaire": questionnaire},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	return questionnaire, err
}

func GetQuestionnaireByQuestion(ctx context.Context, question string) (*dtoAssesment.CreateQuestionnaireRequest, error) {
	collection := db.GetCollection("assesment")
	var questionnaire dtoAssesment.CreateQuestionnaireRequest

	filter := bson.M{
		"questionnaire.question": question,
	}

	err := collection.FindOne(ctx, filter).Decode(&questionnaire)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &questionnaire, nil
}

func GetAllDataQuestionnaires(ctx context.Context, assesmentId primitive.ObjectID) ([]dtoAssesment.CreateQuestionnaireRequest, error) {
	collection := db.GetCollection("assesment")
	var result []dtoAssesment.CreateQuestionnaireRequest

	filter := bson.M{
		"_id":        assesmentId,
		"deleted_at": bson.M{"$eq": nil},
	}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var assesment struct {
			Questionnaire []dtoAssesment.CreateQuestionnaireRequest `bson:"questionnaire"`
		}

		if err := cursor.Decode(&assesment); err != nil {
			return nil, err
		}
		for _, questionnaire := range assesment.Questionnaire {
			if questionnaire.DeletedAt == nil {
				result = append(result, questionnaire)
			}
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func GetAllNonActiveQuestionnaires(ctx context.Context, assesmentId primitive.ObjectID) ([]dtoAssesment.CreateQuestionnaireRequest, error) {
	collection := db.GetCollection("assesment")
	var result []dtoAssesment.CreateQuestionnaireRequest

	filter := bson.M{
		"_id":        assesmentId,
		"deleted_at": bson.M{"$eq": nil},
	}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var assesment struct {
			Questionnaire []dtoAssesment.CreateQuestionnaireRequest `bson:"questionnaire"`
		}

		if err := cursor.Decode(&assesment); err != nil {
			return nil, err
		}
		for _, questionnaire := range assesment.Questionnaire {
			if questionnaire.DeletedAt != nil {
				result = append(result, questionnaire)
			}
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func GetQuestionnaireById(ctx context.Context, id primitive.ObjectID) (*dtoAssesment.CreateQuestionnaireRequest, error) {
	collection := db.GetCollection("assesment")
	var questionnaire dtoAssesment.CreateQuestionnaireRequest

	filter := bson.M{"questionnaire._id": id}

	err := collection.FindOne(ctx, filter).Decode(&questionnaire)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &questionnaire, nil
}

func UpdateQuestionnaireById(ctx context.Context, id primitive.ObjectID, updateData dtoAssesment.UpdateQuestionnaireRequest) (*dtoAssesment.UpdateQuestionnaireRequest, error) {
	collection := db.GetCollection("assesment")
	var updated dtoAssesment.UpdateQuestionnaireRequest

	filter := bson.M{"questionnaire._id": id}

	update := bson.M{
		"$set": bson.M{
			"questionnaire.$.question":   updateData.Question,
			"questionnaire.$.updated_by": updateData.UpdatedBy,
			"questionnaire.$.updated_at": updateData.UpdatedAt,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(ctx, filter).Decode(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func ActiveQuestionnaireById(ctx context.Context, id primitive.ObjectID, updateData dtoAssesment.ActiveQuestionnaireRequest) (*dtoAssesment.ActiveQuestionnaireRequest, error) {
	collection := db.GetCollection("assesment")
	var updated dtoAssesment.ActiveQuestionnaireRequest

	filter := bson.M{"questionnaire._id": id}

	update := bson.M{
		"$set": bson.M{
			"questionnaire.$.updated_by": updateData.UpdatedBy,
			"questionnaire.$.updated_at": updateData.UpdatedAt,
			"questionnaire.$.deleted_at": updateData.DeletedAt,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(ctx, filter).Decode(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

// Answer
func CreateAnswer(ctx context.Context, questionnaireId primitive.ObjectID, answer dtoAssesment.CreateAnswerRequest) (dtoAssesment.CreateAnswerRequest, error) {
	collection := db.GetCollection("assesment")
	filter := bson.M{"questionnaire._id": questionnaireId}
	update := bson.M{
		"$push": bson.M{"questionnaire.$.answer": answer},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	return answer, err
}

func GetAllActiveAnswers(ctx context.Context, questionnaireId primitive.ObjectID) ([]dtoAssesment.CreateAnswerRequest, error) {
	collection := db.GetCollection("assesment")
	var result []dtoAssesment.CreateAnswerRequest

	filter := bson.M{"questionnaire._id": questionnaireId}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var assesment struct {
			Questionnaire []struct {
				Id      primitive.ObjectID                 `bson:"_id"`
				Answers []dtoAssesment.CreateAnswerRequest `bson:"answer"`
			} `bson:"questionnaire"`
		}

		if err := cursor.Decode(&assesment); err != nil {
			return nil, err
		}
		for _, questionnaire := range assesment.Questionnaire {
			if questionnaire.Id == questionnaireId {
				result = append(result, questionnaire.Answers...)
			}
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func GetAnswerByValue(ctx context.Context, questionnaireId primitive.ObjectID, value bool) (*dtoAssesment.CreateAnswerRequest, error) {
	collection := db.GetCollection("assesment")
	var answer dtoAssesment.CreateAnswerRequest

	filter := bson.M{"questionnaire.answer._id_questionnaire": questionnaireId, "questionnaire.answer.value": value}

	err := collection.FindOne(ctx, filter).Decode(&answer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &answer, nil
}

func GetAllAnswers(ctx context.Context, questionnaireId primitive.ObjectID) ([]dtoAssesment.CreateAnswerRequest, error) {
	collection := db.GetCollection("assesment")
	var result []dtoAssesment.CreateAnswerRequest

	filter := bson.M{"questionnaire._id": questionnaireId}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var questionnaire struct {
			Answers []dtoAssesment.CreateAnswerRequest `bson:"answers"`
		}

		if err := cursor.Decode(&questionnaire); err != nil {
			return nil, err
		}
		for _, answer := range questionnaire.Answers {
			if answer.DeletedAt == nil {
				result = append(result, answer)
			}
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func GetAnswerById(ctx context.Context, id primitive.ObjectID) (*dtoAssesment.CreateAnswerRequest, error) {
	collection := db.GetCollection("assesment")
	var result struct {
		Questionnaire []struct {
			Answers []dtoAssesment.CreateAnswerRequest `bson:"answer"`
		} `bson:"questionnaire"`
	}

	filter := bson.M{"questionnaire.answer._id": id}

	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	for _, questionnaire := range result.Questionnaire {
		for _, answer := range questionnaire.Answers {
			if answer.Id == id {
				return &answer, nil
			}
		}
	}

	return nil, nil
}

func UpdateAnswerById(ctx context.Context, id primitive.ObjectID, updateData dtoAssesment.UpdateAnswerRequest) (*dtoAssesment.UpdateAnswerRequest, error) {
	collection := db.GetCollection("assesment")
	var updated dtoAssesment.UpdateAnswerRequest

	filter := bson.M{"questionnaire.answer._id": id}

	update := bson.M{
		"$set": bson.M{
			"questionnaire.$.answer.$[elem].description": updateData.Description,
			"questionnaire.$.answer.$[elem].updated_at":  updateData.UpdatedAt,
			"questionnaire.$.answer.$[elem].updated_by":  updateData.UpdatedBy,
		},
	}
	arrayFilters := options.ArrayFilters{
		Filters: []interface{}{bson.M{"elem._id": id}},
	}
	opts := options.UpdateOptions{
		ArrayFilters: &arrayFilters,
	}

	_, err := collection.UpdateOne(ctx, filter, update, &opts)
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(ctx, filter).Decode(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}
