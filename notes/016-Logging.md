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

I'd like to try `apex/log` because TJ Holowaychuk built it, but the last commit is from 2020, Apex's SSL cert has been revoked (site dead), and TJ hasn't done much on GitHub since 2020 (and by TJ's historical standards, 2020 was a slow year).

`logrus` is in maintenance mode, meaning no new features, only security and bug fixes. (It points to the two Z-loggers above and `apex/log`, was put into maintenance in February 2020.) Even so, it's more forked and starred than either of the Z-loggers, so gets a look.

I'll test the experimental `slog`, `logrus`, `zap`, and `zerolog`. Besides the main criteria defined above, my concerns are about how reasonable the logging code and output looks.

## How to test

I expect most of my logs wil be driven on Golang errors, `structs`, or both, so I want to test logging both. I think primitive data type logs will be rare, but I'll test them too. I want to understand Golang's context concept and how it might apply to logging. (Can a context carry data for errors?)

The basic test program is:

* Create a logger that logs to stdout
* Log an error
* Log a struct with nested a nested struct, array, map
* Log a message string with formatting (embedding values) plus specific values in fields

I'll work in `016-Logging`. `main.go` will import from packages for each logger `logslog`, `loglogrus`, `logzap`, and `logzerolog`, which have specifics for each logger.

## slog

I used information from:

* [Example from slog docs](https://pkg.go.dev/golang.org/x/exp/slog#example-package-Wrapping)
* [Logging in Go with slog](https://thedevelopercafe.com/articles/logging-in-go-with-slog-a7bb489755c2)

I need to `go get golang.org/x/exp/slog` to pull the experimental package. (After 1.21 releases, I should be able to import from `log/slog` with no `go get`.)

In `logslog.go`, I'll build an example using `slog`.

Printing a formatted line (third message) uses `fmt.Sprintf()` to format the string and key/value parameter pairs to include the values as logged fields (map, Code).

Text output

```
time=2023-06-27T02:06:12.043Z level=ERROR source=/workspace/016-Logging/logslog.go:76 msg="Log an error" moduleName=logslog exampleInt=42 error="original err wrapped error"
time=2023-06-27T02:06:12.044Z level=INFO source=/workspace/016-Logging/logslog.go:78 msg="Log a complex data structure" moduleName=logslog exampleInt=42 data="{FileName:main.go FunctionName:main LineNumber:32 Message:test log message Code:test ErrorData:{Name:Joe Stuff:{Line1:123 Elm St Line2:Apt 987} Arry:[2 42 32 1]} CanRetry:false OriginalError:original err wrapped error Amap:map[key1:3 key2:1 key32:98232]}"
time=2023-06-27T02:06:12.044Z level=WARN source=/workspace/016-Logging/logslog.go:80 msg="Log format string map[string]int 32 test" moduleName=logslog exampleInt=42 map="map[key1:3 key2:1 key32:98232]" Code=test
```

JSON output

```json
{"time":"2023-06-27T02:05:39.80322462Z","level":"ERROR","source":{"function":"main.main","file":"/workspace/016-Logging/logslog.go","line":76},"msg":"Log an error","moduleName":"logslog","exampleInt":42,"error":"original err wrapped error"}
{"time":"2023-06-27T02:05:39.803321469Z","level":"INFO","source":{"function":"main.main","file":"/workspace/016-Logging/logslog.go","line":78},"msg":"Log a complex data structure","moduleName":"logslog","exampleInt":42,"data":{"FileName":"main.go","FunctionName":"main","LineNumber":32,"Message":"test log message","Code":"test","ErrorData":{"Name":"Joe","Stuff":{"Line1":"123 Elm St","Line2":"Apt 987"},"Arry":[2,42,32,1]},"CanRetry":false,"OriginalError":{},"Amap":{"key1":3,"key2":1,"key32":98232}}}
{"time":"2023-06-27T02:05:39.803397411Z","level":"WARN","source":{"function":"main.main","file":"/workspace/016-Logging/logslog.go","line":80},"msg":"Log format string map[string]int 32 test","moduleName":"logslog","exampleInt":42,"map":{"key1":3,"key2":1,"key32":98232},"Code":"test"}
```

I note that, by default, `slog` includes `source`, which refers to the log line. If logging happens removed from the source of data logged, `source` may be less valuable. To remove it, use `ReplaceAttr` and the example function in the code. Note that `ReplaceAttr` can remove or change other attributes. (An `Attr` has a `Key` and a `Value`)

```json
{"time":"2023-06-27T02:12:55.60012797Z","level":"ERROR","msg":"Log an error","moduleName":"logslog","exampleInt":42,"error":"original err wrapped error"}
{"time":"2023-06-27T02:12:55.600225626Z","level":"INFO","msg":"Log a complex data structure","moduleName":"logslog","exampleInt":42,"data":{"FileName":"main.go","FunctionName":"main","LineNumber":32,"Message":"test log message","Code":"test","ErrorData":{"Name":"Joe","Stuff":{"Line1":"123 Elm St","Line2":"Apt 987"},"Arry":[2,42,32,1]},"CanRetry":false,"OriginalError":{},"Amap":{"key1":3,"key2":1,"key32":98232}}}
{"time":"2023-06-27T02:12:55.600317891Z","level":"WARN","msg":"Log format string map[string]int 32 test","moduleName":"logslog","exampleInt":42,"map":{"key1":3,"key2":1,"key32":98232},"Code":"test"}
```

**COMMIT:** FEAT: demo slog logger

Preparing to test logrus, I moved the test data into a package (`testdata`) and changed `logslog.go` to use it.

**COMMIT:** REFACTOR: move test data into a package so it's easier to share

## logrus