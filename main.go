package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/debdutdeb/silver-garbanzo/handlers"
)


func main() {
    h, err := handlers.New()
    if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
    }

	t := tea.NewProgram(h)
	if err := t.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
}
