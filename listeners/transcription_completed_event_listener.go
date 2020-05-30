package listeners

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/saiprasadkrishnamurthy/interviews-api/config"
	"github.com/saiprasadkrishnamurthy/interviews-api/repositories"
	"github.com/saiprasadkrishnamurthy/interviews-api/utils"

	textmatch "github.com/SaiprasadKrishnamurthy/text-match"
	"github.com/nats-io/nats.go"
	"github.com/saiprasadkrishnamurthy/interviews-api/models"
)

// TranscriptionCompletedReceivedEventListener foo listener.
type TranscriptionCompletedReceivedEventListener struct {
	*BaseListener
}

// Handle on message function of all the listeners.
func (l *TranscriptionCompletedReceivedEventListener) Handle(msg *nats.Msg) {
	event := models.TranscriptionCompletedEvent{}
	json.Unmarshal(msg.Data, &event)
	_, questions := repositories.GetQuestions(event.SessionID)

	fmt.Println("\n\n\n\n ")
	for _, tr := range event.TranscriptionResults {
		question := utils.Filter(questions,
			func(q models.QuestionMetadata) string {
				return q.Question.QuestionID
			},
			func(qid string) bool {
				return qid == tr.Question
			},
		)[0] // First match. There must be one.

		expectedAnswer := question.AnswerText
		expectedKeywords := question.ImportantKeywords
		answerTxtFile := filepath.Join(config.GetConfig().VideoStore.WorkDir, tr.SessionID, tr.CandidateID, tr.Question, "answer.txt")
		actualAnswerBytes, _ := ioutil.ReadFile(answerTxtFile)
		actualAnswer := string(actualAnswerBytes)
		answerScore := textmatch.Similarity(expectedAnswer, actualAnswer, true)
		keywordsScore := textmatch.Similarity(expectedKeywords, actualAnswer, true)

		qr := models.QuestionResult{
			SessionID:                         event.SessionID,
			CandidateID:                       event.CandidateID,
			QuestionID:                        tr.Question,
			AutoAnswerSimilarityScore:         answerScore.CosineSimilarityScore,
			AutoAnswerAbsoluteMatchingScore:   answerScore.AbsoluteSimilarityScore,
			AutoKeywordsSimilarityScore:       keywordsScore.CosineSimilarityScore,
			AutoKeywordsAbsoluteMatchingScore: keywordsScore.AbsoluteSimilarityScore,
			Confidence:                        tr.Confidence,
		}

		fmt.Printf(" Candidate: %s, Question: %s, Answer Score: %#v, Keyword Score: %#v \n\n\n", event.CandidateID, tr.Question, answerScore, keywordsScore)

		// Async save.
		go repositories.SaveQuestionResult(&qr)

	}

}
