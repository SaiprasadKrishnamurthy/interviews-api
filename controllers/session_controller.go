package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
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
	session := repositories.GetSession(sessionID)
	rw.Header().Set("Content-Type", "application/json")
	session.ToJSON(rw)
	return nil // no error
}
