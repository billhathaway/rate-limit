package main

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestTiming(t *testing.T) {
	tenLineInput := `1
2
3
4
5
6
7
8
9
10
`

	buf := bytes.Buffer{}
	start := time.Now()
	err := loopLines(10, strings.NewReader(tenLineInput), &buf)
	duration := time.Since(start)
	if err != nil {
		t.Error(err)
	}
	if duration.Seconds() < .9 || duration.Seconds() > 1.1 {
		t.Fail()
	}
	t.Logf("loopLines took %.3f seconds\n", duration.Seconds())
	if buf.String() != tenLineInput {
		t.Fail()
		t.Log("Output did not match input")
	} else {
		t.Log("Output matched input")
	}
}
