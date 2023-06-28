package main

import (
	"fmt"
	"logging/testdata"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func newLogger(useJson bool) *logrus.Logger {
	log := logrus.New()

	if useJson {
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
			// FieldMap is supposed to work for JSONFormatter and TextFormatter
			// but the compiler is complaining that it doesn't recognize it FieldKeyTime and FieldKeyFunc
			// FieldMap: logrus.FieldMap{
			// 	FieldKeyTime: "time",
			// 	FieldKeyFunc: "caller",
			// },
			PrettyPrint: true, // indents json
			// other options in docs: https://pkg.go.dev/github.com/sirupsen/logrus?utm_source=godoc#JSONFormatter
		})
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: time.RFC3339,
			// other options in docs: https://pkg.go.dev/github.com/sirupsen/logrus?utm_source=godoc#TextFormatter
		})
	}

	log.SetReportCaller(true)

	// default is stderr, which is fine for my test case, but I wanted to show how to set the io.Writer
	log.SetOutput(os.Stdout)

	// log level names end in Level
	log.SetLevel(logrus.InfoLevel)

	return log
}

func main() {
	log := newLogger(true)

	// Set standard fields by changing the logger to include them.
	// ctxLog is not a logrus.Logger, but acts like one.
	ctxLog := log.WithFields(logrus.Fields{
		"moduleName": "loglogrus",
		"exampleInt": 42,
	})

	ctxLog.Info(testdata.TStruct)

	ctxLog.WithField("error", testdata.TStruct.OriginalError).
		Error("Log an error ")

	ctxLog.WithField("data", fmt.Sprintf("%+v", testdata.TStruct)).
		Info("Log a complex data structure")

	ctxLog.WithFields(logrus.Fields{
		"map":  testdata.TStruct.Amap,
		"Code": testdata.TStruct.Code,
	}).
		Warn(fmt.Sprintf("Log format string %T %d %v", testdata.TStruct.Amap, testdata.TStruct.LineNumber, testdata.TStruct.Code))

	ctxLog.Debug("Debug log should not appear with level set to Info.")
}
