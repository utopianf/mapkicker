package web

import (
	"encoding/json"
	"log"
	"mapkicker/domain"
	"net/http"
)

// App is application
type App struct {
	r        domain.Repository
	handlers map[string]http.HandlerFunc
}

// NewApp creates new app object
func NewApp(r domain.Repository, cors bool) App {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // TODO: コンストラクタにしてまとめる
	app := App{
		handlers: make(map[string]http.HandlerFunc),
	}
	mappoolHandler := app.GetMappool
	if !cors {
		mappoolHandler = disableCors(mappoolHandler)
	}
	app.handlers["/api/mappool"] = mappoolHandler
	app.handlers["/echo"] = Echo
	app.handlers["/"] = http.FileServer(http.Dir("/webapp")).ServeHTTP
	return app
}

// Serve listens to port 8080
func (a *App) Serve() error {
	r := newRoom()
	http.Handle("/room", r)
	for path, handler := range a.handlers {
		http.Handle(path, handler)
	}
	log.Println("Web server is available on port 8080")
	return http.ListenAndServe(":8080", nil)
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
