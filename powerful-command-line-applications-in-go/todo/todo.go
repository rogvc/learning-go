package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type task struct {
	What        string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []task

func (ls *List) Add(w string) {
	t := task{
		What:        w,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*ls = append(*ls, t)
}

func (ls *List) Complete(i int) error {
	if i < 0 || i > len(*ls) || len(*ls) == 0 {
		return fmt.Errorf("item #%d does not exist", i)
	}

	(*ls)[i].Completed = true
	(*ls)[i].CompletedAt = time.Now()

	return nil
}

func (ls *List) Delete(i int) error {
	if i < 0 || i > len(*ls) || len(*ls) == 0 {
		return fmt.Errorf("item #%d does not exist", i)
	}

	*ls = append((*ls)[:i], (*ls)[i+1:]...)

	return nil
}

func (ls *List) Save(f string) error {
	js, err := json.Marshal(ls)
	if err != nil {
		return err
	}

	return os.WriteFile(f, js, 0644)
}

func (ls *List) Get(fn string) error {
	f, err := os.ReadFile(fn)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil // File doesn't exist. Nothing to read.
		}
		return err
	}

	if len(f) == 0 {
		return nil
	}

	return json.Unmarshal(f, ls)
}

func (ls *List) String() string {
	var out string
	for i, t := range *ls {
		out += fmt.Sprintf("(%d) ", i)
		out += t.What
		if t.Completed {
			out += " [✅]"
		} else {
			out += " [ ]"
		}
		if i < len(*ls)-1 {
			out += fmt.Sprintln()
		}
	}

	return out
}
