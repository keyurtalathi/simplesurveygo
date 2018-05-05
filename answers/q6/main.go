package main

import (
	"net/http"
	"simplesurveygo/answers/q6"
	"simplesurveygo/servicehandlers"
)

// function to limit requests using semaphores
func maxClients(h http.Handler, sema chan struct{}) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sema <- struct{}{}
		defer func() { <-sema }()

		h.ServeHTTP(w, r)
	})
}

func main() {

	var sema chan struct{}
	sema = make(chan struct{}, 10)
	// Serves the html pages
	http.Handle("/", http.FileServer(http.Dir("./static")))

	pingHandler := servicehandlers.PingHandler{}
	authHandler := servicehandlers.UserValidationHandler{}
	sessionHandler := servicehandlers.SessionHandler{}
	surveyHandler := servicehandlers.SurveyHandler{}
	userSurveyHandler := servicehandlers.UserSurveyHandler{}
	signupHandler := servicehandlers.SignupHandler{}

	// Serves the API content
	http.Handle("/api/v1/ping/", maxClients(pingHandler, sema))

	http.Handle("/api/v1/signup/", maxClients(signupHandler, sema))
	http.Handle("/api/v1/authenticate/", maxClients(authHandler, sema))
	http.Handle("/api/v1/validate/", maxClients(sessionHandler, sema))

	http.Handle("/api/v1/survey/{surveyname}", maxClients(surveyHandler, sema))
	http.Handle("/api/v1/survey/", maxClients(surveyHandler, sema))

	http.Handle("/api/v1/usersurvey/", maxClients(userSurveyHandler, sema))
	go q6.Deactivete_serveys()
	// Start Server
	http.ListenAndServe(":3000", nil)
}
