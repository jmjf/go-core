package main

import (
	"fmt"
	"logging/testdata"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newLogger(useJson bool) *zap.Logger {

	// A custom config might look like this
	// zapConfig := zap.Config{
	// 	Level: zap.InfoLevel,
	// 	DisableCaller: false,
	// 	DisableStackTrace: false,
	// 	Encoding: "json",
	// 	OutputPaths: []string{"stdout"},
	// 	InitialFields: map[string]interface{}{
	// 		"moduleName": "logzap",
	// 		"exampleInt": 42,
	// 	},
	// }

	atLevel := zap.NewAtomicLevel()
	atLevel.SetLevel(zapcore.InfoLevel)

	encConfig := zap.NewProductionEncoderConfig()
	encConfig.EncodeTime = zapcore.ISO8601TimeEncoder // or RFC3339TimeEncoder
	encConfig.TimeKey = "time"

	var encoder zapcore.Encoder
	if useJson {
		encoder = zapcore.NewJSONEncoder(encConfig)
	} else {
		encConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encConfig)
	}

	logger := zap.New(
		zapcore.NewCore(
			encoder,
			zapcore.Lock(os.Stdout), // makes stdout safe for concurrent use by wrapping in a mutex
			atLevel),
		zap.AddCaller(),
	).With( // default fields
		zap.String("moduleName", "logzap"),
		zap.Int("exampleInt", 42),
		zap.Int64("pid", int64(os.Getpid())),
	)

	return logger
}

func newSugar(useJson bool) *zap.SugaredLogger {
	return newLogger(useJson).Sugar()
}

// Demonstrates using zap with a plain logger
func testLogger(useJson bool) {
	log := newLogger(useJson)
	defer log.Sync()

	// Uses reflection to build an object field in the message
	// Alternatively, Object(key, zapcore.ObjectMarshaler)
	// For ObjectMarshaler, attach `MarshalLogObject(enc zapcore.ObjectEncoder) error`
	// to the struct and use `enc.Add*(key, value)` methods to avoid reflection overhead.
	// https://pkg.go.dev/go.uber.org/zap#Object
	log.Info("", zap.Reflect("data", testdata.TStruct))

	log.Error("Log an error ", zap.Error(testdata.TStruct.OriginalError))

	log.Info("Log a complex data structure", zap.Reflect("data", testdata.TStruct))

	log.Warn(fmt.Sprintf("Log format string %T %d %v", testdata.TStruct.Amap, testdata.TStruct.LineNumber, testdata.TStruct.Code),
		zap.Reflect("map", testdata.TStruct.Amap),
		zap.String("Code", testdata.TStruct.Code),
	)

	log.Debug("Debug log should not appear with level set to Info.")
}

func testSugar(useJson bool) {
	sugar := newSugar(useJson)
	defer sugar.Sync()

	sugar.Infow("",
		"data", testdata.TStruct,
	)

	sugar.Errorw("Log an error ",
		"error", testdata.TStruct.OriginalError,
	)

	sugar.Infow("Log a complex data structure",
		"data", testdata.TStruct,
	)

	sugar.Warnf("Log format string with Warnf %T %d %v", testdata.TStruct.Amap, testdata.TStruct.LineNumber, testdata.TStruct.Code)

	sugar.Warnw(fmt.Sprintf("Log format string with Warnw %T %d %v", testdata.TStruct.Amap, testdata.TStruct.LineNumber, testdata.TStruct.Code),
		"map", testdata.TStruct.Amap,
		"Code", testdata.TStruct.Code,
	)

	sugar.Debug("Debug log should not appear with level set to Info.")
}

func main() {
	useJson := true
	fmt.Println("***** Plain Logger *****\n")
	testLogger(useJson)

	fmt.Println("\n***** Sugared Logger *****\n")
	testSugar(useJson)
}
