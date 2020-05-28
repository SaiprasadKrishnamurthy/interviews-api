package repositories

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/saiprasadkrishnamurthy/interviews-api/config"
	"github.com/saiprasadkrishnamurthy/interviews-api/models"
)

// GetQuestions get questions.
func GetQuestions(sessionID string) models.Questions {
	client := config.ElasticClient()
	configuration := config.GetConfig()
	elasticConf := configuration.Elastic
	uri := elasticConf.URI + elasticConf.QuestionsMetadataIndex + "/_search"
	queryJSON := `{
		"query": {
		  "bool": {
			"filter": [
			  {
				"term": {
				  "sessionId": "%s"
				}
			  }
			]
		  }
		},
		"sort": [
		  {
			"sequence": {
			  "order": "asc"
			}
		  }
		]
	  }`
	query := fmt.Sprintf(queryJSON, sessionID)

	req, err := http.NewRequest("POST", uri, strings.NewReader(query))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(configuration.Elastic.Username, configuration.Elastic.Password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	var response map[string]interface{}
	json.Unmarshal(bodyText, &response)
	h := response["hits"].(map[string]interface{})
	hits := h["hits"].([]interface{})
	questions := []models.Question{}
	for _, hit := range hits {
		obj := hit.(map[string]interface{})
		src := obj["_source"]
		srcJSON, _ := json.Marshal(src)
		qm := models.QuestionMetadata{}
		json.Unmarshal(srcJSON, &qm)
		questions = append(questions, qm.Question)
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
	_, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		log.Println(e)
	}
	return e
}
