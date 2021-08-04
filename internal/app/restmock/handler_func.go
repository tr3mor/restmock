package restmock

import (
	"fmt"
	"log"
	"net/http"
)

func NewHandlerFunc(interaction Interaction, l *HttpLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch interaction.Response.Type {
		case "plain":
			w.Header().Set("Content-Type", "text/plain")
		case "json":
			w.Header().Set("Content-Type", "application/json")
		}
		w.WriteHeader(interaction.Response.StatusCode)
		_, err := fmt.Fprintln(w, interaction.Response.Body)
		if err != nil {
			log.Fatalln("cant return response")
		}
		l.logRequest(r, interaction.Response.StatusCode)
	}
}
