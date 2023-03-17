package engine

import (
	"reflect"
	"strings"
)

const (
	cyclicalErr           = "The itinerary is a cycle"
	multipleStartsEndsErr = "Multiple starting or destination points found"
)

type BadDataErr struct {
	Value string
	Type  reflect.Type
	Start []string
	End   []string
}

func newBadDataErr(msg string, start, end []string) error {
	return &BadDataErr{
		Value: msg,
		Start: start,
		End:   end,
	}
}

func (e *BadDataErr) Error() string {
	return e.String()
}

func (e *BadDataErr) String() string {
	var b strings.Builder
	b.WriteString(e.Value)
	b.WriteString(": ")
	if e.Start != nil {
		b.WriteString("Start ")
		b.WriteString(strings.Join(e.Start, ", "))
	}
	if e.End != nil {
		b.WriteString("End ")
		b.WriteString(strings.Join(e.End, ", "))
	}
	return b.String()
}
