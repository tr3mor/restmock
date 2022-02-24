package restmock

import (
	"fmt"
	"net/http"
)

func NewHandlerFunc(interaction Interaction, l *HttpLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		l.logRequest(r, interaction.Response.StatusCode)
		switch interaction.Response.Type {
		case "plain":
			w.Header().Set("Content-Type", "text/plain")
		case "json":
			w.Header().Set("Content-Type", "application/json")
		}
		w.WriteHeader(interaction.Response.StatusCode)
		_, err := fmt.Fprintln(w, interaction.Response.Body)
		if err != nil {
			l.logger.Error("cant return response:", err)
		}

	}
}
