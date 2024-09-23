package main

import (
	"log"
	"net/http"
	"quiz-api/api"
	"time"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/auth", api.BasicAuthHandler)
	
	mux.HandleFunc("/questions", api.GetQuestionsHander)
	mux.HandleFunc("/question", api.GetQuestionHander)
	mux.HandleFunc("/result", api.GetResultHandler)
	mux.HandleFunc("/stats", api.GetStatsHandler)
	mux.HandleFunc("/answers", api.PostAnswersHandler)
    srv := &http.Server{
        Addr:         ":8081",
        Handler:      mux,
        IdleTimeout:  time.Minute,
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 30 * time.Second,
    }
	log.Printf("starting server on %s", srv.Addr)
	http.ListenAndServe(srv.Addr, mux)
}
