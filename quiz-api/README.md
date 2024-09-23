# Quiz API

### Description

Directory that contains the quiz endpoints.

### Endpoints

The table bellow contains the endpoins supported.

Request | Endpoint  | Description
------------- | ------------- | -------------
GET | `curl -u alice:rainbow http://localhost:8081/auth`  | Performs basic authentication and returns an API Key
GET | `curl 'http://localhost:8081/questions' --request "GET" -H 'X-API-Key: VAFJWEKSFS'`  | Gets the quiz form
GET | `curl 'http://localhost:8081/question?id=1' --request "GET" -H 'X-API-Key: VAFJWEKSFS'`| Gets a quiz question
POST | `curl http://localhost:8081/answers --request "POST" --data '[{"question_id": "1","option_id": "a"}]' -H 'X-API-Key: VAFJWEKSFS'`  | Answers the quiz questions
GET | `curl 'http://localhost:8081/result' --request "GET" -H 'X-API-Key: VAFJWEKSFS'`  | Gets the total number of correct quiz answers
GET | `curl 'http://localhost:8081/stats' --request "GET" -H 'X-API-Key: VAFJWEKSFS'`  | Gets the quiz stats

### Authentication

This implementation relies on an authentication via API Key, provided and linked to a specific user when registering in the client application.

The table bellow contains the registered mocked api keys that can be used to test these endpoints.

API Key  | 
------------ | 
`VAFJWEKSFS` | 
`FEJRGIERGJ` | 
`PQIENFJRGR` | 