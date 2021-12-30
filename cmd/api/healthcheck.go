package main

import (
	"net/http"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "the server encountered a problem and could not process the request", http.StatusInternalServerError)
	}

}
