package gotest_test

import (
	"testing"

	"github.com/krhubert/gotest/go/ostest"
)

func ExampleOsOpen(t *testing.T) {
	ostest.Open(t, "null")
}
