package main_test

import (
	"bytes"
	"io"
	"testing"

	. "github.com/gugahoi/go-junit-md"
)

func TestFailureBlock(t *testing.T) {
	testCases := []struct {
		desc   string
		writer io.Writer
	}{
		{
			desc: "good writer", writer: bytes.NewBufferString(""),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// arrange

			// assign
			err := FailureBlock(tC.writer, "name", "class", "body")

			// assert
			if err != nil {
				t.Fatalf("Expected err to be nil, got %v", err)
			}
		})
	}
}
