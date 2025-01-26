package handlers

import (
	"context"
	"fmt"
	"net/http"
	dtoAssesment "server/dto/assesment"
	dto "server/dto/result"
	"server/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Assesment

func GetAllActiveAssesmentsHandler(c *gin.Context) {
	assesments, err := repositories.GetAllAssesments(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    assesments})
}

func GetAllNonActiveAssesmentsHandler(c *gin.Context) {
	assesments, err := repositories.GetAllNonActiveAssesments(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    assesments})
}

func CreateAssesmentHandler(c *gin.Context) {
	var req dtoAssesment.CreateAssesmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	if req.Name == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	_, err := primitive.ObjectIDFromHex(req.Id.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Product ID"})
		return
	}

	assesmentName, _ := repositories.GetAssesmentByName(context.Background(), req.Name)
	if assesmentName != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	assesment := dtoAssesment.CreateAssesmentRequest{
		Id:            primitive.NewObjectID(),
		Name:          req.Name,
		DeletedAt:     nil,
		Questionnaire: []dtoAssesment.CreateQuestionnaireRequest{},
		CreatedAt:     currentTime,
		CreatedBy:     "System",
		UpdatedAt:     currentTime,
		UpdatedBy:     "System",
	}

	data, err := repositories.CreateAssesment(context.Background(), assesment)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.SuccessResult{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    data})
}

func UpdateAssesmentHandler(c *gin.Context) {
	id := c.Param("id")
	var req dtoAssesment.CreateAssesmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	if req.Name == "" || id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	assesmentId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Product ID"})
		return
	}

	existingAssesment, _ := repositories.GetAssesmentById(context.Background(), assesmentId)
	if existingAssesment == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Not Exist"})
		return
	}

	assesmentName, _ := repositories.GetAssesmentByName(context.Background(), req.Name)
	if assesmentName != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	assesment := dtoAssesment.UpdateAssesmentRequest{
		Name:      req.Name,
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	_, err = repositories.UpdateAssesmentById(context.Background(), assesmentId, assesment)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.SuccessResult{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    assesment})
}

func DeleteAssesmentHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	questionnaireId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Assesment ID"})
		return
	}

	existingAssesment, _ := repositories.GetAssesmentById(context.Background(), questionnaireId)
	if existingAssesment == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Not Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	assesment := dtoAssesment.ActiveAssesmentRequest{
		DeletedAt: &currentTime,
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	_, err = repositories.ActiveAssesmentById(context.Background(), questionnaireId, assesment)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    assesment})
}

func ActiveAssesmentHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	assesmentId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Assesment ID"})
		return
	}

	existingAssesment, _ := repositories.GetAssesmentById(context.Background(), assesmentId)
	if existingAssesment == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Not Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	activeAssesment := dtoAssesment.ActiveAssesmentRequest{
		DeletedAt: nil,
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	_, err = repositories.ActiveAssesmentById(context.Background(), assesmentId, activeAssesment)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    activeAssesment})
}

// Questionnaire

func GetAllActiveQuestionnaireHandler(c *gin.Context) {
	id := c.Param("id")
	assesmentId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Assesment ID"})
		return
	}
	assesments, err := repositories.GetAllDataQuestionnaires(context.Background(), assesmentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    assesments})
}

func GetAllNonActiveQuestionnaireHandler(c *gin.Context) {
	id := c.Param("id")
	assesmentId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Assesment ID"})
		return
	}
	assesments, err := repositories.GetAllDataQuestionnaires(context.Background(), assesmentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    assesments})
}

func CreateQuestionnaireHandler(c *gin.Context) {
	var req dtoAssesment.CreateQuestionnaireRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	if req.Question == "" || req.IdAsessment.Hex() == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	assesmentId, err := primitive.ObjectIDFromHex(req.IdAsessment.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Product ID"})
		return
	}

	nameQuestion, _ := repositories.GetQuestionnaireByQuestion(context.Background(), req.Question)
	if nameQuestion != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Exist"})
		return
	}

	existingQuestionnaire, _ := repositories.GetAssesmentById(context.Background(), assesmentId)
	if existingQuestionnaire == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Assesment Not Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	newQuestionnaire := dtoAssesment.CreateQuestionnaireRequest{
		Id:          primitive.NewObjectID(),
		IdAsessment: req.IdAsessment,
		Question:    req.Question,
		Answer:      []dtoAssesment.CreateAnswerRequest{},
		CreatedAt:   currentTime,
		CreatedBy:   "System",
		UpdatedAt:   currentTime,
		UpdatedBy:   "System",
	}

	data, err := repositories.CreateQuestionnaire(context.Background(), assesmentId, newQuestionnaire)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.SuccessResult{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    data})
}

func UpdateQuestionnaireHandler(c *gin.Context) {
	id := c.Param("id")
	var req dtoAssesment.UpdateQuestionnaireRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	if req.Question == "" || id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	questionnaireId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Questionnaire ID"})
		return
	}

	existingQuestionnaire, _ := repositories.GetQuestionnaireById(context.Background(), questionnaireId)
	if existingQuestionnaire == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Not Exist"})
		return
	}

	nameQuestion, _ := repositories.GetQuestionnaireByQuestion(context.Background(), req.Question)
	if nameQuestion != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	updatedQuestionnaire := dtoAssesment.UpdateQuestionnaireRequest{
		Question:  req.Question,
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	_, err = repositories.UpdateQuestionnaireById(context.Background(), questionnaireId, updatedQuestionnaire)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    updatedQuestionnaire})
}

func DeleteQuestionnaireHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	questionnaireId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Questionnaire ID"})
		return
	}

	existingQuestionnaire, _ := repositories.GetQuestionnaireById(context.Background(), questionnaireId)
	if existingQuestionnaire == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Not Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	softDeletedQuestionnaire := dtoAssesment.ActiveQuestionnaireRequest{
		DeletedAt: &currentTime,
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	_, err = repositories.ActiveQuestionnaireById(context.Background(), questionnaireId, softDeletedQuestionnaire)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    softDeletedQuestionnaire})
}

func ActiveQuestionnaireHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	questionnaireId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Questionnaire ID"})
		return
	}

	existingQuestionnaire, _ := repositories.GetQuestionnaireById(context.Background(), questionnaireId)
	if existingQuestionnaire == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Not Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	ActiveQuestionnaire := dtoAssesment.ActiveQuestionnaireRequest{
		DeletedAt: nil,
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	_, err = repositories.ActiveQuestionnaireById(context.Background(), questionnaireId, ActiveQuestionnaire)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    ActiveQuestionnaire})
}

// Answer
func GetAllActiveAnswersHandler(c *gin.Context) {
	id := c.Param("id")
	questionnaireId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Questionnaire ID"})
		return
	}
	answers, err := repositories.GetAllActiveAnswers(context.Background(), questionnaireId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    answers})
}

func CreateAnswerHandler(c *gin.Context) {
	var req dtoAssesment.CreateAnswerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	if req.Description == "" || req.IdQuestionnaire.Hex() == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	questionnaireId, err := primitive.ObjectIDFromHex(req.IdQuestionnaire.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Questionnaire ID"})
		return
	}

	existingQuestionnaire, _ := repositories.GetQuestionnaireById(context.Background(), questionnaireId)
	if existingQuestionnaire == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Questionnaire Not Exist"})
		return
	}

	valueAnswer, _ := repositories.GetAnswerByValue(context.Background(), questionnaireId, req.Value)
	fmt.Println(valueAnswer, "valueAnswer")
	if valueAnswer != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	newAnswer := dtoAssesment.CreateAnswerRequest{
		Id:              primitive.NewObjectID(),
		IdQuestionnaire: req.IdQuestionnaire,
		Value:           req.Value,
		Description:     req.Description,
		CreatedAt:       currentTime,
		CreatedBy:       "System",
		UpdatedAt:       currentTime,
		UpdatedBy:       "System",
	}

	data, err := repositories.CreateAnswer(context.Background(), questionnaireId, newAnswer)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.SuccessResult{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    data})
}

func UpdateAnswerHandler(c *gin.Context) {
	id := c.Param("id")
	var req dtoAssesment.UpdateAnswerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	if req.Description == "" || id == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	answerId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Answer ID"})
		return
	}

	existingAnswer, _ := repositories.GetAnswerById(context.Background(), answerId)
	if existingAnswer == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Not Exist"})
		return
	}

	currentTime := time.Now().In(time.FixedZone("UTC+7", 7*3600))

	updatedAnswer := dtoAssesment.UpdateAnswerRequest{
		Description: req.Description,
		UpdatedAt:   currentTime,
		UpdatedBy:   "System",
	}

	_, err = repositories.UpdateAnswerById(context.Background(), answerId, updatedAnswer)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "success",
		Data:    updatedAnswer})
}
