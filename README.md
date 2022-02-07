## My Theresa Test

***All of this you must run into app directory***

#### First step to run app
You must have `docker` running on your machine

#### Run tests
important integration tests works with a docker container inside of `.it` directory

`make test`

#### Run application 
`make app-up`

#### Down application
`make app-down`

**Workaround if you can't run with commands above**

Run tests
1. `cd .it`
2. `docker-compose up`
3. `go test ./...`

Run Application
1. First stop docker container of tests `cd .it && docker-compose down`
2. In app root directory `docker-compose up`

Ready! You can test if all went good with
`curl -X GET 127.0.0.1:8080/products`

###Steps

#### Infrastructure with docker compose:
- Created a postgres database from an image
- Simple API with a “hello world” response from an endpoint and db connection for very basic infrastructure test from url
- Created a migration.sql file that will run after postgres image is up to create database, tables, and insert the data from the exercise
- Test of db connection

#### Service API structure:
- A simple structure with a success route from handler to repository
- Creation of requests and responses that will be used in the respective layers.

#### Repository layer:
- Started with a repository layer creating the sql queries to represent the business needed.
- Created a basic layer of db that will have only queries that will be executed in the repository.
- Filled repository GetProducts method with all necesary to return slice of Product structs
- Creation of environment for integration tests with docker compose and added makefile to throw tests
- Added unit tests for queries and integration tests for repository
- Added app up and down to Makefile, stopping test environment to ensure it will run successfully

#### Service Layer:
- Created apply discount helper to apply percentage discount to a price
- Logic to get products method in service
- Created discounter struct to manage the discounts logic and fill price struc
- tested with use cases

#### Handler
- Get Products handler logic
- Made helper function to parse query string to struct request
- I can’t do tests for this, but I would have done integration tests
