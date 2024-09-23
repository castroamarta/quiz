# Quiz API

### Description

Directory that contains the quiz endpoints.

### Endpoints

The table bellow contains the endpoins supported.

Request | Endpoint  | Description
------------- | ------------- | -------------
GET | `curl -u alice:rainbow http://localhost:8081/auth`  | Performs basic authentication and returns an API Key
GET | `curl -u alice:rainbow 'http://localhost:8081/questions' --request "GET"`  | Gets the quiz form
GET | `curl -u alice:rainbow 'http://localhost:8081/question?id=1' --request "GET"`| Gets a quiz question
POST | `curl -u alice:rainbow http://localhost:8081/answers --request "POST" --data '[{"question_id": "1","option_id": "a"}]'`  | Answers the quiz questions
GET | `curl -u alice:rainbow 'http://localhost:8081/result' --request "GET"`  | Gets the total number of correct quiz answers
GET | `curl -u alice:rainbow 'http://localhost:8081/stats' --request "GET"`  | Gets the quiz stats

### Authentication

This implementation relies on a basic username password authentication.

The table bellow contains the registered mocked credentials that can be used to test these endpoints.

Username  | Password  | 
------------ | ------------ |  
`alice` | `rainbow` | 
`bob` | `flower` | 
`eve` | `boat` | 