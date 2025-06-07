package api

import (
	"encoding/json"
	"net/http"
	"taskman/internal/task"
)

func RegisterRoutes() {
	http.HandleFunc("/tasks", taskHandler)
}


func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tasks := task.GetAll()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)

	case http.MethodPost:
		var t task.Task
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request body"))
			return
		}

		added := task.Add(t.Name)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(added)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}