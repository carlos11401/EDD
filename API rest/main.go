package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type task struct {
	Id      int    `json:Id`
	Name    string `json:Name`
	Content string `json:Content`
}
type allTasks []task

var tasks = allTasks{
	{
		Id:      1,
		Name:    "no 1",
		Content: "con 1",
	},
}

// Just show a massage
func indexRout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "WELCOME")
}

// receive the information of the body
func getTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	// get the information of BODY
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil { // If nothing is in the BODY
		fmt.Fprintf(w, "Insert a Valid Task")
	}
	// Pass dates of the BODY(json) at newTask
	json.Unmarshal(reqBody, &newTask)
	// add date ID a newTask. Count all tasks and plus one
	newTask.Id = len(tasks) + 1
	// add newTask a tasks
	tasks = append(tasks, newTask)
	// just to the type of date
	w.Header().Set("Content-Type", "aplication/json")
	// if went correct
	w.WriteHeader(http.StatusCreated)
	// Send date
	json.NewEncoder(w).Encode(newTask)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	// get value of parameters
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}
	for _, task := range tasks {
		if task.Id == taskId {
			// type of date to show
			w.Header().Set("Content-Type", "application/json")
			//show date
			json.NewEncoder(w).Encode(task)
		}
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}
	for i, task := range tasks {
		if task.Id == taskId {
			// [:i} conserve all before, [i+1:]... conserve all after
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprintf(w, "Delete Succesfully")
		}
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var updateTask task
	taskId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter Valid Date")
	}
	json.Unmarshal(reqBody, &updateTask)

	for i, task := range tasks {
		if task.Id == taskId {
			tasks = append(tasks[:i], tasks[i+1:]...)
			updateTask.Id = taskId
			tasks = append(tasks, updateTask)
			fmt.Fprintf(w, "ID %v Has been Updates Succesfully", taskId)
		}
	}
}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRout)
	// show tasks
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	//create task
	router.HandleFunc("/task", createTask).Methods("POST")
	// search task for index
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	// delete for index
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	// update for index
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", router))
}
