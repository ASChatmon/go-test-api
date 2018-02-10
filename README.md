# go-test-api

    This service is a directed graph that will represent all the workflow transitions.

## Running

    ./go-test-api

## Routes



## Test

    ./devtools/run-unit-tests

## Code Coverage

To get basic coverage:

    go test -coverprofile=coverage.out ./handlers/

To get details of that coverage, run after the above command:

    go tool cover -func=coverage.out

To get a line by line report, run this instead:

    go tool cover -html=coverage.out

## Lint

    ./devtools/fmt-lint

## Maintainer

    ASChatmon