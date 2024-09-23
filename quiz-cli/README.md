# Quiz CLI

### Description

Directory that contains the quiz command line interface.

### Endpoints

The table bellow contains the commands supported.

Command | Example
------------- | ------------- 
 auth | `./bin/quiz auth --username alice --password rainbow`  
 get-questions | `./bin/quiz get-questions`  
 get-question  | `./bin/quiz get-question 1`
 select-options | `./bin/quiz select-options 1:b,2:b,3:a`
 get-result  | `./bin/quiz get-result`
 get-stats  | `./bin/quiz get-stats`