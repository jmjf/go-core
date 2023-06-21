## Deeper dive into testing

## Side note

I've looked at the community testing libraries a bit. I'm personally more comfortable with Ginkgo's BDD style. Testify seems to be closer to the native test runner's model. GoConvey's syntax is a bit mind bending and I'm not really interested in a browser-based test viewer. Httpexpect reminds me of some HTTP test helpers from JS-land and might be interesting.

I want to work with the native test runner for now and see how it might compare to Testify or Ginkgo and decide which I prefer.

Also, I spotted a couple of other interesting things related to testing.

* [GoTestSum](https://github.com/gotestyourself/gotestsum) -- uses `go test` but gets JSON output and formats it nicely. Output looks like it may be easier to read.
* [Mockery](https://vektra.github.io/mockery/) -- builds mock implementations for interfaces. This could be handy for some of the tests I want to write. Seems related to Testify.

## Test conventions and strategies

Golang has some conventions the tools expect code to follow.

* Tests add `_test` to the filename of the file they test; identifies tests and excludes from binaries
* Test names begin with `TestX` where `X` is the thing being tested
* Tests accept `t *testing.T`, which gives access to the test infrastructure; `t` isn't magic, can be anything
* For black box tests, test packages end in `_test`; ex: `main` -> `main_test`
* For white box tests, tests are in the same package so they have access to private code

Prefer black box tests so you're less likely to test implementation details. Consider the code below. If I want to test `privateFunc` directly, I must use a whie box tests. But, if I black box test `PublicFunc` with values 9.9, 10.0, and 10.1 and confirm that PublicFunc returns -1/err, -1/err, and 20.2/err, I don't need a separate test for `publicFunc`.

```golang
type MyObject struct {}

type MyInterface interface {
   privateFunc(i int) (int, err)
   PublicFunc(x float64) (int, err)
}

func (o MyObject) privateFunc(i int) (int, err) {
   if i > 100 {
      return 2, nil
   }
   return -1, errors.new("error from privateFunc()")
}

func (o MyObject) PublicFunc(x float64) (int, err) {
   i := int(x * 10)
   res, err := privateFunc(i)
   if err != nil {
      return res, err
   }
   return int(x * i)
}
```

Integration and end-to-end tests are usually best done black box and tests kept in a directory separate from the package(s) being tested.

When testing, we can report failures to the test engine that are noted/reported, but allow tests to continue, and failures that stop testing immediately. We might use the latter in an integration test that requires a database connection if we can't connect to the database (and expect it to succeed).

Stop all testing (stops running the test function; runs test functions after):

* `t.FailNow()` -- just stop
* `t.Fatal(args ...interface{})` -- stop and print the args
* `t.Fatalf(format string, args ...interface{})` -- stop and print a formatted string

Continue testing (runs code in the test function after the fail):

* `t.Fail()` -- just fail the test
* `t.Error(args ...interface{})` -- fail the test and print args
* `t.Errorf(format string, args ...interface{})` -- fail the test and print a formatted string

## Testing and test examples

Run tests with `go test`. The command has a lot of options and docs available with `go help test` and `go help testflags`. Among the interesting options, `go test -run <testpattern>` runs only tests that match `<testpattern>`. For example, `go test -run Welc` will run any test function whose name includes `Welc`. If it finds no matching tests, it ends with a PASS result and a warning that no tests ran. The comparison is case sensitive.

**COMMIT:** TEST: show how "stop all" and "continue" functions above behave

**COMMIT:** TEST: show how white box testing works

If I change the package name to `greeter_test`, I'll get an error when calling methods from `greeter`. I need to import `moreTesting/greeter` and prefix calls with `greeter` (`greeter.Welcome()`). I can't call `greeter.buhbye()` because it isn't exported (is private) and I'm in a different package.

**COMMIT:** TEST: show how black box testing works

## Coverage

***OPINION***

Test coverage matters, but there is no magic number. More coverage is desirable until it becomes overkill and slows testing and builds for little or no extra benefit. If possible, testing outer level methods is better. For example see `privateFunc` and `PublicFunc` discussion earlier in this document.

As another example, when writing an application where a use case (DDD/clean architecture concept) calls a domain object for data handling, repo to manage persistent data, etc., I can set up use case tests and cover the repo and domain object without needing separate tests for them. Assume I mock the database connection and set up the mock to return a valid row when it gets a SELECT query. I can test my use case that calls the repo's get method and don't need separate tests for the repo. Another test can return an error, which my use case should recognize and handle appropriately so needs to be tested too.

Those two tests give me confidence my use case and repo handle found and error correctly. Adding tests to cover the repo methods separately is more test code to write, more test code to execute, and more test code to maintain.

So, the goal is not to write a test for every function. The goal is to write tests that cover the important parts of the code so integration and end to end tests are less about the code you're writing and more about ensuring the external services (database, message bus, logging destination, etc.) behave as expected, performance is decent, service load is tolerable, etc.

***/OPINION***

Back to test coverage reporting in golang. The magic command is `go test -cover`, which runs tests and gives a coverage percent.

`go test -coverprofile <coverfile>` writes details to a file, then run `go tool cover -func <coverfile>` to see a summary of which code is and isn't covered. Example output below.

```
moreTesting/greeter/greeter.go:6:       Welcome         100.0%
moreTesting/greeter/greeter.go:11:      buhbye          0.0%
total:                                  (statements)    50.0%
```

`go tool -html <coverfile>` gives an HTML view in browser that color codes text to identify code that is covered and not covered. I'm working in a dev container that doesn't have a browser installed or a display. So `go tool -html <coverfile> -o <htmlfile>` writes to a file, which I can open in the host browser. See `./013-MoreTesting/greeter/cover.html` for an example.

`go test -coverprofile callcount.tst -covermode count` gives output that, when converted to HTML, color codes relative coverage level (shades of green). As far as I can tell, the only difference vs. normal `-coverprofile` (set mode) is that one file begins with `mode: set` and the other begins with `mode: count`.

In the `013-MoreTesting` directory, `cover.tst` and `cover.html` are based on commenting the condition in `Welcome` and `cover2.tst` and `cover2.html` are based on uncommenting it. The second set shows that the output file tracks specific lines of code. So, it knows that the `if` was executed but the `fmt.Println()` inside the `if` was not.

When running `go test -cover -v`, the output doesn't seem to consider what part of a function is covered. It reports 50% coverage whether the `if` is commented or not. But `go tool cover -func cover2.tst` shows percent tested for each function (66.7% for `Welcome` because one of three executable lines is not executed). The total coverage is still 50%, though.

I think the HTML outputs would get overwhelming for a large body of code, so is most useful for one or two files at a time. Also, the fact that it's color coded makes it less useful to colorblind people. The different shades of green may be difficult to distinguish for some people. Maybe more interesting would be to filter the standard line output to identify code that is only partially covered.

The native test runner has the main features I might want for coverage reporting. Some of the results may be misleading and some of the output modes may be difficult to use in larger projects.

**COMMIT:** TEST: try different test coverage features in the standard test runner
