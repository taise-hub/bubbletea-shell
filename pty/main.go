package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type shellFinishedMsg struct{ err error }

type model struct {
	err error
}

func openShell() tea.Cmd {
	return tea.Exec(&Terminal{}, func(err error) tea.Msg {
		return shellFinishedMsg{err}
	})
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "s":
			return m, openShell()
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case shellFinishedMsg:
		if msg.err != nil {
			m.err = msg.err
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return "Error: " + m.err.Error() + "\n"
	}
	return "Press 's' to open your SHELL.\nPress 'q' to quit.\n"
}

func main() {
	m := model{}
	if err := tea.NewProgram(m).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
