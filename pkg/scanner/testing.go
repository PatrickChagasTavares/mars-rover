package scanner

import (
	"fmt"
	"testing"
)

type TestScan struct {
	Value []any
	Err   error
}

func (ts TestScan) ScanTerminal(format string, args ...any) error {
	if ts.Err != nil {
		return ts.Err
	}

	for i := 0; i < len(args); i++ {
		switch v := args[i].(type) {
		case *string:
			*v = ts.Value[i].(string)
		case *int:
			*v = ts.Value[i].(int)
		case *interface{}:
			*v = ts.Value[i].(interface{})
		default:
			return fmt.Errorf("faile decode")
		}
	}
	ts.Value = ts.Value[len(args):]

	return nil
}

func (ts TestScan) ScanLine(_, format string, args ...any) error {
	if ts.Err != nil {
		return ts.Err
	}

	for i := 0; i < len(args); i++ {
		switch v := args[i].(type) {
		case *string:
			*v = ts.Value[i].(string)
		case *int:
			*v = ts.Value[i].(int)
		case *interface{}:
			*v = ts.Value[i].(interface{})
		default:
			return fmt.Errorf("faile decode")
		}
	}
	ts.Value = ts.Value[len(args):]

	return nil
}

func SetTestMode(t *testing.T, scanner Scanner) {
	_scanner = scanner
	t.Cleanup(func() {
		_scanner = _defaultScanner
	})
}
