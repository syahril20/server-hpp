package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupAssesmentRoutes(r *gin.Engine) {
	Assesment := r.Group("/assesment")
	{
		Assesment.GET("/", handlers.GetAllActiveAssesmentsHandler)
		Assesment.GET("/non-active", handlers.GetAllNonActiveAssesmentsHandler)
		Assesment.POST("/", handlers.CreateAssesmentHandler)
		Assesment.PUT("/:id", handlers.UpdateAssesmentHandler)
		Assesment.PUT("/active/:id", handlers.ActiveAssesmentHandler)
		Assesment.DELETE("/:id", handlers.DeleteAssesmentHandler)
	}

	Questionnaire := r.Group("/questionnaire")
	{
		Questionnaire.GET("/:id", handlers.GetAllActiveQuestionnaireHandler)
		Questionnaire.GET("/non-active/:id", handlers.GetAllNonActiveQuestionnaireHandler)
		Questionnaire.POST("/", handlers.CreateQuestionnaireHandler)
		Questionnaire.PUT("/:id", handlers.UpdateQuestionnaireHandler)
		Questionnaire.PUT("/active/:id", handlers.ActiveQuestionnaireHandler)
		Questionnaire.DELETE("/:id", handlers.DeleteQuestionnaireHandler)
	}

	Answer := r.Group("/answer")
	{
		Answer.GET("/:id", handlers.GetAllActiveAnswersHandler)
		Answer.POST("/", handlers.CreateAnswerHandler)
		Answer.PUT("/:id", handlers.UpdateAnswerHandler)
	}
}
