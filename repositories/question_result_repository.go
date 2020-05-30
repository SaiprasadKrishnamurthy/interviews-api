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

// SaveQuestionResult get questions.
func SaveQuestionResult(qr *models.QuestionResult) error {
	client := config.ElasticClient()
	configuration := config.GetConfig()
	elasticConf := configuration.Elastic
	q := (*qr)
	docID := strings.Join([]string{q.SessionID, q.CandidateID, q.QuestionID}, "_")
	uri := elasticConf.URI + elasticConf.QuestionResultIndex + "/_doc/" + docID
	fmt.Println(uri)
	json, _ := json.Marshal(qr)
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(configuration.Elastic.Username, configuration.Elastic.Password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	r, e := ioutil.ReadAll(resp.Body)
	fmt.Println(string(r))
	if e != nil {
		log.Fatal(e)
	}
	return e
}
