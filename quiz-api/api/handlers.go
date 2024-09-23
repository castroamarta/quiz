package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func BasicAuthHandler(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if ok { 
		for user := range UserSecretsMock {
			if user == username && password == UserSecretsMock[user] {
				return
			}
		}
	}
	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "user authentication failed", http.StatusUnauthorized)
}

func GetQuestionsHander(w http.ResponseWriter, r *http.Request) {
	username, password, _ := r.BasicAuth()
	if err := validateCredentials(username, password); err != nil {
		http.Error(w, UnauthorizedErrorMessage, http.StatusUnauthorized)
		return
	}
	var quiz = QuizMock
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quiz)
}

func GetQuestionHander(w http.ResponseWriter, r *http.Request) {
	username, password, _ := r.BasicAuth()
	if err := validateCredentials(username, password); err != nil {
		http.Error(w, UnauthorizedErrorMessage, http.StatusUnauthorized)
		return
	}
	question, err := getQuestion(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(question)
}

func GetResultHandler(w http.ResponseWriter, r *http.Request) {
	username, password, _ := r.BasicAuth()
	if err := validateCredentials(username, password); err != nil {
		http.Error(w, UnauthorizedErrorMessage, http.StatusUnauthorized)
		return
	}
	result := fmt.Sprintf("Number of correct answers: %d", getQuizResult())
	fmt.Fprintln(w, result)
}

func GetStatsHandler(w http.ResponseWriter, r *http.Request) {
	username, password, _ := r.BasicAuth()
	if err := validateCredentials(username, password); err != nil {
		http.Error(w, UnauthorizedErrorMessage, http.StatusUnauthorized)
		return
	}
	stats := fmt.Sprintf("You were better than %.f%% of all quizzers", getQuizComparisonStats(username))
	fmt.Fprintln(w, stats)
}

func PostAnswersHandler(w http.ResponseWriter, r *http.Request) {
	username, password, _ := r.BasicAuth()
	if err := validateCredentials(username, password); err != nil {
		http.Error(w, UnauthorizedErrorMessage, http.StatusUnauthorized)
		return
	}
	answers := []Answer{}
    err := json.NewDecoder(r.Body).Decode(&answers)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	err = validateAnswers(answers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updateQuizStats(username)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(answers)
}

