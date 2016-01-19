package pubsub_test

import (
	"testing"

	"github.com/hatajoe/pubsub"
)

func TestLeave(t *testing.T) {
	done := make(chan int)
	ps := pubsub.New()

	f := func(i int) {
		done <- i
	}
	s1 := pubsub.NewSubscriber(f)
	s2 := pubsub.NewSubscriber(f)
	ps.Sub(s1)
	ps.Sub(s2)
	ps.Pub(1)
	i1 := <-done
	i2 := <-done
	if i1 != 1 || i2 != 1 {
		t.Fatal("Expected multiple subscribers")
	}
	ps.UnSub(s1)
	ps.UnSub(s2)
	ps.Pub(2)
	select {
	case <-done:
		t.Fatal("WTF")
	default:
	}
}
