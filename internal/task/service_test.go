package task

import "testing"

func TestAddGetAll(t *testing.T) {
	Clear()
	t1 := Add("Test Task 1")
	if t1.ID != 1 {
		t.Errorf("Expected ID 1, got %d", t1.ID)
	}

	t2 := Add("Test Task 2")
	if t2.ID != 2 {
		t.Errorf("Expected ID 1, got %d", t2.ID)
	}

	tasks := GetAll()
	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
}