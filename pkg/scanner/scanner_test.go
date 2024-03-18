package scanner_test

import (
	"errors"
	"testing"

	"github.com/patrickchagastavares/rover/pkg/scanner"
	"github.com/stretchr/testify/assert"
)

func TestScanTerminal(t *testing.T) {
	tests := map[string]struct {
		format string
		input  string
		value  any
		err    error
	}{
		"string":  {format: "%s", input: "Hello", value: new(string), err: nil},
		"integer": {format: "%d", input: "42", value: new(int), err: nil},
		"float":   {format: "%f", input: "3.14", value: new(float64), err: nil},
		"error":   {format: "%d", input: "not_an_integer", value: new(int), err: errors.New("expected integer")},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			scanner.SetTestMode(t, &scanner.TestScan{
				Value: []any{test.value},
				Err:   test.err,
			})

			err := scanner.ScanTerminal(test.format, &test.value)
			assert.Equal(t, test.err, err)
		})
	}
}
