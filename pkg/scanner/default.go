package scanner

import "fmt"

type defaultScanner struct{}

func (s defaultScanner) ScanTerminal(format string, value ...any) error {
	_, err := fmt.Scanf(format, value)
	return err
}

func (s defaultScanner) ScanLine(line, format string, value ...any) error {
	_, err := fmt.Sscanf(line, format, value...)
	return err
}
