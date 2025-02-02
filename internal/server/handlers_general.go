package server

import (
	"github.com/gorilla/websocket"
	"github.com/reaper47/recipya/internal/app"
	"github.com/reaper47/recipya/internal/models"
	"github.com/reaper47/recipya/web/components"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(_ *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (s *Server) downloadHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file := r.PathValue("tmpFile")
		data, err := s.Files.ReadTempFile(file)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", http.DetectContentType(data))
		w.Header().Set("Content-Disposition", `attachment; filename="`+file+`"`)
		w.Header().Set("Content-Length", strconv.Itoa(len(data)))
		_, _ = w.Write(data)
	}
}

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	if app.Config.Server.IsAutologin || isAuthenticated(r, s.Repository.GetAuthToken) {
		middleware := s.mustBeLoggedInMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			s.recipesHandler().ServeHTTP(w, r)
		}))
		middleware.ServeHTTP(w, r)
		return
	}

	http.Redirect(w, r, "/guide", http.StatusSeeOther)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_ = components.SimplePage("Page Not Found", "The page you requested to view is not found. Please go back to the main page.").Render(r.Context(), w)
}

func (s *Server) userInitialsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(s.Repository.UserInitials(getUserID(r))))
	}
}

func (s *Server) wsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			w.Header().Set("HX-Trigger", models.NewWarningToast("", "Could not upgrade connection.", "").Render())
			return
		}

		userID := getUserID(r)
		broker := models.NewBroker(userID, s.Brokers, ws)
		s.Brokers[userID] = broker
	}
}
