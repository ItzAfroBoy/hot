package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	// tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
	// "github.com/mattn/go-isatty"
)

var helpStyle = lg.NewStyle().Foreground(lg.Color("241")).Render

// type model struct {
// 	filename string
// 	command string
// }

// func (m model) Init() tea.Cmd {}
// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {}
// func (m model) View() string {}

func exit(str string, code int) {
	fmt.Printf("Error: %s", str)
	os.Exit(code)
}

func watchFiles(w *fsnotify.Watcher, files []string, cmd string) {
	for {
		select {
		case err, ok := <-w.Errors:
			if !ok {
				return
			}
			fmt.Printf("Error: %s\n", err.Error())
		case e, ok := <-w.Events:
			if !ok {
				return
			}

			var found bool
			for _, f := range files {
				if f == e.Name {
					found = true
				}
			}

			if !found {
				continue
			}

			if e.Has(fsnotify.Write) {
				args := strings.Split(cmd, ";")
				fmt.Printf("File has changed. Executing command: %s\n", strings.Join(args, " "))
				out, err := exec.Command(args[0], args[1:]...).Output()
				if err != nil {
					fmt.Printf("Error: %s", err.Error())
				} else {
					fmt.Printf("Command output: %s\n", string(out))
				}
			}
		}
	}
}

func main() {
	cmd := flag.String("cmd", "go;build", "Command to exec on file change. Args to be sperated with ';'")

	flag.Parse()
	files := flag.Args()

	if len(files) < 1 {
		exit("", 1)
	}

	w, err := fsnotify.NewWatcher()
	if err != nil {
		exit(err.Error(), 1)
	}
	defer w.Close()

	go watchFiles(w, files, *cmd)
	for _, p := range files {
		st, err := os.Lstat(p)
		if err != nil {
			exit(err.Error(), 1)
		}

		if st.IsDir() {
			exit(p+" is a directory, not a file", 1)
		}

		err = w.Add(filepath.Dir(p))
		if err != nil {
			exit(p+": "+err.Error(), 1)
		}
	}

	fmt.Println(helpStyle("hot is ready: press ^C to exit"))
	<-make(chan struct{})
}
