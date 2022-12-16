package main

import (
	"fmt"
	"strings"
)

func main() {

	var firstName string
	var task string
	var taskInput int
	var input int

	tasksList := map[string][]string{
		"anind": {"clean room"},
        "achint": {"Play fifa"},
        "sriya": {"eat"},
	}

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Do you wish to add task?")
	fmt.Println("Press 1 to add, 2 to delete and 3 to view")
	fmt.Scan(&input)

	if input == 1 {
		fmt.Println("Please enter the task")
		fmt.Scan(&task)

		// Check if the user's first name is already in the map
		if tasksList[strings.ToLower(firstName)] != nil {
			// If the first name is already in the map, append the new task to the existing list of tasks
			tasksList[strings.ToLower(firstName)] = append(tasksList[strings.ToLower(firstName)], task)
		} else {
			// If the first name is not in the map, create a new entry with the first name and the new task
			tasksList[strings.ToLower(firstName)] = []string{task}
		}

		for k, v := range tasksList {
			fmt.Printf("%v your tasks are :- \n", k)
			for a, b := range v {
				fmt.Println(a+1, "-", b)
			}
		}

	} else if input == 2 {
		// Ask the user to enter the index of the task they want to delete
		fmt.Println("Please enter the index of the task you want to delete:")
		fmt.Scan(&taskInput)

		// Delete the task using the slice function
		tasksList[strings.ToLower(firstName)] = append(tasksList[strings.ToLower(firstName)][:taskInput-1], tasksList[strings.ToLower(firstName)][taskInput:]...)
        for k, v := range tasksList {
            fmt.Printf("%v your tasks are :- \n", k)
            for a, b := range v {
                fmt.Println(a+1, "-", b)
            }
        }
	} else {
		for k, v := range tasksList {
			fmt.Printf("%v your tasks are :- \n", k)
			for a, b := range v {
				fmt.Println(a+1, "-", b)
			}
		}

	}

}
