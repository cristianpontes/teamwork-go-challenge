![alt text](https://www.teamwork.com/app/themes/teamwork-theme/dist/images/twork-slate.svg "Teamwork")
[![Go Report Card](https://goreportcard.com/badge/github.com/cristianpontes/teamwork-go-challenge)](https://goreportcard.com/report/github.com/cristianpontes/teamwork-go-challenge)
[![Build Status](https://travis-ci.org/cristianpontes/teamwork-go-challenge.svg?branch=master)](https://travis-ci.org/cristianpontes/teamwork-go-challenge)

# Teamwork - Go Challenge
This repository contains my solution for the Teamwork coding challenge assignment.

## Original Spec
```
// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
```

## Assumptions and Versions

The way I understood the spec, was that goal was to _group_ customers by their email's domain. However, I wasn't sure whether the expected output was

`gmail.com -> X` where `x` would be `len(customersUsingGmail)` 

OR

`gmail.com -> {X, Y, Z, ...}` where `x` `y` `z` represent full customer objects

So, I implemented both views (examples bellow)

Although, some examples of the expected output would have been quite useful for this assignment

**Go Version**
```
$ go version
go version go1.14 darwin/amd64
```

## Project Structure
- `cmd/{command-name}` entry point for the application via CLI
- `pkg/{package-name}` concrete implementations and shared packages across the project
- `pkg/{package-name}/testing/mocks` exposes a mockable instance that fullfils the package's contract

## Project Set Up / Useful Commands
- Configuring the project locally: `make config`
- Installing the project locally: `make install`
- Running linters: `make lint`
- Running tests: `make test`
- Running test coverage report (UI) `make coverage-report`

## Project Features
- Abstraction between concrete implementations (`./pkg`) and the presentation layer (`./cmd`)
- Packages rely on interfaces/contracts which facilitates dependency injection and testing
- The project is fully covered with tests
- Benchmarks for the core feature can be found in `./pkg/customer/group_test.go`


## Walkthrough - Running the project

First, install the project `make config && make install`

The project will be installed as a command called `tw-go-challenge`

CLI instructions can be found via `tw-go-challenge importer --help` 

```
Customer importer reads from a csv file and returns a sorted list of email domains along with the number of customers with e-mail addresses for each domain.

Usage:
  tw-go-challenge importer [flags]

Flags:
      --detailed-report   Show detailed import report by showing the full list of customers by domain email
      --file string       Path to a CSV file containing the list of the customers to be imported.
  -h, --help              help for importer

```

### Customer importer - Simplified Report
```
$ tw-go-challenge importer  --file ./cmd/customer/importer/testing/stubs/customer-import.csv    

sourceforge.net -> 7
foxnews.com -> 5
chicagotribune.com -> 7
sciencedirect.com -> 4
...
```

### Customer importer - Detailed Report
```
$ tw-go-challenge importer  --file ./cmd/customer/importer/testing/stubs/customer-import.csv --detailed-report=true

  "360.cn": [
    {
      "first_name": "Justin",
      "last_name": "Hansen",
      "email": "jhansen3@360.cn",
      "gender": "Male",
      "ip_address": "251.166.224.119"
    }
  ],
  "cyberchimps.com": [
    {
      "first_name": "Bonnie",
      "last_name": "Ortiz",
      "email": "bortiz1@cyberchimps.com",
      "gender": "Female",
      "ip_address": "197.54.209.129"
    }
  ],
  ...
```

## What would I have done differently?

Since we're talking about grouping customers by their email's domain, which would often be represented as `map[string][]*Customer`, it didn't make much sense to make the process concurrent, there wouldn't be in major performance gains since sharing a map among a pool of workers would require to make it thread-safe (ie: via mutex), which would turn out to be the by the main bottleneck. 

In a real project, where there're millions of customers and different requirements/features, I would have developed this in a more distributed way. For example: reading the file in chunks, having a pipeline for channels, maybe allowing for pre-sorted data so grouping can be more efficient, etc.