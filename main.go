package main

import (
	"cli-installer/cmd"
	"fmt"
	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)
//helloo
type model struct {
	choices []string      // List of technologies
	cursor  int           // Which choice our cursor is pointing at
	confirm bool          // Confirmation flag to proceed
	choice  string        // Final choice of technology
	quit    bool          // Quit flag to exit
	loading bool          // Loading flag
	spinner spinner.Model // Spinner model
}

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	return model{
		// Initial list of technologies
		choices: []string{"golang", "laravel", "java", "nodejs", "reactjs", "php"},
		spinner: s,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quit = true
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter":
			if m.confirm {
				m.choice = m.choices[m.cursor]
				m.loading = true
				return m, tea.Batch(m.spinner.Tick, waitForConfirmation())
			}
			m.confirm = true
		}
	case spinner.TickMsg:
		if m.loading {
			var cmd tea.Cmd
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd
		}
	case waitForMsg:
		m.loading = false
		m.quit = false
		return m, tea.Quit
	}
	return m, nil
}

func (m model) View() string {
	if m.quit {
		return ""
	}

	if m.loading {
		return fmt.Sprintf("%s Loading... Please wait.", m.spinner.View())
	}

	if m.confirm {
		return fmt.Sprintf("You selected: %s\nPress 'Enter' to confirm, 'q' to quit.\n", m.choices[m.cursor])
	}

	s := "What do you want to work with?\n\n"

	for i, choice := range m.choices {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress 'Enter' to select, 'q' to quit.\n"

	return s
}

type waitForMsg struct{}

func waitForConfirmation() tea.Cmd {
	return tea.Tick(3*time.Second, func(time.Time) tea.Msg {
		return waitForMsg{}
	})
}

func main() {
	p := tea.NewProgram(initialModel())
	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	m := finalModel.(model)
	fmt.Println(m.quit)
	if m.quit || m.choice == "" {
		fmt.Println("No technology selected. Exiting.")
		os.Exit(0)
	}

	action := m.choice

	actions := []string{"install", "update", "version"}
	var actionChoice string

	actionPrompt := &survey.Select{
		Message: "Select the action you want to perform:",
		Options: actions,
	}

	survey.AskOne(actionPrompt, &actionChoice)

	switch actionChoice {
	case "install":
		cmd.RunInstall(action)
	case "update":
		cmd.RunUpdate(action)
	case "version":
		cmd.RunVersion(action)
	default:
		fmt.Println("Invalid action selected")
	}
}
