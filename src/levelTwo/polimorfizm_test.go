package main

import (
	"strings"
	"bytes"
	"testing"
	"bufio"
)

var testData = `1
2
3
4
5
6
`

var checkData = `1
2
3
4
5
6
`

func TestOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testData))
	out := new(bytes.Buffer)
	err := uniq(in, out)
	if err != nil {
		t.Error("test failed")
	}
	resultData := out.String()
	if resultData != checkData {
		t.Errorf("result not valid %v %v", checkData, resultData)
	}
}
