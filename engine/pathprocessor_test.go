package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_process(t *testing.T) {
	testData := []struct {
		name     string
		expected []string
		err      error
		input    [][]string
	}{
		{"Happy single", []string{"SFO", "EWR"}, nil, [][]string{{"SFO", "EWR"}}},
		{"Happy easy", []string{"SFO", "EWR"}, nil, [][]string{{"ATL", "EWR"}, {"SFO", "ATL"}}},
		{"Happy more complex", []string{"SFO", "EWR"}, nil,
			[][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}}},
		{"Happy with inner cycle", []string{"SFO", "EWR"}, nil,
			[][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}, {"EWR", "ATL"}, {"ATL", "SFO"}, {"SFO", "EWR"}}},
		{"Error Simple Cycle", nil, newBadDataErr(cyclicalErr, nil, nil), [][]string{{"ATL", "EWR"}, {"SFO", "ATL"}, {"EWR", "SFO"}}},
		{"Error Complex Cycle", nil, newBadDataErr(cyclicalErr, nil, nil),
			[][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}, {"EWR", "ATL"}, {"ATL", "SFO"}}},
		{"Error multiple disjoined paths (added)", nil, newBadDataErr(multipleStartsEndsErr, nil, nil),
			[][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}, {"FOO", "BAR"}}},
		{"Error multiple disjoined paths (removed)", nil, newBadDataErr(multipleStartsEndsErr, nil, nil),
			[][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}}},
		{"Error multiple disjoined paths (duplicates)", nil, newBadDataErr(multipleStartsEndsErr, nil, nil),
			[][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}}},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			res, err := process(td.input)
			if td.err == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorContains(t, err, td.err.Error())
			}
			assert.Equal(t, td.expected, res)
		})
	}
}
