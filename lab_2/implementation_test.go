package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixToPostfix_CorrectInput(t *testing.T) {
	res, err := PrefixToPostfix("+ 5 * - 4 2 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "4 2 - 3 * 5 +", res)
	}
}

func TestPrefixToPostfix_CorrectInputBigNumbers(t *testing.T) {
	res, err := PrefixToPostfix("+ 51 * - 43 25 36")
	if assert.Nil(t, err) {
		assert.Equal(t, "43 25 - 36 * 51 +", res)
	}
}

func TestPrefixToPostfix_BadChar(t *testing.T) {
	res, err := PrefixToPostfix("aaaa")
	if err != nil {
		assert.Equal(t, "bad char", err.Error())
		assert.Equal(t, "", res)
	} else {
		t.Errorf("Unexpected result")
	}
}

func TestPrefixToPostfix_CorrectInputStartSpaces(t *testing.T) {
	res, err := PrefixToPostfix("      + 5 * - 4 2 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "4 2 - 3 * 5 +", res)
	}
}

func TestPrefixToPostfix_CorrectInputEndSpaces(t *testing.T) {
	res, err := PrefixToPostfix("+ 5 * - 4 2 3           ")
	if assert.Nil(t, err) {
		assert.Equal(t, "4 2 - 3 * 5 +", res)
	}
}

func TestPrefixToPostfix_CorrectInputInnerSpaces(t *testing.T) {
	res, err := PrefixToPostfix("+ 5 * -         4 2 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "4 2 - 3 * 5 +", res)
	}
}

func TestPrefixToPostfix_InvExprNoOper(t *testing.T) {
	res, err := PrefixToPostfix("451")
	if err != nil {
		assert.Equal(t, "invalid expression", err.Error())
		assert.Equal(t, "", res)
	} else {
		t.Errorf("Unexpected result")
	}
}

func TestPrefixToPostfixInvExpr2Symb(t *testing.T) {
	res, err := PrefixToPostfix("-5")
	if err != nil {
		assert.Equal(t, "invalid expression", err.Error())
		assert.Equal(t, "", res)
	} else {
		t.Errorf("Unexpected result")
	}
}

func TestPrefixToPostfix_InvExprNoOper1(t *testing.T) {
	res, err := PrefixToPostfix("4-34")
	if err != nil {
		assert.Equal(t, "invalid expression", err.Error())
		assert.Equal(t, "", res)
	} else {
		t.Errorf("Unexpected result")
	}
}

func ExamplePrefixToPostfix() {
	res, _ := PrefixToPostfix("+ 2 2")
	fmt.Println(res)
}
