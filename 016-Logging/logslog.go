package main

import (
	"fmt"
	"os"

	"logging/testdata"

	"golang.org/x/exp/slog"
)

func newLogger(useJson bool) *slog.Logger {
	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
		//This function removes the source attribute
		// ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
		// 	if a.Key == slog.SourceKey {
		// 		return slog.Attr{}
		// 	}
		// 	return a
		// },
	}

	var handler slog.Handler

	if useJson {
		handler = slog.NewJSONHandler(os.Stdout, &opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, &opts)
	}

	return slog.New(handler.WithAttrs([]slog.Attr{slog.String("moduleName", "logslog"), slog.Int("exampleInt", 42)}))
}

func main() {
	// pass true for JSON, false for text
	logger := newLogger(true)

	logger.Error("Log an error", "error", testdata.TStruct.OriginalError)

	logger.Info("Log a complex data structure", "data", testdata.TStruct)

	logger.Warn(fmt.Sprintf("Log format string %T %d %v", testdata.TStruct.Amap, testdata.TStruct.LineNumber, testdata.TStruct.Code),
		"map", testdata.TStruct.Amap, "Code", testdata.TStruct.Code)
}
