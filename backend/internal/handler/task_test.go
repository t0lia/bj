package handler

import "testing"

func TestGetTask(t *testing.T) {
	got := getTask()

	allowed := map[string]bool{
		"buy milk":        true,
		"walk the dog":    true,
		"reply to emails": true,
	}

	if !allowed[got] {
		t.Fatalf("getTask() = %q; expected one of %#v", got, defaultTasks)
	}
}
