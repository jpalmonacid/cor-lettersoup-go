# Lettersoup - Home test for COR
This project contains the code which solves the [given exercise](test-back.pdf) sent by COR.

## Requirements
* [docker](https://docs.docker.com/engine/install/)

## Run the unit tests

Run the following command from the project's root directory to run the tests:
```
docker run -it --rm -v"$PWD":/go golang:1.15-alpine3.12 sh -c "go test -v"
```

This should run all tests found at [lettersoup_test.go](lettersoup_test.go). All
test cases described on the pdf are found in this file.
You should find the actual code at [lettersoup.go](lettersoup.go).
