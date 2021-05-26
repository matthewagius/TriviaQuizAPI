# Quiz API

## About the Project

A simple quiz with 10 questions, having one correct answer per question, composed of REST API and CLI that talks with the API

Just in memery, so no database.

## Requirements

- user should be presented with questions having multiple answers
- user should be able to select one answer per question
- user should be able to answer all the questions and then post his answers to get back how many correct answers they had and be displayed to the user
- user should see how well he rated compared to others that have taken the quiz "You scored higher than 60% of all quizzers

## Built With

- [Go](https://golang.org/)
- [gofiber](https://github.com/gofiber/fiber)

## Prerequisites

- go installed on your machine

## Running this API

- open TriviaQuestionsAPI
- run $ go run main.go
- open browser to http://localhost:8000/swagger/ - api documentation can be viewed here via Swagger

## References

- https://dev.to/devsmranjan/golang-build-your-first-rest-api-with-fiber-24eh
- https://github.com/stretchr/testify
- https://github.com/arsmn/fiber-swagger
