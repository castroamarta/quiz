package api

import "fmt"

func getQuizResult() int {
	numCorrectAnswers := 0
	for questionName := range UserSolution {
		if UserSolution[questionName] == SolutionMock[questionName] {
			numCorrectAnswers += 1
		}
	}
	return numCorrectAnswers
}

func getQuizStats(userID string) float64 {
	numCorrectAnswers := getQuizResult()
	StatsMock[userID] = float64(numCorrectAnswers*100) / float64(NumQuizQuestions)
	quizStat := StatsMock[userID]
	numUsers := 0 // number of quizzers that performed worse
	for userID := range StatsMock {
		if StatsMock[userID] < quizStat {
			numUsers += 1
		}
	}
	return float64((numUsers)*100) / float64(len(StatsMock)-1)
}

func validateAnswers(answers []Answer) error {
	if len(answers) == 0 {
		return fmt.Errorf("no answers provided")
	}
	for _, answer := range answers {
		err := validateAnswer(answer)
		if err != nil {
			return err
		}
		UserSolution[answer.QuestionID] = answer.OptionID
	}
	return nil
}

func validateAnswer(answer Answer) error {
	question, err := getQuestion(answer.QuestionID)
	if err != nil {
		return err
	}
	options := question.Options
	isValidOption := false
	for _, op := range options {
		if answer.OptionID == op.ID {
			isValidOption = true
			break
		}
	}
	if !isValidOption {
		return fmt.Errorf("question %v only has the options %s", answer.QuestionID, options)
	}
	return nil
}

func getQuestion(id string) (Question, error) {
	question := Question{}
	if id == "" {
		return question, fmt.Errorf("no question ID provided")
	}
	isValidQuestion := false
	for _, q := range QuizMock {
		if id == q.ID {
			question = q
			isValidQuestion = true
			break
		}
	}
	if !isValidQuestion {
		return question, fmt.Errorf("invalid question ID provided: %v", id)
	}
	return question, nil
}

func validateAPIKey(APIKey string) error {
	for key := range UserAPIKeysMock {
		if UserAPIKeysMock[key] == APIKey {
			return nil
		}
	}
	return fmt.Errorf("user unauthorized to perform method call")
}