package fsm_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"free5gc/lib/fsm"
	"testing"
)

const (
	ACITVE    fsm.State = "ACTIVE"
	INACITVE  fsm.State = "INACITVE"
	EXCEPTION fsm.State = "EXCEPTION"
)

const (
	MESSAGE fsm.Event = "MESSAGE"
)

func Activefunc(sm *fsm.FSM, event fsm.Event, args fsm.Args) error {
	switch event {
	case fsm.EVENT_ENTRY:
		fmt.Printf("Enter Active state\n")
	case MESSAGE:
		a := args["Message"].(string)
		fmt.Printf("Active : %s\n", a)
	}
	return nil
}
func Inactivefunc(sm *fsm.FSM, event fsm.Event, args fsm.Args) error {
	switch event {
	case fsm.EVENT_ENTRY:
		fmt.Printf("Enter Inactive state\n")
	case MESSAGE:
		a := args["Message"].(string)
		fmt.Printf("Inactive : %s\n", a)
	}
	return nil
}
func Exceptionfunc(sm *fsm.FSM, event fsm.Event, args fsm.Args) error {
	switch event {
	case fsm.EVENT_ENTRY:
		param1 := args["Init"].(string)
		fmt.Printf("Enter Exception state with param %s\n", param1)
	case MESSAGE:
		a := args["Message"].(string)
		fmt.Printf("Exception : %s\n", a)
	}
	return nil
}

func TestInitFSM(t *testing.T) {
	table := fsm.NewFuncTable()
	table[ACITVE] = Activefunc
	table[INACITVE] = Inactivefunc
	sm, err := fsm.NewFSM(ACITVE, table)
	assert.Equal(t, nil, err)
	a := "hahaha"
	err = sm.SendEvent(MESSAGE, fsm.Args{"Message": a})
	assert.Equal(t, nil, err)
	err = sm.Transfer(INACITVE, nil)
	assert.Equal(t, nil, err)
	err = sm.SendEvent(MESSAGE, fsm.Args{"Message": a})
	assert.Equal(t, nil, err)
	err = sm.Transfer(EXCEPTION, fsm.Args{"Message": "p1"})
	assert.Equal(t, true, err != nil)
	sm.AddState(EXCEPTION, Exceptionfunc)
	err = sm.Transfer(EXCEPTION, fsm.Args{"Init": "p1"})
	assert.Equal(t, nil, err)
	err = sm.SendEvent(MESSAGE, fsm.Args{"Message": a})
	assert.Equal(t, nil, err)
	err = sm.Transfer("asd", nil)
	assert.Equal(t, true, err != nil)
	sm.PrintStates()
	assert.Equal(t, true, sm.Check(EXCEPTION))
}
