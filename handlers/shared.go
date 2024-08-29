package handlers

import (
	"log/slog"
	"net/http"

	"github.com/VladanT3/IT_Ticketing_Platform/views/layouts"
	"github.com/a-h/templ"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request) error

func Make(h HTTPHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			slog.Error("HTTP handler error", "msg", err, "path", r.URL.Path)
		}
	}
}

func Render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	return c.Render(r.Context(), w)
}

func ShowError(w http.ResponseWriter, r *http.Request) error {
	err_msg := w.Header().Get("ErrorMessage")

	if LoggedInUserType != "" {
		return Render(w, r, layouts.ErrorMessage(LoggedInUserType, err_msg))
	} else {
		return Render(w, r, layouts.ErrorMessage("", err_msg))
	}
}
