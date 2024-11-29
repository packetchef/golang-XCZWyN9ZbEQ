package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("######## welcome to TODOLIST ########")
	http.HandleFunc("/", handleSayHello)
	http.HandleFunc("/hello", handleSayHello)
	http.HandleFunc("/tasks", handleReturnListWithNewTask)
	http.ListenAndServe(":8888", nil)
}

func handleSayHello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, justSayHello())
}

func handleReturnListWithNewTask(writer http.ResponseWriter, request *http.Request) {
	tList := buildTaskList()
	for _, task := range tList {
		fmt.Fprintln(writer, task)
	}
}

func buildTaskList() []string {
	var t0 = "Watch full course"
	var t1 = "Watch crash course"
	var t2 = "Eat a donut"

	builtList := []string{t0, t1, t2}
	return builtList
}

func mainDemo() {
	var t3 = strings.Repeat("abc", 5)
	fmt.Print(t3, "\n")

	t4 := "abc"
	fmt.Print(t4, "\n")

	tList := buildTaskList()
	fmt.Print("Tasks: ", tList)

	fmt.Println("Printing tasks with a loop:")
	printTasks(tList)

	fmt.Println("Adding a new task:")
	addTask(tList, "Scrubbadub")

	fmt.Println("Printing tasks with a loop:")
	printTasks(tList)

	fmt.Println("adding a new task again:")
	addTaskToOriginal(&tList, "Scoobydooby")
	printTasks(tList)

	fmt.Println("adding a new task again:")
	tList = returnListWithNewTask(tList, "Scrappledapple")
	printTasks(tList)

}

func printTasks(pt []string) {
	fmt.Println("************* PRINTING *************")
	tCount := 1
	// for index, t := range tList {
	for _, t := range pt {
		//fmt.Printf("Task item %d: %s\n", index, t)
		fmt.Printf("Task item %d: %s\n", tCount, t)
		tCount += 1
	}
}

func addTask(oList []string, newTask string) {
	newList := append(oList, newTask)
	printTasks(newList)
}

func addTaskToOriginal(origList *[]string, newTask string) {
	updatedList := append(*origList, newTask)
	fmt.Println("List updation: ", updatedList)
}

func returnListWithNewTask(origList []string, newTask string) []string {
	updatedList := append(origList, newTask)
	return updatedList
}

func justSayHello() string {
	jshStr := "hello world"
	return jshStr
}
