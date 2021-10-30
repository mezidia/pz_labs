package lab2

import (
	"bytes"
	lab2 "github.com/mezidia/pz_labs/tree/main/lab_2"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestComputeHandler_CorrectInput(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	handler := lab2.ComputeHandler{
		In:  bytes.NewBufferString("+ 5 * - 4 2 3"),
		Out: w,
	}
	err := handler.Compute()
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	if assert.Nil(t, err) {
		assert.Equal(t, "postfix: 4 2 - 3 * 5 +\n", string(out[:]))
	} else {
		t.Errorf("Unexpected result")
	}
}
func TestComputeHandler_BadChar(t *testing.T) {
	handler := lab2.ComputeHandler{
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
	handler := lab2.ComputeHandler{
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
