package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"powerful-command-line-applications-in-go/todo"
)

const todoFile = "todo.json"
const todoFileEnvVar = "TODO_FILE_PATH"

func main() {
	var (
		t, df string
		l, a  bool
		c, d  int
	)
	flag.StringVar(&t, "task", "", "Task to add")
	flag.StringVar(&t, "t", "", "Task to add")
	flag.StringVar(&df, "data", todoFile, "Path of the data file to use (does not overwrite environment setting)")
	flag.StringVar(&df, "D", todoFile, "Path of the data file to use (does not overwrite environment setting)")

	flag.BoolVar(&l, "list", false, "List current tasks")
	flag.BoolVar(&l, "l", false, "List current tasks")
	flag.BoolVar(&a, "add", false, "Use this flag if you want to add tasks by typing them into the program")
	flag.BoolVar(&a, "a", false, "Use this flag if you want to add tasks by typing them into the program")

	flag.IntVar(&c, "complete", -1, "Index of the task to be marked as complete")
	flag.IntVar(&c, "c", -1, "Index of the task to be marked as complete")
	flag.IntVar(&d, "delete", -1, "Index of the task to be deleted")
	flag.IntVar(&d, "d", -1, "Index of the task to be deleted")

	flag.Parse()

	ls := &todo.List{}

	// Checks if data file path has been set in the environment
	if edf := os.Getenv(todoFileEnvVar); edf != "" {
		df = edf
	}

	if err := ls.Get(df); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case t != "":
		ls.Add(t)
		saveList(*ls, df)
	case c >= 0:
		ls.Complete(c)
		saveList(*ls, df)
	case d >= 0:
		ls.Delete(d)
		saveList(*ls, df)
	case a:
		fmt.Printf("Please enter your task: ")
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		ls.Add(t)
		saveList(*ls, df)
	case l:
		if len(*ls) == 0 {
			fmt.Println("Nothing to do!")
		} else {
			fmt.Println(ls.String())
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid usage.")
		flag.Usage()
		os.Exit(1)
	}
}

// Helper functions
func saveList(ls todo.List, df string) {
	if err := ls.Save(df); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		return "", fmt.Errorf("task cannot be blank")
	}

	return s.Text(), nil
}
