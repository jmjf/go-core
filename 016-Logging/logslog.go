package main

import (
	"fmt"
	"os"

	"golang.org/x/exp/slog"
)

type TestStruct struct {
	FileName      string
	FunctionName  string
	LineNumber    int
	Message       string
	Code          string
	ErrorData     any
	CanRetry      bool
	OriginalError error
	Amap          map[string]int
}

type StuffType struct {
	Line1 string
	Line2 string
}

type TestErrData struct {
	Name  string
	Stuff StuffType
	Arry  []int
}

var TEData = TestErrData{
	Name: "Joe",
	Stuff: StuffType{
		Line1: "123 Elm St",
		Line2: "Apt 987",
	},
	Arry: []int{2, 42, 32, 1},
}

var TStruct = TestStruct{
	FileName:      "main.go",
	FunctionName:  "main",
	LineNumber:    32,
	Message:       "test log message",
	Code:          "test",
	ErrorData:     TEData,
	CanRetry:      false,
	OriginalError: fmt.Errorf("original err %w", fmt.Errorf("wrapped error")),
	Amap:          map[string]int{"key1": 3, "key2": 1, "key32": 98232},
}

func NewSlogLogger(useJson bool) *slog.Logger {
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
	logger := NewSlogLogger(true)

	logger.Error("Log an error", "error", TStruct.OriginalError)

	logger.Info("Log a complex data structure", "data", TStruct)

	logger.Warn(fmt.Sprintf("Log format string %T %d %v", TStruct.Amap, TStruct.LineNumber, TStruct.Code), "map", TStruct.Amap, "Code", TStruct.Code)
}
