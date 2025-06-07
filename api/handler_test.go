package api

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/yourname/taskman/internal/task"
// )

// func TestTaskAPI(t *testing.T) {
// 	task.Clear()
// 	RegisterRoutes()

// 	t.Run("GET /tasks", func(t *testing.T) {
// 		req := httptest.NewRequest("GET", "/tasks", nil)
// 		w := httptest.NewRecorder()

// 		http.DefaultServeMux = http.NewServeMux()
// 		RegisterRoutes()
// 		http.DefaultServeMux.ServeHTTP(w, req)

// 		if w.Code != http.StatusOK {
// 			t.Errorf("Expected 200 OK, got %d", w.Code)
// 		}
// 		var got []task.Task
// 		err := json.NewDecoder(w.Body).Decode(&got)
// 		if err != nil || len(got) != 0 {
// 			t.Errorf("Expected empty task list, got %v", got)
// 		}
// 	})

// 	t.Run("POST /tasks", func(t *testing.T) {
// 		body := bytes.NewBuffer([]byte(`{"name": "New task"}`))
// 		req := httptest.NewRequest("POST", "/tasks", body)
// 		req.Header.Set("Content-Type", "application/json")

// 		w := httptest.NewRecorder()
// 		http.DefaultServeMux = http.NewServeMux()
// 		RegisterRoutes()
// 		http.DefaultServeMux.ServeHTTP(w, req)

// 		if w.Code != http.StatusOK {
// 			t.Errorf("Expected 200 OK, got %d", w.Code)
// 		}
// 		var added task.Task
// 		json.NewDecoder(w.Body).Decode(&added)
// 		if added.Name != "New task" {
// 			t.Errorf("Expected 'New task', got %v", added.Name)
// 		}
// 	})
// }
