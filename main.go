package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(`
Exactly one parameter is required: the name of the text editor.

Supported options:
  With Sublime on Linux: ./sp.exe subl
With Notepad on Windows: sp.exe notepad
		`)
		return
	}

	p := os.Args[1]
	if p == "subl" || p == "notepad" {
		launch(p, notesPath())
	}
}

func launch(editor string, fn string) {
	cmd := exec.Command(editor, fn)
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func notesPath() string {
	return filepath.Join(utcYear(), txtFilename())
}

func utcYear() string {
	return fmt.Sprint(time.Now().UTC().Year())
}

func txtFilename() string {
	return fmt.Sprintf("%s.txt", formattedTimestamp(time.Now()))
}

func formattedTimestamp(now time.Time) string {
	return now.Format("2006-01-02-1504")
}
