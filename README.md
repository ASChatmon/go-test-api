# go-test-api

    This service shows basic api structure with auth, logging, metrics, and db usage.

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
