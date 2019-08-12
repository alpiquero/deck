package diff

import (
	"github.com/alpiquero/deck/crud"
)

// Event represents an event to perform
// an imperative operation
// that gets Kong closer to the target state.
type Event struct {
	Op     crud.Op
	Kind   crud.Kind
	Obj    interface{}
	OldObj interface{}
}
