package goblueprintbuilder

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type SpinnerModel struct {
	spinner     spinner.Model
	showSpinner bool
	done        bool
	result      string
}

func NewSpinnerModel() SpinnerModel {
	sp := spinner.New()
	sp.Spinner = spinner.Line

	m := SpinnerModel{
		spinner: sp,
	}
	return m
}

func (m SpinnerModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		m.prepareProject,
	)
}

func (m SpinnerModel) View() string {
	if m.done {
		return m.result
	}
	var s string
	s += fmt.Sprintf("\n %s %s \n\n", m.spinner.View(), "Prepareing your project...")
	return s
}

func (m SpinnerModel) prepareProject() tea.Msg {
	time.Sleep(5 * time.Second)
	return spinnerFinshedMsg{}
}

type spinnerFinshedMsg struct{}

func (m SpinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case spinnerFinshedMsg:
		m.done = true
		m.result = "Build Project Successful!\n\n"
		m.result += fmt.Sprintf("cd ./%s\n", projectName)
		return m, nil
	default:
		return m, nil
	}
}
