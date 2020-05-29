package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/saiprasadkrishnamurthy/interviews-api/config"
	"github.com/saiprasadkrishnamurthy/interviews-api/controllers"
)

const apiBase = "/api"

// InitialiseAllRoutes initialises all routes in the API.
func InitialiseAllRoutes(r *httprouter.Router) {
	// CORS.
	r.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", r.Header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})
	apiContext := apiBase + "/" + config.APIVersion()
	appController := controllers.BaseController{}

	// List all your controllers here.
	questionsController(apiContext, r, appController)
	sessionController(apiContext, r, appController)
	answerController(apiContext, r, appController)
}

func questionsController(apiContext string, r *httprouter.Router, baseController controllers.BaseController) {
	c := &controllers.QuestionsController{BaseController: baseController}
	r.GET(apiContext+"/questions", c.Action(c.Questions))
	r.GET(apiContext+"/question/:questionId", c.Action(c.QuestionVideo))
	r.PUT(apiContext+"/question", c.Action(c.SaveQuestionMetadata))
}

func sessionController(apiContext string, r *httprouter.Router, baseController controllers.BaseController) {
	c := &controllers.SessionController{BaseController: baseController}
	r.GET(apiContext+"/session/:sessionId", c.Action(c.Session))
	r.PUT(apiContext+"/session", c.Action(c.CreateInterviewSession))
}

func answerController(apiContext string, r *httprouter.Router, baseController controllers.BaseController) {
	c := &controllers.AnswerController{BaseController: baseController}
	r.PUT(apiContext+"/answer-video/:candidateId/:sessionId/:questionId", c.Action(c.SaveAnswerVideo))
	r.POST(apiContext+"/interview-completion/:candidateId/:sessionId", c.Action(c.InterviewCompleted))
}
