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

**COMMIT:** TEST: show how "stop all" and "continue" functions above behave
