package main

import (
	"bufio"
	"fmt"
	"os"
)
import (
	"flag"
)

func main() {
	command := flag.String("command", "nothing", "to say your name")
	flag.Parse()
	for {
		runCommand(*command)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Please enter another command:")
		scanner.Scan()
		*command = scanner.Text()
	}
}

func runCommand(command string) {
	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "exit":
		fmt.Println("see you later ;)")
		os.Exit(0)
	}
}
func createTask() {
	scanner := bufio.NewScanner(os.Stdin)
	var name, dodate, category string

	fmt.Println("please enter task NAME")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter task DoDATE")
	scanner.Scan()
	dodate = scanner.Text()

	fmt.Println("please enter task CATEGORY")
	scanner.Scan()
	category = scanner.Text()

	fmt.Println("Task: ", name, dodate, category)
}

func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)
	var name, color string

	fmt.Println("please Enter category Name")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please Enter category color")
	scanner.Scan()
	color = scanner.Text()

	fmt.Println("Category: ", name, color)
}
