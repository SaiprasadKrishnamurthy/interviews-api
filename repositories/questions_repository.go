package repositories

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/saiprasadkrishnamurthy/interviews-api/config"
	"github.com/saiprasadkrishnamurthy/interviews-api/models"
)

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
func SaveQuestionMetadata(questionMetadata *models.QuestionMetadata) error {
	client := config.ElasticClient()
	configuration := config.GetConfig()
	elasticConf := configuration.Elastic
	uri := elasticConf.URI + elasticConf.QuestionsMetadataIndex + "/_doc/" + questionMetadata.QuestionID
	fmt.Println(uri)
	json, _ := json.Marshal(questionMetadata)
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(configuration.Elastic.Username, configuration.Elastic.Password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	fmt.Println(s)
	return err
}
