package handlers

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *Background) View() string {
	s := "Choose a file:-\n\n"

	for num, file := range m.CurrDirs {
		var choice string = " "
		var selected string = " "
		if m.cursor == num {
			choice = "?"
		}
		if m.selectedIndex == num {
			selected = "*"
		}
		s += fmt.Sprintf("%s [%s] %s\n", choice, selected, file)
	}

	s += "\npress q to quit\n"

	return s
}

func (m *Background) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		{
			switch msg.String() {
			case "q", "ctrl+c":
				{
					return m, tea.Quit
				}
			case "enter", " ":
				{
					m.selectedIndex = m.cursor
				}
			case "j":
				{
					if m.cursor < len(m.CurrDirs) {
                        m.cursor++
                    }
				}
			case "k":
				{
                    if m.cursor > 0 {
                        m.cursor--
                    }
				}
			}
		}
	}

	return m, nil
}

func (m *Background) Init() tea.Cmd {
	return nil
}
