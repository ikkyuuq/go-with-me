package goblueprintbuilder

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Model struct {
	form    formModel
	spinner SpinnerModel
	state   int
}

const (
	StateForm = iota
	StateSpinner
	StateCompleted
)

func NewBuilder() *Model {
	fm := NewFormModel()
	sp := NewSpinnerModel()
	return &Model{
		form:    fm,
		spinner: sp,
	}
}

func (m Model) Init() tea.Cmd {
	return m.form.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	switch m.state {
	case StateForm:
		_, cmd := m.form.Update(msg)
		if m.form.State() == huh.StateCompleted {
			m.state = StateSpinner
			m.spinner = NewSpinnerModel()
			return m, m.spinner.Init()
		}
		return m, cmd
	case StateSpinner:
		spinnerModel, cmd := m.spinner.Update(msg)
		m.spinner = spinnerModel.(SpinnerModel)
		if m.spinner.done {
			m.state = StateCompleted
		}
		return m, cmd
	default:
		return m, nil
	}
}

func (m Model) View() string {
	if m.form.form == nil {
		return "Starting..."
	}
	switch m.state {
	case StateCompleted:
		return "Completed"
	case StateSpinner:
		return m.spinner.View()
	default:
		v := strings.TrimSuffix(m.form.View(), "\n\n")
		return v
	}
}
