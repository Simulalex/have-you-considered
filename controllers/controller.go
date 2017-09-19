package controllers

import (
	"github.com/Simulalex/haveyouconsidered/db"
	"github.com/gorilla/mux"
	"net/http"
	"html/template"
)

type SummaryHandler struct {
	db db.Database
}

type Config struct {
	Db db.Config
}

func Run(cfg *Config) error {
	db, err := db.NewDatabase(&cfg.Db);
	if err != nil {
		return err
	}

	summaryHandler := SummaryHandler{db}

	router := mux.NewRouter()
	router.HandleFunc("/", GetLanding)
	router.HandleFunc("/{name}", summaryHandler.GetSummary)

	return http.ListenAndServe(":8080", router)
}

func (handler SummaryHandler) GetSummary(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	technology, err := handler.db.GetSummary(name)
	if err != nil {
		http.ServeFile(w, r, "views/404.html")
	}

	t, _ := template.ParseFiles("views/summary.html")
	t.Execute(w, technology)
}

func GetLanding(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "views/index.html")
}
