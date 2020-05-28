package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/saiprasadkrishnamurthy/interviews-api/models"
	"github.com/saiprasadkrishnamurthy/interviews-api/repositories"
)

// SessionController controller for Session.
type SessionController struct {
	BaseController
}

// Session from database.
// Session.
// @Summary Get Session by session id.
// @Description Get Session by session id.
// @Produce  json
// @Param sessionId path string true "Session id"
// @Success 200 {object} models.Session
// @Header 200 {string} Token "qwerty"
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /session/{sessionId} [get]
func (c *SessionController) Session(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	sessionID := p.ByName("sessionId")
	if session := repositories.GetSession(sessionID); session != nil {
		rw.Header().Set("Content-Type", "application/json")
		session.ToJSON(rw)
	} else {
		rw.WriteHeader(http.StatusNotFound)
	}
	return nil // no error
}

// CreateInterviewSession creates an interview session.
// creates an interview session.
// @Summary creates an interview session.
// @Description creates an interview session.
// @Produce  json
// @Accept json
// @Param account body models.Session true "creates an interview session"
// @Success 202
// @Header 200 {string} Token "qwerty"
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /session [put]
func (c *SessionController) CreateInterviewSession(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	go func() {
		session := models.Session{}
		session.FromJSON(r.Body)
		repositories.CreateSession(&session)
	}()
	rw.WriteHeader(http.StatusAccepted)
	return nil // never fails cause it's async.
}
