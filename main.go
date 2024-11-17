package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

type model struct {
	count int // The counter value
}

// Init initializes the model. It's used to set up anything your app needs at the start.
func (m model) Init() tea.Cmd {
	return nil // No commands at start
}

// Update handles messages and updates the model's state.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg: // If a key is pressed
		switch msg.String() {
		case "q": // Quit the app
			return m, tea.Quit
		case "+": // Increment the counter
			m.count++
		case "-": // Decrement the counter
			m.count--
		}
	}
	return m, nil
}

var style = lipgloss.NewStyle().Foreground(lipgloss.Color("202")).Bold(true)

// View renders the UI. It returns the string that represents the UI
func (m model) View() string {
	return style.Render(fmt.Sprintf(
		"Counter: %d\n\n[+] Increment\n[-] Decrement\n[q] Quit",
		m.count,
	))
}

func main() {
	// Initialise the program
	p := tea.NewProgram(model{})

	// Run the program
	if err, _ := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting app: %s\n", err)
		os.Exit(1)
	}
}
