package testdata

import "fmt"

type TestStruct struct {
	FileName      string `json:"fileName"`
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
	Name  string    `json:"name"`
	Stuff StuffType `json:"stuff"`
	Arry  []int     `json:"arry"`
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
