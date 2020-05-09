package web

import (
	"encoding/json"
	"log"
	"mapkicker/db"
	"net/http"
)

// App is application
type App struct {
	d        db.DB
	handlers map[string]http.HandlerFunc
}

// NewApp creates new app object
func NewApp(d db.DB, cors bool) App {
	app := App{
		d:        d,
		handlers: make(map[string]http.HandlerFunc),
	}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	r := newRoom()
	techHandler := app.GetTechnologies
	mappoolHandler := app.GetMappool
	roomHandler := r.ServeHTTP
	if !cors {
		techHandler = disableCors(techHandler)
		mappoolHandler = disableCors(mappoolHandler)
		// roomHandler = disableCors(roomHandler)
	}
	app.handlers["/api/technologies"] = techHandler
	app.handlers["/api/mappool"] = mappoolHandler
	app.handlers["/room"] = roomHandler
	app.handlers["/"] = http.FileServer(http.Dir("/webapp")).ServeHTTP
	return app
}

// Serve listens to port 8080
func (a *App) Serve() error {
	for path, handler := range a.handlers {
		http.Handle(path, handler)
	}
	log.Println("Web server is available on port 8080")
	return http.ListenAndServe(":8080", nil)
}

// GetTechnologies returns technologies
func (a *App) GetTechnologies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	technologies, err := a.d.GetTechnologies()
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(technologies)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

// GetMappool returns mappool
func (a *App) GetMappool(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mappools := []string{
		"Eternal Empire LE",
		"Ever Dream LE",
		"Golden Wall LE ",
		"Nightshade LE",
		"Purity and Industry LE",
		"Simulacrum LE",
		"Zen LE",
	}
	err := json.NewEncoder(w).Encode(mappools)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

// Needed in order to disable CORS for local development
func disableCors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		h(w, r)
	}
}
