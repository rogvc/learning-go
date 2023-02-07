package main_test

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var (
	tdir = "test"              // Directory where to place temp files
	bin  = tdir + "/todo"      // Name of the binary to use for testing
	df   = tdir + "/todo.json" // Name of the data file to use for testing
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")
	if runtime.GOOS == "windows" {
		bin += ".exe"
	}

	build := exec.Command("go", "build", "-o", bin)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool '%s'. Aborting with error: %s", bin, err)
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	r := m.Run()

	fmt.Println("Cleaning up...")
	os.RemoveAll(tdir)

	os.Exit(r)
}

func TestTodoCLI(t *testing.T) {
	tsk := "First Task!"
	tsk2 := "Second task!"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cp := filepath.Join(dir, bin)

	t.Run("AddNewTask", func(t *testing.T) {
		cmd := exec.Command(cp, "--task", tsk, "--data", df)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("AddNewStdinTask", func(t *testing.T) {
		cmd := exec.Command(cp, "--add", "--data", df)
		cmdStdin, err := cmd.StdinPipe()
		if err != nil {
			t.Fatal(err)
		}
		io.WriteString(cmdStdin, tsk2)
		cmdStdin.Close()

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cp, "--list", "--data", df)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		if !strings.Contains(string(out), tsk) {
			t.Errorf("Expected '%s', but got '%s' instead.", tsk, string(out))
		}
	})

	t.Run("DeleteTasks", func(t *testing.T) {
		cmd := exec.Command(cp, "--delete", "0", "--data", df)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}

		lst := exec.Command(cp, "--list")
		out, err := lst.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		if strings.Contains(string(out), tsk) {
			t.Errorf("Expected '%s' to have been removed, but it was still present.", tsk)
		}
	})

	t.Run("CompleteTasks", func(t *testing.T) {
		cmd := exec.Command(cp, "--complete", "0", "--data", df)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}

		lst := exec.Command(cp, "--list")
		out, err := lst.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		if !strings.Contains(string(out), "[✅]") {
			t.Errorf("Expected '%s' to have been marked as complete, but it was not completed.", tsk2)
		}
	})
}
