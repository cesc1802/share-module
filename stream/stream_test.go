package stream

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewConcurrent(t *testing.T) {
	str, push, _close := NewConcurrent(1, 2, 3)

	go func() {
		defer _close()
		push(context.TODO(), 5, 9)
		push(context.TODO(), -3)
	}()

	vals, err := All(str)
	if err != nil {
		t.Fatalf("drain stream: %v", err)
	}

	want := []int{1, 2, 3, 5, 9, -3}
	if !cmp.Equal(want, vals) {
		t.Fatalf("stream returned wrong values\n%s", cmp.Diff(want, vals))
	}
}

func TestConcurrent(t *testing.T) {
	str := make(chan int)
	push := Concurrent(str)

	go func() {
		defer close(str)
		push(context.TODO(), 5)
		push(context.TODO(), 9, -3)
	}()

	vals, err := All(str)
	if err != nil {
		t.Fatalf("drain stream: %v", err)
	}

	want := []int{5, 9, -3}
	if !cmp.Equal(want, vals) {
		t.Fatalf("stream returned wrong values\n%s", cmp.Diff(want, vals))
	}
}

//func TestBefore(t *testing.T) {
//	original := []event.Event{
//		event.New("foo", test.FooEventData{}).Any(),
//		event.New("bar", test.BarEventData{}).Any(),
//		event.New("baz", test.BazEventData{}).Any(),
//	}
//	add := []event.Event{
//		event.New("foo", test.FooEventData{}).Any(),
//		event.New("foobar", test.FoobarEventData{}).Any(),
//	}
//
//	str := New(original)
//
//	str = Before(str, func(evt event.Event) []event.Event {
//		if evt.Name() == "foo" || evt.Name() == "baz" {
//			return add
//		}
//		return nil
//	})
//
//	events, err := Drain(context.Background(), str)
//	if err != nil {
//		t.Fatalf("drain stream: %v", err)
//	}
//
//	want := append(append(add, original[:2]...), append(add, original[2])...)
//
//	if !cmp.Equal(want, events) {
//		t.Fatalf("stream returned wrong events\n%s", cmp.Diff(want, events))
//	}
//}
