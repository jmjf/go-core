package main

import (
	"fmt"
	"logging/testdata"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func newLogger(useJson bool) *zerolog.Logger {
	zerolog.TimeFieldFormat = time.RFC3339

	var log zerolog.Logger

	log = zerolog.New(os.Stdout).
		Level(zerolog.InfoLevel).
		With().
		Str("moduleName", "logzerolog").
		Int("exampleInt", 42).
		Timestamp(). // default format is RFC3339
		Caller().
		Logger()

	// default is JSON logging; override with console output for non-JSON
	if !useJson {
		log = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	return &log
}

func main() {
	log := newLogger(true)

	log.Info().
		Interface("data", testdata.TStruct).
		Send() // no message

	log.Error().
		Err(testdata.TStruct.OriginalError).
		Msg("Log an error ")

	log.Info().
		Interface("data", testdata.TStruct).
		Msg("Log a complex data structure")

	log.Warn().
		Interface("map", testdata.TStruct.Amap).
		Str("Code", testdata.TStruct.Code).
		Msg(fmt.Sprintf("Log format string %T %d %v", testdata.TStruct.Amap, testdata.TStruct.LineNumber, testdata.TStruct.Code))

	log.Debug().Msg("Debug log should not appear with level set to Info.")
}
