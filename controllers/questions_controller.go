package controllers

import (
	"fmt"
	"io/ioutil"

	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/saiprasadkrishnamurthy/interviews-api/config"
	"github.com/saiprasadkrishnamurthy/interviews-api/models"
	"github.com/saiprasadkrishnamurthy/interviews-api/repositories"
)

// QuestionsController controller for Questions.
type QuestionsController struct {
	BaseController
}

// Questions from database.
// Questions from db.
// @Summary Get Questions by session id.
// @Description Get Questions by session id.
// @Produce  json
// @Param sessionId query string true "Session id"
// @Success 200 {object} models.Questions
// @Header 200 {string} Token "qwerty"
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /questions [get]
func (c *QuestionsController) Questions(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	sessionID := r.URL.Query().Get("sessionId")
	questions := repositories.GetQuestions(sessionID)
	rw.Header().Set("Content-Type", "application/json")
	questions.ToJSON(rw)
	return nil // no error
}

// QuestionVideo video from database.
// Questions.
// @Summary Get Question video by question id.
// @Description Get Question video by question id.
// @Produce  video/mp4
// @Accept json
// @Param questionId path string true "question id"
// @Success 200
// @Header 200 {string} Token "qwerty"
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /question/{questionId} [get]
func (c *QuestionsController) QuestionVideo(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	questionID := p.ByName("questionId")

	rw.Header().Set("Content-Type", "video/mp4")
	url := fmt.Sprintf(config.GetVideoQuestionsURL(), questionID)
	fmt.Println(url)
	rw.Write(downloadFile(url))
	return nil // no error
}

// SaveQuestionMetadata saves question metadata.
// SaveQuestionMetadata.
// @Summary SaveQuestionMetadata saves question metadata.
// @Description SaveQuestionMetadata saves question metadata.
// @Produce  json
// @Accept json
// @Param account body models.QuestionMetadata true "Add account"
// @Success 202
// @Header 200 {string} Token "qwerty"
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /question [put]
func (c *QuestionsController) SaveQuestionMetadata(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	go func() {
		questionMetadata := models.QuestionMetadata{}
		questionMetadata.FromJSON(r.Body)
		repositories.SaveQuestionMetadata(&questionMetadata)
	}()
	rw.WriteHeader(http.StatusAccepted)
	return nil // never fails cause it's async.
}

func downloadFile(url string) []byte {
	//Get the response bytes from the url
	response, err := http.Get(url)
	if err != nil {
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return body
}
