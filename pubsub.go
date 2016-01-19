// Package pubsub can be realize pattern of pub-sub in simple way
package pubsub

import (
	"errors"
	"reflect"
	"sync"
)

// PubSub control pub-sub message passing
type PubSub struct {
	c chan interface{}
	m sync.RWMutex
	s []*sub
	e chan error
}

type sub struct {
	f interface{}
}

func NewSubscriber(f interface{}) *sub {
	return &sub{f: f}
}

// New instanciate PubSub object
func New() *PubSub {
	ps := PubSub{
		c: make(chan interface{}),
		e: make(chan error),
	}
	call := func(rf reflect.Value, in []reflect.Value) {
		defer func() {
			if e := recover(); e != nil {
				ps.e <- e.(error)
			}
		}()
		rf.Call(in)
	}
	go func() {
		for v := range ps.c {
			rv := reflect.ValueOf(v)
			ps.m.Lock()
			for _, s := range ps.s {
				go call(reflect.ValueOf(s.f), []reflect.Value{rv})
			}
			ps.m.Unlock()
		}
	}()
	return &ps
}

// Pub publish to subscribers
func (m *PubSub) Pub(v interface{}) {
	m.c <- v
}

// Sub subscribe published data
func (m *PubSub) Sub(s *sub) error {
	rf := reflect.ValueOf(s.f)
	if rf.Kind() != reflect.Func {
		return errors.New("Not a function")
	}
	if rf.Type().NumIn() != 1 {
		return errors.New("Number of arguments should be 1")
	}
	m.m.Lock()
	defer m.m.Unlock()
	m.s = append(m.s, s)
	return nil
}

// UnSub unsubscribe published data
func (m *PubSub) UnSub(us *sub) {
	if us == nil {
		return
	}
	m.m.Lock()
	defer m.m.Unlock()
	result := make([]*sub, 0, len(m.s))
	last := 0
	for i, s := range m.s {
		if reflect.ValueOf(s).Pointer() == reflect.ValueOf(us).Pointer() {
			result = append(result, m.s[last:i]...)
			last = i + 1
		}
	}
	m.s = append(result, m.s[last:]...)
}
