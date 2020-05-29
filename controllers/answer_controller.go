package controllers

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
	"github.com/saiprasadkrishnamurthy/interviews-api/config"
	"github.com/saiprasadkrishnamurthy/interviews-api/models"
)

// AnswerController controller for Answer.
type AnswerController struct {
	BaseController
}

// SaveAnswerVideo as a video.
// SaveAnswerVideo.
// @Summary SaveAnswerVideo.
// @Description SaveAnswerVideo.
// @Produce  json
// @Accept multipart/form-data
// @Param  file formData file true  "this is a test file"
// @Param candidateId path string true "Candidate id"
// @Param sessionId path string true "Session id"
// @Param questionId path string true "Question id"
// @Success 200 {object} models.Session
// @Header 200 {string} Token "qwerty"
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /answer-video/{candidateId}/{sessionId}/{questionId} [put]
func (c *AnswerController) SaveAnswerVideo(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	candidateID, sessionID, questionID := p.ByName("candidateId"), p.ByName("sessionId"), p.ByName("questionId")
	clientFile, _, _ := r.FormFile("file") // r is *http.Request
	defer clientFile.Close()
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, clientFile)

	newpath := filepath.Join(config.GetConfig().VideoStore.WorkDir, sessionID, candidateID, questionID)
	os.MkdirAll(newpath, os.ModePerm)
	file, _ := os.Create(newpath + "/answer.webm")
	defer file.Close()

	bufferedWriter := bufio.NewWriter(file)
	bufferedWriter.Write(buf.Bytes())
	rw.WriteHeader(http.StatusCreated)
	return nil // no error
}

// InterviewCompleted InterviewCompleted
// InterviewCompleted.
// @Summary InterviewCompleted.
// @Description called when the interview is completed.
// @Produce  json
// @Accept json
// @Param candidateId path string true "Candidate id"
// @Param sessionId path string true "Session id"
//
// @Success 200
// @Header 200 {string} Token "qwerty"
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /interview-completion/{candidateId}/{sessionId} [post]
func (c *AnswerController) InterviewCompleted(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	candidateID, sessionID := p.ByName("candidateId"), p.ByName("sessionId")
	interviewCompletedEvent := models.InterviewCompletedEvent{SessionID: sessionID, CandidateID: candidateID}
	json, _ := json.Marshal(interviewCompletedEvent)
	natsConn := config.GetNatsConnection()
	natsSubject := config.GetConfig().Nats.InterviewCompletedSubject
	err := natsConn.Publish(natsSubject, json)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Published message %s on subject %s ", json, natsSubject)

	rw.WriteHeader(http.StatusAccepted)
	interviewCompletedEvent.ToJSON(rw)
	return nil // no error
}
