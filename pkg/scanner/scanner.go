package scanner

type Scanner interface {
	ScanTerminal(format string, value ...any) error
	ScanLine(line, format string, value ...any) error
}

var (
	_defaultScanner         = &defaultScanner{}
	_scanner        Scanner = _defaultScanner
)

func ScanTerminal(format string, value ...any) error {
	return _scanner.ScanTerminal(format, value...)
}

func ScanLine(line, format string, value ...any) error {
	return _scanner.ScanLine(line, format, value...)
}
