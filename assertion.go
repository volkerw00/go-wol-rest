package main

import (
	"strings"
	"testing"
)

var test *testing.T

func assertThatInt(actual int, matcher func(int) (bool, func())) {
	result, printer := matcher(actual)
	if !result {
		printer()
	}
}

func isInt(expected int) func(int) (bool, func()) {
	return func(actual int) (bool, func()) {
		if expected == actual {
			return true, func() {
			}
		}
		return false, func() {
			isErrorInt(actual, expected)
		}
	}
}

func isErrorInt(actual, expected int) {
	test.Errorf("expected %d but was %d", expected, actual)
}

func assertThatString(actual string, matcher func(string) (bool, func())) {
	result, printer := matcher(actual)
	if !result {
		printer()
	}
}

func containsString(expected string) func(string) (bool, func()) {
	return func(actual string) (bool, func()) {
		if strings.Index(actual, expected) > 0 {
			return true, func() {
			}
		}
		return false, func() {
			containsErrorString(actual, expected)
		}
	}
}

func containsErrorString(actual, expected string) {
	test.Errorf("expected body to contain '%s' but was '%s'", expected, actual)
}
