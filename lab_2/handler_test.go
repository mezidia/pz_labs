package lab2

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeHandler_CorrectInput(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	handler := ComputeHandler{
		In:  bytes.NewBufferString("+ 5 * - 4 2 3"),
		Out: w,
	}
	err := handler.Compute()
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	if assert.Nil(t, err) {
		assert.Equal(t, "[+ 5 * - 4 2 3]\npostfix: 4 2 - 3 * 5 +\n", string(out[:]))
	} else {
		t.Errorf("Unexpected result")
	}
}
func TestComputeHandler_BadChar(t *testing.T) {
	handler := ComputeHandler{
		In:  bytes.NewBufferString("aaa"),
		Out: os.Stdout,
	}
	err := handler.Compute()
	if err != nil {
		assert.Equal(t, "bad char", err.Error())
	} else {
		t.Errorf("Unexpected result")
	}
}
func TestComputeHandler_WrongExpression(t *testing.T) {
	handler := ComputeHandler{
		In:  bytes.NewBufferString("1"),
		Out: os.Stdout,
	}
	err := handler.Compute()
	if err != nil {
		assert.Equal(t, "invalid expression", err.Error())
	} else {
		t.Errorf("Unexpected result")
	}
}
