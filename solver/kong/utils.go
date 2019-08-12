package kong

import (
	"github.com/alpiquero/deck/crud"
	"github.com/alpiquero/deck/diff"
)

func eventFromArg(arg crud.Arg) diff.Event {
	event, ok := arg.(diff.Event)
	if !ok {
		panic("unexpected type, expected diff.Event")
	}
	return event
}
