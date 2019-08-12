package dry

import (
	"github.com/alpiquero/deck/crud"
	"github.com/alpiquero/deck/diff"
	"github.com/alpiquero/deck/print"
	"github.com/alpiquero/deck/state"
	"github.com/alpiquero/deck/utils"
	"github.com/hbagdi/go-kong/kong"
)

// TargetCRUD implements Actions interface
// from the github.com/kong/crud package for the Target entitiy of Kong.
type TargetCRUD struct {
	// client    *kong.Client
	// callbacks []Callback // use this to update the current in-memory state
}

func targetFromStuct(arg diff.Event) *state.Target {
	target, ok := arg.Obj.(*state.Target)
	if !ok {
		panic("unexpected type, expected *state.Target")
	}

	return target
}

// Create creates a fake Target.
// The arg should be of type diff.Event, containing the target to be created,
// else the function will panic.
// It returns a the created *state.Target.
func (s *TargetCRUD) Create(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	target := targetFromStuct(event)
	print.CreatePrintln("creating target", *target.Target.Target,
		"on upstream", *target.Upstream.Name)
	target.ID = kong.String(utils.UUID())
	return target, nil
}

// Delete deletes a fake Target.
// The arg should be of type diff.Event, containing the target to be deleted,
// else the function will panic.
// It returns a the deleted *state.Target.
func (s *TargetCRUD) Delete(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	target := targetFromStuct(event)
	print.DeletePrintln("deleting target", *target.Target.Target,
		"from upstream", *target.Upstream.Name)
	return target, nil
}

// Update updates a fake Target.
// The arg should be of type diff.Event, containing the target to be updated,
// else the function will panic.
// It returns a the updated *state.Target.
func (s *TargetCRUD) Update(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	target := targetFromStuct(event)
	oldTarget, ok := event.OldObj.(*state.Target)
	if !ok {
		panic("unexpected type, expected *state.Target")
	}
	print.DeletePrintln("deleting target", *oldTarget.Target.Target,
		"from upstream", *oldTarget.Upstream.Name)
	print.CreatePrintln("creating target", *target.Target.Target,
		"on upstream", *target.Upstream.Name)
	return target, nil
}
