package controllers

import (
	"fmt"
	"io/ioutil"

	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/saiprasadkrishnamurthy/interviews-api/config"
	"github.com/saiprasadkrishnamurthy/interviews-api/repositories"
)

// QuestionsController controller for Questions.
type QuestionsController struct {
	BaseController
}

// Questions from database.
// Questions.
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
// @Produce  json
// @Param questionId query string true "question id"
// @Success 200
// @Header 200 {string} Token "qwerty"
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /question/{questionName}/{questionId} [get]
func (c *QuestionsController) QuestionVideo(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	questionID := p.ByName("questionId")

	rw.Header().Set("Content-Type", "video/mp4")
	url := fmt.Sprintf(config.GetVideoQuestionsURL(), questionID)
	fmt.Println(url)
	rw.Write(downloadFile(url))
	return nil // no error
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
