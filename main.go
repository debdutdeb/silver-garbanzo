package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices   []string
	cursor    int
	selection map[int]struct{}
}

func initModel() *model {
	return &model{
		choices:   []string{"red", "green", "black", "white"},
		cursor:    0,
		selection: make(map[int]struct{}),
	}
}

func (m *model) View() string {
	s := "Which color?\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	return s
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		{
			switch msg.String() {
			case "j": // down
				{
					m.cursor = (m.cursor + 1) % len(m.choices)
				}
			case "k": // up
				{
					m.cursor = len(m.choices) + (m.cursor - 1)
				}
			}
		}
	}
	return m, nil
}

func (m *model) Init() tea.Cmd {
	return nil
}

func main() {
	t := tea.NewProgram(initModel())
	if err := t.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
}
