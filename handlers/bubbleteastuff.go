package handlers

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *Background) View() string {
	s := "Choose a file:-\n\n"

	for num, file := range m.CurrDirs {
		var choice string = " "
		if m.cursor == num {
			choice = "*"
		}
		s += fmt.Sprintf("[%s] %s\n", choice, file)
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
					m.selection = m.cursor
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
			case "l":
				{
					m.MoveTo(m.CurrDirs[m.cursor])
				}
			case "b":
				{
                    // restore last position
                    b, _ := m.Pos.Pop().(Background)
                    m = &b
				}
			}
		}
	}

	return m, nil
}

func (m *Background) Init() tea.Cmd {
	return nil
}
