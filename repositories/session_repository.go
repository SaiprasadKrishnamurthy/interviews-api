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

// GetSession get session.
func GetSession(sessionID string) *models.Session {
	client := config.ElasticClient()
	configuration := config.GetConfig()
	elasticConf := configuration.Elastic
	uri := elasticConf.URI + elasticConf.SessionIndex + "/_search"
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
		}
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
	sessions := []models.Session{}
	for _, hit := range hits {
		obj := hit.(map[string]interface{})
		src := obj["_source"]
		srcJSON, _ := json.Marshal(src)
		sess := models.Session{}
		json.Unmarshal(srcJSON, &sess)
		sessions = append(sessions, sess)
	}
	if len(sessions) == 0 {
		return nil
	}
	return &sessions[0]
}

// CreateSession save an interview session.
func CreateSession(session *models.Session) error {
	client := config.ElasticClient()
	configuration := config.GetConfig()
	elasticConf := configuration.Elastic
	uri := elasticConf.URI + elasticConf.SessionIndex + "/_doc/" + (*session).SessionID
	fmt.Println(uri)
	json, _ := json.Marshal(session)
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
