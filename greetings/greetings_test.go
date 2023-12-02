package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "yonghu"
	want := regexp.MustCompile(`\b` + name + `\b`)

	message, err := Hello(name)

	if err != nil || !want.MatchString(message) {
		t.Fatalf("Hello(yonghu) = %q, %v, want match for %#q", message, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
