package todo_test

import (
	"encoding/json"
	"io"
	"os"
	"testing"

	"powerful-command-line-applications-in-go/todo"
)

func TestAdd(t *testing.T) {
	ls := todo.List{}
	w := "Test my code"
	ls.Add(w)

	if ls[0].What != w {
		t.Errorf("Expected '%s', but got '%s' instead.", w, ls[0].What)
	}
}

func TestComplete(t *testing.T) {
	ls := todo.List{}
	w := "Test my code, again!"
	ls.Add(w)

	if ls[0].Completed {
		t.Errorf("Item '%s' should not be completed by default.", ls[0].What)
	}

	ls.Complete(0)
	if !ls[0].Completed {
		t.Errorf("Item '%s' should be marked as completed after 'Complete' call.", ls[0].What)
	}
}

func TestDelete(t *testing.T) {
	ls := todo.List{}
	w := "Test my code, yet again!"
	ls.Add(w)
	ls.Add("Some")
	ls.Add("random")
	ls.Add("text")

	ls.Delete(0)
	if contains(ls, w) {
		t.Errorf("Item '%s' should have been deleted after 'Delete' call.", w)
	}
}

func TestSave(t *testing.T) {
	ls := todo.List{}
	w := "Test my code, one more time!"
	ls.Add(w)

	f, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("Could not create temporary file. Aborting with error %s", err)
	}
	defer os.Remove(f.Name()) // Defers file deletion to when function ends.

	if err := ls.Save(f.Name()); err != nil {
		t.Fatalf("Could not save todo JSON data to file '%s'. Aborting with error %s", f.Name(), err)
	}

	fc, err := io.ReadAll(f) // Reads file contents.
	if err != nil {
		t.Fatalf("Could not read from file '%s'. Aborting with error %s", f.Name(), err)
	}

	var ls2 todo.List
	if err := json.Unmarshal([]byte(fc), &ls2); err != nil { // Parses JSON data from file
		t.Fatalf("Could not parse todo JSON data from file '%s'. Aborting with error %s", f.Name(), err)
	}

	if ls[0].What != ls2[0].What {
		t.Errorf("Expected todo list with item '%s', but got item '%s' instead.", w, ls2[0].What)
	}
}

func TestGet(t *testing.T) {
	ls := todo.List{}
	w := "Test my code, one last time!"
	ls.Add(w)

	f, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("Could not create temporary file. Aborting with error %s", err)
	}
	// Defers the following call until when TestSave returns.
	defer os.Remove(f.Name())

	if err := ls.Save(f.Name()); err != nil {
		t.Fatalf("Could not save todo JSON data to file '%s'. Aborting with error %s", f.Name(), err)
	}

	ls2 := todo.List{}
	if err := ls2.Get(f.Name()); err != nil { // Parses JSON data from file
		t.Fatalf("Could not Get todo list data from file '%s'. Aborting with error %s", f.Name(), err)
	}

	if ls[0].What != ls2[0].What {
		t.Errorf("Expected todo list with item '%s', but got item '%s' instead.", w, ls2[0].What)
	}
}

// Helper functions
func contains(l todo.List, s string) bool {
	for _, i := range l {
		if i.What == s {
			return true
		}
	}
	return false
}
