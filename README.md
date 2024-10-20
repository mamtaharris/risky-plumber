# Risky Plumber 

As the Risky Plumbers team works primarily in the creation of web services, this assignment is intended to allow you to demonstrate your abilities towards this end. 
This assignment can be placed in a GitHub repository (please use a GH generic name for the repo) or e-mailed to the recruiter in a zip file. 
Please treat this as if you were writing production quality code. 
Task 
The task is to create a Risk application which: 
. Listens on port 8080 for standard HTTP traffic (not HTTPS); and 
2. Can return a list of Risk given a GET to /v1/risks on the endpoint; and . Can create a new Risk given a POST to /v1/risks on the endpoint; and . Can retrieve an individual Risk given a GET to /v1/risks/<id> . 
Risk Information 
A Risk should consist of: 
A Risk ID in the form of a UUID. 
This ID should be auto-generated on creation. 
A state value as a string 
This can be one of [open, closed, accepted, investigating] 
MUST be present for all Risks. 
A Risk title as a string. 
A Risk description as a string. 
Additional Information
The endpoints are NOT secured for the purposes of this assignment. Data transfer should be done in JSON. 
The storage of risks can be done in memory; no database creation or usage is required. 
The endpoints should use standard HTTP response status codes. E.g.: 200 OK for a successful GET to the /v1/risks/ or /v1/risks/<id> endpoints 500 Internal Server Error for problem with the internal server; etc. 
You may use any Golang framework or library to assist in the creation of this app. Please include: 
Instructions via a README.md or README.adoc file including how to run your service and any tests you may have written. 
In a For-Reviewers.md file, include any notes, thoughts, etc. which you would like the interviewers to be aware of.

Please check thw wiki page for more details.

## Getting Started

Follow below steps to get the service running on your local machine.

### Prerequisites

- Golang

### Local setup

Open terminal in your project root directory and run below command

```bash
make server
```

## Project Overview

- cmd - This package contains different commands that the service provides. Currently, we can start server, migrate the database and seed some user info.

- config - This directory contains configuration files or code related to your application's configuration, such as database connections or environment variables.

- internal - This directory is added at the root level and contains packages that should not be imported or accessed by external packages. It can include packages like internal models, repositories, services, or any other internal implementation details.

- pkg - This directory is used to store packages that can be potentially reused across multiple applications. In this case, we have two subdirectories: /database, /logger, /server and /middlewares.

- main.go - This is the entry point of the application which calls all the initialization scripts and starts the http server.

- Makefile - This file defines the set of tasks to be executed.

- RiskyPlumber.postman_collection.json - This file can be imported to your postman workspace.

## Makefile

Run the application

```bash
make server
```

Run linters

```bash
make lint
```

Run all tests. Current test coverage is ~95%

```bash
make test
```

Generates mock files for unit testing

```bash
make mock
```
