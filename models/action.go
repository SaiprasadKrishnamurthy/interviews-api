package models

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Action defines a standard function signature for us to use when creating
// controller actions. A controller action is basically just a method attached to // a controller.
type Action func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error
