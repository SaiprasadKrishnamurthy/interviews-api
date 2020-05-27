package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/saiprasadkrishnamurthy/interviews-api/models"
)

// BaseController This is our Base Controller
type BaseController struct {
}

// Action The action function helps with error handling in a controller
func (c *BaseController) Action(a models.Action) httprouter.Handle {
	return httprouter.Handle(func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if err := a(rw, r, p); err != nil {
			http.Error(rw, err.Error(), 500)
		}
	})
}
