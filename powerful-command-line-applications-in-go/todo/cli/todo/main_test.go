package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var (
	binName  = "todo"
	fileName = "todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool '%s'. Aborting with error: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	r := m.Run()

	fmt.Println("Cleaning up...")
	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(r)
}

func TestTodoCLI(t *testing.T) {
	tsk := "Test it again!"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cp := filepath.Join(dir, binName)

	t.Run("AddNewTask", func(t *testing.T) {
		cmd := exec.Command(cp, "--add", tsk)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cp, "--list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		if !strings.Contains(string(out), tsk) {
			t.Errorf("Expected '%s', but got '%s' instead.", tsk, string(out))
		}
	})

	t.Run("DeleteTasks", func(t *testing.T) {
		cmd := exec.Command(cp, "--delete", "0")
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}

		lst := exec.Command(cp, "--list")
		out, err := lst.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		if !strings.Contains(string(out), "Nothing to do!") {
			t.Errorf("Expected '%s' to have been removed, but it was still present.", tsk)
		}
	})
}
