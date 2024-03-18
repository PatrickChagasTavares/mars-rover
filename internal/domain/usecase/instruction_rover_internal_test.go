package usecase

import (
	"errors"
	"fmt"
	"testing"

	"github.com/patrickchagastavares/rover/internal/domain/model"
	"github.com/patrickchagastavares/rover/pkg/scanner"
	"github.com/stretchr/testify/assert"
)

func TestMoveRover(t *testing.T) {
	cases := map[string]struct {
		input        model.Rover
		scanMock     *scanner.TestScan
		expectedErr  error
		expectedData string
	}{
		"success": {
			input:        model.NewRover(1, 2, 5, 5, "N"),
			scanMock:     &scanner.TestScan{Value: []any{"LMLMLMLMM"}, Err: nil},
			expectedErr:  nil,
			expectedData: "1 3 N",
		},
		"success right": {
			input:        model.NewRover(3, 3, 5, 5, "E"),
			scanMock:     &scanner.TestScan{Value: []any{"MMRMMRMRRM"}, Err: nil},
			expectedErr:  nil,
			expectedData: "5 1 E",
		},
		"success with lower case": {
			input:        model.NewRover(1, 2, 5, 5, "N"),
			scanMock:     &scanner.TestScan{Value: []any{"lmlmlmlmm"}, Err: nil},
			expectedErr:  nil,
			expectedData: "1 3 N",
		},
		"don't exceed limit": {
			input:        model.NewRover(1, 2, 5, 5, "N"),
			scanMock:     &scanner.TestScan{Value: []any{"lmlmlmlmmmmmmmlllmmmmm"}, Err: nil},
			expectedErr:  nil,
			expectedData: "5 5 E",
		},
		"don't exceed min value": {
			input:        model.NewRover(1, 2, 5, 5, "N"),
			scanMock:     &scanner.TestScan{Value: []any{"rrMMMMMmmmmmmmrmmmmmmmm"}, Err: nil},
			expectedErr:  nil,
			expectedData: "0 0 W",
		},
		"failed scan terminal": {
			input:        model.NewRover(1, 2, 5, 5, "N"),
			scanMock:     &scanner.TestScan{Value: []any{}, Err: errors.New("failed internal")},
			expectedErr:  fmt.Errorf("failed scan instructions: %w", errors.New("failed internal")),
			expectedData: "1 2 N",
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			scanner.SetTestMode(t, test.scanMock)

			err := moveRover(&test.input)

			assert.Equal(t, test.expectedData, test.input.String())
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestParseRoverData(t *testing.T) {
	cases := map[string]struct {
		inputLine            string
		inputMaxX, inputMaxY int
		scanMock             *scanner.TestScan
		expectedErr          error
		expectedData         model.Rover
	}{
		"success": {
			inputLine: "1 2 n",
			inputMaxX: 5, inputMaxY: 5,
			scanMock:     &scanner.TestScan{Value: []any{1, 2, "n"}, Err: nil},
			expectedErr:  nil,
			expectedData: model.NewRover(1, 2, 5, 5, "n"),
		},
		"failed scan line": {
			inputLine: "1.2 2.3 n",
			inputMaxX: 5, inputMaxY: 5,
			scanMock:     &scanner.TestScan{Value: []any{}, Err: errors.New("failed scan")},
			expectedErr:  fmt.Errorf("failed to scan start position at rover: %w", errors.New("failed scan")),
			expectedData: model.Rover{},
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			scanner.SetTestMode(t, test.scanMock)

			data, err := parseRoverData(test.inputLine, test.inputMaxX, test.inputMaxY)

			assert.Equal(t, test.expectedData, data)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestReadPlateauSize(t *testing.T) {
	cases := map[string]struct {
		scanMock             *scanner.TestScan
		expectedErr          error
		expectedX, expectedY int
	}{
		"success": {
			scanMock:    &scanner.TestScan{Value: []any{5, 5}, Err: nil},
			expectedErr: nil,
			expectedX:   5, expectedY: 5,
		},
		"failed scan line": {
			scanMock:    &scanner.TestScan{Value: []any{}, Err: errors.New("failed scan")},
			expectedErr: fmt.Errorf("failed to scan upper-right coordinates: %w", errors.New("failed scan")),
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			scanner.SetTestMode(t, test.scanMock)

			maxX, maxY, err := readPlateauSize()

			assert.Equal(t, test.expectedX, maxX)
			assert.Equal(t, test.expectedY, maxY)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}
