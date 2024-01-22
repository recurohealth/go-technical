# Go Technical

![Golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

## Table of Contents

- [Technology Stack](#technology-stack)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running](#running)
  - [Usage](#usage)
- [Development](#development)
  - [Pre-commit Hooks](#pre-commit-hooks)

## Technology Stack

- HTTP web framework - [echo](https://github.com/labstack/echo)
- Environment variable loader - [godotenv](https://github.com/joho/godotenv)
- Database driver - [pgxpool (postgres)](https://github.com/jackc/pgx) or [mongo-go-driver](https://github.com/mongodb/mongo-go-driver)
- Database migration - [migrate](https://github.com/golang-migrate/migrate)
- Database query builder - [goqu](https://github.com/doug-martin/goqu)
- Data validator - [validator](https://github.com/go-playground/validator)
- Logger - [zerolog](https://github.com/rs/zerolog)
- Test assertions and mocks - [testify](https://github.com/stretchr/testify)
- Test data faking - [gofakeit](https://github.com/brianvoe/gofakeit/v6)

## Getting Started

### Prerequisites

Install containerization tool `docker`:

```bash
# Make sure to turn docker on after installation
brew install --cask docker

```

Install go version manager `gobrew`:

```bash
curl -sLk https://raw.githubusercontent.com/kevincobain2000/gobrew/master/git.io.sh | sh
```

Add `PATH` setting your shell configuration file (`.bashrc` or `.zshrc`):

```bash
export PATH="$HOME/.gobrew/current/bin:$HOME/.gobrew/bin:$PATH"
export GOROOT="$HOME/.gobrew/current/go"
```

Install go version:

```bash
gobrew use 1.20.4
```

Install `pre-commit` hooks:

```bash
make pc-install
```

### Installation

Install project dependencies:

```bash
make packages
```

It's also recommended to install [pre-commit hooks](#pre-commit-hooks) to statically analyze development work before pushing in any commits.

### Running

It's important to run these in the order as it is instructed below because the server attempts to connect to the database to run migration scripts.

Run the database:

```bash
make db
```

Run the server:

```bash
make server
```

### Usage

List of available endpoints:

```file
POST localhost:1323/api/v1/patients
```

An example of creating a patient:

```bash
curl -X POST 'localhost:1323/api/v1/patients' \
   -H 'Content-Type: application/json' \
   -d '{ "id": "4cb1eb83-197c-4e43-9522-0abb669914c6", "memberId": "123456789", "clientId": "2ba36ffb-22d5-4fdf-80ae-fe09393fbab9", "firstName": "Foo", "lastName": "Bar", "gender": "M", "dateOfBirth": "01-01-1990", "phoneNumber": "999-999-9999", "email": "foobar@gmail.com", "address": { "city": "Green Bay", "addressLineOne": "1234 Lone St", "addressLineTwo": "", "state": "WI", "zip": "55914", "country": "USA" } }'
```

## Development

### Tasks

Development tasks can be found in the `Makefile`. Currently available tasks are:

```bash
# Install pre-commit
make pc-install
# Run pre-commit on all files
make pc-run
# Add/remove dependencies based on imports
make packages
# Format go code
make format
# Run unit tests
make test
# Get test report of unit tests run
make test-report
# Stand up server
make server
# Stand up database
make db
```

These tasks standardize helpful and often used tasks so that we don't have to remember all the different tools and their commands.

### Pre-commit Hooks

Pre-commit hooks ensure that we are coding up to a standard and will also help fix some issues automatically (removing trailing whitespace, removing unused imports, etc). There are of course, issues that will have to be addressed manually and the hooks will point that out. This aligns with the concept of fail-fast as ideally we shouldn't have to wait until our code runs through the CI/CD pipeline to fail or reviewed by a colleague to be caught.

Install pre-commit hooks with the following command if you haven't already:

```bash
make pc-install
```

Notice that a `.pre-commit-config.yaml` file exists at the root (or will be created as a part of the above command if already doesn't exist). The script will use that configuration file to create a `pre-commit` file in `.git/hooks` so that whenever an attempt is made to push into a branch, the hooks will run to check if there are any errors. If errors exist, the push will fail and the errors will first have to be addressed.

After running the install pre-commit script, you can double-check that the pre-commit hooks will run as expected with the following command:

```bash
make pc-run
```

An error will usually occur if you don't have [gocyclo](https://github.com/fzipp/gocyclo) and [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports) installed in your `$GOPATH`. Both these packages can be installed with the following commands:

```bash
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
go install golang.org/x/tools/cmd/goimports@latest
```

It's also important to make sure to have the `$GOPATH` set up as well as `$PATH` pointing to your `$GOPATH`'s `bin` in your shell profile. For example:

```file
# .zshrc file

export GOPATH=$HOME/go
export PATH="$GOPATH/bin:$PATH"
```
