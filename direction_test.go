package chp

import (
	"reflect"
	"testing"
)

func TestReceiveOnly(t *testing.T) {
	var bi chan string
	r := receiveOnly(bi)
	if reflect.TypeOf(r).ChanDir() != reflect.RecvDir {
		t.Errorf("wrong direction %v", reflect.TypeOf(r).ChanDir())
	}
}

func TestSendOnly(t *testing.T) {
	var bi chan string
	s := sendOnly(bi)
	if reflect.TypeOf(s).ChanDir() != reflect.SendDir {
		t.Errorf("wrong direction %v", reflect.TypeOf(s).ChanDir())
	}
}
