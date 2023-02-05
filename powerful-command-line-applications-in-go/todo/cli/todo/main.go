package main

import (
	"flag"
	"fmt"
	"os"

	"powerful-command-line-applications-in-go/todo"
)

const todoFile = "todo.json"

func main() {
	var (
		a    string
		l    bool
		c, d int
	)
	flag.StringVar(&a, "add", "", "Task to add")
	flag.StringVar(&a, "a", "", "Task to add")
	flag.BoolVar(&l, "list", false, "List current tasks")
	flag.BoolVar(&l, "l", false, "List current tasks")
	flag.IntVar(&c, "complete", -1, "Index of the task to be marked as complete")
	flag.IntVar(&c, "c", -1, "Index of the task to be marked as complete")
	flag.IntVar(&d, "delete", -1, "Index of the task to be deleted")
	flag.IntVar(&d, "d", -1, "Index of the task to be deleted")
	flag.Parse()

	ls := &todo.List{}
	if err := ls.Get(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case a != "":
		ls.Add(a)
		saveList(*ls)
	case c >= 0:
		ls.Complete(c)
		saveList(*ls)
	case d >= 0:
		ls.Delete(d)
		saveList(*ls)
	case l:
		if len(*ls) == 0 {
			fmt.Println("Nothing to do!")
		}
		for i, t := range *ls {
			fmt.Printf("(%d) ", i)
			fmt.Print(t.What)
			if t.Completed {
				fmt.Println(" [✅]")
			} else {
				fmt.Println(" [ ]")
			}
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid usage.")
		flag.Usage()
		os.Exit(1)
	}
}

// Helper functions
func saveList(ls todo.List) {
	if err := ls.Save(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
