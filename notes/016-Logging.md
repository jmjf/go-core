# Logging in Go

## Background

When building applications, especially backend, I want to be able to write logs so I can understand and fix problems.

I want to support:

* Log levels like fatal, error, info, debug, trace so I can control the amount of detail logged based on need. (info is normal, turn on debug or trace when diagnosing problems)
* Structured JSON logs so data logged is easier to interpret (different values identified and can be machine parsed for analysis)
* Easily switchable output format
  * Human readable console output in development (formatted JSON) so I can read the logged data easily
  * Machine readable console output in production (linear JSON) so I can send logs to tools like Splunk efficiently
  * All based on JSON

## Options

Go includes a `log` package, but it is weak as a structured logging tool.

Go 1.21 will include `log/slog`, a good, baseline structured logging tool, but I don't want to wait until it's ready.

The experimental/pre-release version of `log/slog` is available as `golang.org/x/exp/slog`.

For third party options, I looked at all the options listed at [Awesome Go](https://awesome-go.com/logging). I eliminated options with low engagement (stars, forks), that were archived, or otherwise seemed not popular, on point or well supported. That got me down the following options.

* [logrus](https://github.com/sirupsen/logrus)
* [zap](https://github.com/uber-go/zap)
* [zerolog](https://github.com/rs/zerolog)

(I'd like to try `apex/log` because TJ Holowaychuk built it, but the last commit is over a year old, so not an option.)

I'll test the experimental `slog`, `logrus`, `zap`, and `zerolog`. Besides the main criteria defined above, my concerns are about how reasonable the logging code and output looks.

## How to test

I expect most of my logs wil be driven on Golang errors, `structs`, or both, so I want to test logging both. I think primitive data type logs will be rare, but I'll test them too. I want to understand Golang's context concept and how it might apply to logging. (Can a context carry data for errors?)

The basic test program is:

* Create a logger that logs to stdout
* Log an error
* Log a struct
* Log a struct with nested a nested struct, array, map
* Log a message string with formatting (embedding values) plus specific values in fields
