package task

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/tasks", ListTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", GetTask).Methods("GET")
	router.HandleFunc("/tasks", CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")
	return router
}

func TestTasks(t *testing.T) {
	router := setupRouter()

	t.Run("ListTasks", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/tasks", nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "application/json", response.Header().Get("Content-Type"))

		var tasks []Task
		json.Unmarshal(response.Body.Bytes(), &tasks)

		assert.Len(t, tasks, 2)
	})

	t.Run("GetTask", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/tasks/1", nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "application/json", response.Header().Get("Content-Type"))

		var task Task
		json.Unmarshal(response.Body.Bytes(), &task)

		assert.Equal(t, 1, task.ID)
	})

	t.Run("GetTask_NotFound", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/tasks/100", nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("CreateTask", func(t *testing.T) {
		task := Task{Title: "New Task", Content: "Content for New Task"}
		taskJSON, _ := json.Marshal(task)
		request, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "application/json", response.Header().Get("Content-Type"))

		var createdTask Task
		json.Unmarshal(response.Body.Bytes(), &createdTask)

		assert.Equal(t, 3, createdTask.ID)
	})

	t.Run("UpdateTask", func(t *testing.T) {
		task := Task{ID: 1, Title: "Updated Task", Content: "Updated Content"}
		taskJSON, _ := json.Marshal(task)
		request, _ := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(taskJSON))
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "application/json", response.Header().Get("Content-Type"))

		var updatedTask Task
		json.Unmarshal(response.Body.Bytes(), &updatedTask)

		assert.Equal(t, "Updated Task", updatedTask.Title)
	})

	t.Run("DeleteTask", func(t *testing.T) {
		request, _ := http.NewRequest("DELETE", "/tasks/1", nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNoContent, response.Code)

		// Check if the task was deleted
		request, _ = http.NewRequest("GET", "/tasks/1", nil)
		response = httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
}
