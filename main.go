package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	DocumentChoices []string
	Cursor          int
	Selected        map[int]struct{}
}

func initialModel() Model {
	return Model{
		DocumentChoices: []string{"a", "b", "c"},
		Selected:        make(map[int]struct{}),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "j":
			if m.Cursor < len(m.DocumentChoices)-1 {
				m.Cursor++
			}
		case "enter":
			if _, ok := m.Selected[m.Cursor]; !ok {
				m.Selected[m.Cursor] = struct{}{}
			} else {
				delete(m.Selected, m.Cursor)
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	s := "Documents:\n\n"

	for i, choice := range m.DocumentChoices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.Selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit."

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	_, err := p.Run()
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
