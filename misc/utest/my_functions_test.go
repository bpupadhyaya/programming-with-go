package utest

import (
    "testing"
    "utest"
)

func TestAdd(t *testing.T) {
    if utest.Add(5, 2) != 7 {
        t.Error(" 5 + 2 is not 7")
    }
}

func TestSubtract(t *testing.T) {
    if utest.Subtract(5, 2) != 3 {
        t.Error("5 -2 is not 3")
    }
}

func TestMultiply(t *testing.T) {
    if utest.Multiply(5, 2) != 10 {
        t.Error("5 * 2 is not 10")
    }
}

func TestDivide(t *testing.T) {
    if utest.Divide(5.0, 2.0) != 2.5 {
        t.Error("5 divided by 2 is not 2.5")
    }
}