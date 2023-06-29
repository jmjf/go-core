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

[GitHub for logrus](https://github.com/sirupsen/logrus) -- has decent documentation.

`go get github.com/sirupsen/logrus`

I have a working version of it, but I see a few issues.

* The syntax for logging individual fields feels clunkier than `slog`'s syntax.
* No direct support for logging structs.
  * I used a cheat (`Sprintf %+v`), but the output isn't JSON. I'd need to convert it to a map (for `Fields`) or JSON (for string).

The JSON output indenter is nice for testing, but not something for production. JSON color coding would be a nice addition.

The text logs are terse, IMO, but color coding makes them more readable.

It's a tossup between `logrus` and `slog` right now.

**COMMIT:** FEAT: add logrus logger example

## zerolog

[GitHub](https://github.com/rs/zerolog)
[A Complete Guide to Logging with zerolog](https://betterstack.com/community/guides/logging/zerolog/)

`go get -u github.com/rs/zerolog/log` -- `-u` updates modules providing dependencies

Example JSON output. I added JSON tags for `fileName` and the members of `ErrorData` (but not members of `stuff`).

```json
{"level":"info","moduleName":"logzerolog","exampleInt":42,"data":{"fileName":"main.go","FunctionName":"main","LineNumber":32,"Message":"test log message","Code":"test","ErrorData":{"name":"Joe","stuff":{"Line1":"123 Elm St","Line2":"Apt 987"},"arry":[2,42,32,1]},"CanRetry":false,"OriginalError":{},"Amap":{"key1":3,"key2":1,"key32":98232}},"time":"2023-06-28T02:32:16Z","caller":"/workspace/016-Logging/logzerolog.go:40"}
{"level":"error","moduleName":"logzerolog","exampleInt":42,"error":"original err wrapped error","time":"2023-06-28T02:32:16Z","caller":"/workspace/016-Logging/logzerolog.go:44","message":"Log an error "}
{"level":"info","moduleName":"logzerolog","exampleInt":42,"data":{"fileName":"main.go","FunctionName":"main","LineNumber":32,"Message":"test log message","Code":"test","ErrorData":{"name":"Joe","stuff":{"Line1":"123 Elm St","Line2":"Apt 987"},"arry":[2,42,32,1]},"CanRetry":false,"OriginalError":{},"Amap":{"key1":3,"key2":1,"key32":98232}},"time":"2023-06-28T02:32:16Z","caller":"/workspace/016-Logging/logzerolog.go:48","message":"Log a complex data structure"}
{"level":"warn","moduleName":"logzerolog","exampleInt":42,"map":{"key1":3,"key2":1,"key32":98232},"Code":"test","time":"2023-06-28T02:32:16Z","caller":"/workspace/016-Logging/logzerolog.go:53","message":"Log format string map[string]int 32 test"}
```

The JSON output looks as good as `slog`'s JSON. `zerolog` does require a lot of function chaining to add values into the log. I can see that as more work or more explicit than `slog`. I think the syntax is a tossup.

Text output is okay. `zerolog` is JSON and CBOR focused, but text is comparable to `logrus`. On the console, it's color coded.

```
2:45AM INF logzerolog.go:39 > data={"Amap":{"key1":3,"key2":1,"key32":98232},"CanRetry":false,"Code":"test","ErrorData":{"arry":[2,42,32,1],"name":"Joe","stuff":{"Line1":"123 Elm St","Line2":"Apt 987"}},"FunctionName":"main","LineNumber":32,"Message":"test log message","OriginalError":{},"fileName":"main.go"} exampleInt=42 moduleName=logzerolog
2:45AM ERR logzerolog.go:43 > Log an error  error="original err wrapped error" exampleInt=42 moduleName=logzerolog
2:45AM INF logzerolog.go:47 > Log a complex data structure data={"Amap":{"key1":3,"key2":1,"key32":98232},"CanRetry":false,"Code":"test","ErrorData":{"arry":[2,42,32,1],"name":"Joe","stuff":{"Line1":"123 Elm St","Line2":"Apt 987"}},"FunctionName":"main","LineNumber":32,"Message":"test log message","OriginalError":{},"fileName":"main.go"} exampleInt=42 moduleName=logzerolog
2:45AM WRN logzerolog.go:52 > Log format string map[string]int 32 test Code=test exampleInt=42 map={"key1":3,"key2":1,"key32":98232} moduleName=logzerolog
```

I think `zerolog` and `slog` are tied in terms of syntax and output. I'll need to do a feature comparison to see if there's notable difference.

(Side note, I've been thinking, with a few tweaks to field names, I could probably feed the JSON logs from `slog` and `zerolog` into `pino-pretty` and get readable output.)

**COMMIT:** FEAT: add zerolog logger example

## zap

[GitHub](https://github.com/uber-go/zap)
[GoDocs - zap](https://pkg.go.dev/go.uber.org/zap)
[GoDocs - zapcore](https://pkg.go.dev/go.uber.org/zap@v1.24.0/zapcore)

`go get -u go.uber.org/zap`

`zap` supports two ways of logging. The "sugared" logger offers a syntax similar to `slog` for a small performance penalty. The plain logger is faster and uses syntax similar to `zerolog`, but without function chaining or `zerolog`'s separate `Msg()`.

```golang
// zerolog

 log.Error().
  Err(testdata.TStruct.OriginalError).
  Msg("Log an error ")

// zap
  log.Error("Log an error ", zap.Error(testdata.TStruct.OriginalError))
```

`zap` has three default configurations (Production, Development, Example) with a `Config` option for custom configuration. The `Config` type includes JSON and YAML tags to allow easy import from JSON or YAML data using Go's native marhalling functions.

* Production writes `InfoLevel` and above JSON logs to stderr.
* Development writes `DebugLevel` and above human-friendly logs to stderr.
* Example writes `DebugLevel` and above JSON logs to stderr.

I have a demo for both a plain logger and a sugared logger.

I like the fact that `zap` natively supports ISO-8601 time formats (RFC3339 with ms). I can support that with other loggers by defining a custom time format, so it isn't a big deal, but it's nice nonetheless.

The loggers include a `Sync` method, which flushes any buffered logs. Adding `defer logger.Sync()` after creating the logger in `main()` or similar seems to be a recommended practice. In my case, it's in the `testLogger` and `testSugar` functions because they create an use the logger and don't expose it.

Console logs aren't color coded, but aren't much less readable than the other loggers that support a non-JSON option.

JSON output sample:

```json
{"level":"info","time":"2023-06-29T01:42:49.230Z","caller":"016-Logging/logzap.go:87","msg":"","moduleName":"logzap","exampleInt":42,"data":{"fileName":"main.go","FunctionName":"main","LineNumber":32,"Message":"test log message","Code":"test","ErrorData":{"name":"Joe","stuff":{"Line1":"123 Elm St","Line2":"Apt 987"},"arry":[2,42,32,1]},"CanRetry":false,"OriginalError":{},"Amap":{"key1":3,"key2":1,"key32":98232}}}
{"level":"error","time":"2023-06-29T01:42:49.230Z","caller":"016-Logging/logzap.go:91","msg":"Log an error ","moduleName":"logzap","exampleInt":42,"error":"original err wrapped error"}
{"level":"info","time":"2023-06-29T01:42:49.230Z","caller":"016-Logging/logzap.go:95","msg":"Log a complex data structure","moduleName":"logzap","exampleInt":42,"data":{"fileName":"main.go","FunctionName":"main","LineNumber":32,"Message":"test log message","Code":"test","ErrorData":{"name":"Joe","stuff":{"Line1":"123 Elm St","Line2":"Apt 987"},"arry":[2,42,32,1]},"CanRetry":false,"OriginalError":{},"Amap":{"key1":3,"key2":1,"key32":98232}}}
{"level":"warn","time":"2023-06-29T01:42:49.230Z","caller":"016-Logging/logzap.go:99","msg":"Log format string with Warnf map[string]int 32 test","moduleName":"logzap","exampleInt":42}
{"level":"warn","time":"2023-06-29T01:42:49.230Z","caller":"016-Logging/logzap.go:101","msg":"Log format string with Warnw map[string]int 32 test","moduleName":"logzap","exampleInt":42,"map":{"key1":3,"key2":1,"key32":98232},"Code":"test"}
```

## Can I get pino-pretty to play nice with JSON logs?

In Node-land, I use `pino` as a logger and [`pino-pretty`](https://github.com/pinojs/pino-pretty) in demos and run-it tests to get readable logs. `pino-pretty` wants a JSON structured log and formats based on certain-named fields. Specifically, it likes `time`, `pid`, `level`, and `msg`.

By default, `pino` writes log level numbers, not strings, and `pino-pretty` interprets them. To get those from the Go loggers, I'd need a way to intercept the log line and add it. Let's see how far we can get anyway.

Pino log levels map as shown below.

```golang
 levelMap := map[string]int{
  "fatal": 60,
  "error": 50,
  "warn": 40,
  "info": 30,
  "debug": 20,
  "trace": 10,
 }
```

In `016-Logging` `npm init` to create `package.json`, then `npm install --save-dev pino-pretty`.

`go run logzap.go | npx pino-pretty` -- except for the lack of log level that `pino-pretty` understands, it works okay.

And if I remove my level rename and let it log as `level`, `pino-pretty` recognizes it and color codes it. And `slog` and `zerolog` work well with `pino-pretty` too. I'm seeing issues with `logrus`, though I'm not sure why yet. It already has a few strikes against it, so not worth solving.

With `pino-pretty` making output easier to read, I'm seeing that `slog`, `zerolog` and `zap` all seem to miss the `error` in `TStruct`. The print `"OriginalError": {},`, instead of `OriginalError:original err wrapped error`. `logrus` does log the embedded error correctly, though that doesn't overcome the other issues. So, when logging `error`s, I need to be sure to pull it out of any objects that wrap it. With `zap`, if I went the `Object` route instead of the `Reflect` route, I could do that. (Based on what I'm seeing, I'm fairly certain reflection is the issue.)

## Conclusion

* `logrus` -- no longer actively developed (but still maintained)
* `slog` -- part of the standard library; features are more limited than the next two
* `zerolog` -- viable option
* `zap` -- viable option

At the moment, I'm on the fence between the last three. There's little to recommend one above the other except `slog`'s more limited feature set.

**COMMIT:** FEAT: add zap logger example; test pino-pretty; summarize
