package task

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
)

// Task represents the task entity
type Task struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var tasks = []Task{
	{ID: 1, Title: "P0", Content: "This is a high priority task"},
	{ID: 2, Title: "P2", Content: "This is a low priority task"},
}

// @Summary List all tasks
// @Description Get a list of all tasks
// @Produce json
// @Success 200 {array} Task
// @Router /tasks [get]
func ListTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// @Summary Get a task by ID
// @Description Get a task by its ID
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} Task
// @Failure 400 {string} string "Invalid task ID"
// @Failure 404 {string} string "Task not found"
// @Router /tasks/{id} [get]
func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	taskID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	for _, task := range tasks {
		if task.ID == taskID {
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	http.NotFound(w, r)
}

// @Summary Create a new task
// @Description Create a new task
// @Accept json
// @Produce json
// @Param task body Task true "Task object"
// @Success 200 {object} Task
// @Router /tasks [post]
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task Task

	_ = json.NewDecoder(r.Body).Decode(&task)

	// Sanitize user provided content to remove any harmful HTML elements
	task.Content = bluemonday.UGCPolicy().Sanitize(task.Content)
	task.Title = bluemonday.UGCPolicy().Sanitize(task.Title)

	// Check if task already exists
	for _, t := range tasks {
		if t.Title == task.Title && t.Content == task.Content {
			json.NewEncoder(w).Encode(t)
			return
		}
	}

	task.ID = len(tasks) + 1
	tasks = append(tasks, task)

	json.NewEncoder(w).Encode(task)
}

// @Summary Update an existing task
// @Description Update an existing task by its ID
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param task body Task true "Task object"
// @Success 200 {object} Task
// @Failure 400 {string} string "Invalid task ID"
// @Failure 404 {string} string "Task not found"
// @Router /tasks/{id} [put]
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	taskID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	for index, task := range tasks {
		if task.ID == taskID {
			var updatedTask Task
			_ = json.NewDecoder(r.Body).Decode(&updatedTask)
			updatedTask.ID = taskID

			// Sanitize user provided content to remove any harmful HTML elements
			updatedTask.Content = bluemonday.UGCPolicy().Sanitize(updatedTask.Content)
			updatedTask.Title = bluemonday.UGCPolicy().Sanitize(updatedTask.Title)
			tasks[index] = updatedTask
			json.NewEncoder(w).Encode(updatedTask)
			return
		}
	}

	http.NotFound(w, r)
}

// @Summary Delete a task
// @Description Delete a task by its ID
// @Param id path int true "Task ID"
// @Success 204 "No Content"
// @Failure 400 {string} string "Invalid task ID"
// @Failure 404 {string} string "Task not found"
// @Router /tasks/{id} [delete]
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	taskID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	for index, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:index], tasks[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.NotFound(w, r)
}
