package main

import (
	"bytes"
	"strings"
	"testing"
)

// Fill stdin with enough characters to cause reader to fill its buffer
func makeLongInputAndError() bytes.Buffer {
	var buffer bytes.Buffer

	// Scanner has a default max token size of 64 * 1024
	buffer.Write([]byte("x"))
	for i := 0; i < 64*64*65*1024; i++ {
		buffer.Write([]byte("a"))
	}
	buffer.Write([]byte("error\n"))

	return buffer
}

func makeLongInputAndNoError() bytes.Buffer {
	var buffer bytes.Buffer

	buffer.Write([]byte("x"))
	for i := 0; i < 65*1024; i++ {
		buffer.Write([]byte("a"))
	}

	return buffer
}

// Both tests check the first character to make sure the
// full error string is being returned.
func TestLongInputError(t *testing.T) {
	c := make(chan string)

	msgBuffer := makeLongInputAndError()
	msgReader := bytes.NewReader(msgBuffer.Bytes())
	go ScanForErrors(c, msgReader)

	res := <-c

	if res == "" {
		t.Fail()
	}

	if !strings.Contains(res, "error") || res[0] != 'x' {
		t.Fail()
	}
}

func TestLongInputNoError(t *testing.T) {
	c := make(chan string)

	msgBuffer := makeLongInputAndNoError()
	msgReader := bytes.NewReader(msgBuffer.Bytes())
	go ScanForErrors(c, msgReader)

	res := <-c

	if len(res) != 0 {
		t.Fail()
	}
}
