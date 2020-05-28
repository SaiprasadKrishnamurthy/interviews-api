package repositories

import "github.com/saiprasadkrishnamurthy/interviews-api/models"

// GetQuestions get questions.
func GetQuestions(sessionID string) models.Questions {
	questions := []models.Question{
		models.Question{
			SessionID:           sessionID,
			Sequence:            1,
			QuestionID:          "question1",
			QuestionText:        "QuestionText1",
			AnswerTimeInSeconds: 60,
		},
		models.Question{
			SessionID:           sessionID,
			Sequence:            2,
			QuestionID:          "question2",
			QuestionText:        "QuestionText2",
			AnswerTimeInSeconds: 60,
		},
	}
	return models.Questions{Questions: questions}
}

// SaveQuestionMetadata get questions.
func SaveQuestionMetadata(questionMetadata models.QuestionMetadata) error {

}
