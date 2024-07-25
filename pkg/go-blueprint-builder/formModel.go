package goblueprintbuilder

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type formModel struct {
	form *huh.Form
}

var (
	projectName string
	database    string
	framework   string
	confirm     bool
)

func NewFormModel() formModel {
	m := formModel{}

	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Name of your project").
				Prompt("? ").
				Value(&projectName).
				Key("projectName").
				Description("This will affect to your go mod"),
			huh.NewSelect[string]().
				Title("Choose your database").
				Value(&database).
				Key("database").
				Options(huh.NewOptions("None", "MySQL", "MongoDB", "SQLite", "Postgre")...).
				Description("This is will affect to your project structure"),
			huh.NewSelect[string]().
				Title("Choose your framework").
				Value(&framework).
				Key("framework").
				Options(huh.NewOptions("Standard Library", "Chi", "Echo", "Fiber", "Gin", "Gorilla", "HttpRouter")...).
				Description("This is will affect to your project structure"),
			huh.NewConfirm().
				Title("Would you like to confirm these setup").
				Value(&confirm).
				Affirmative("Let's go!").
				Negative("Wait, no").
				Validate(func(v bool) error {
					if !v {
						return fmt.Errorf("welp, finish up then")
					}
					return nil
				}),
		),
	).WithWidth(45).
		WithHeight(10).
		WithShowHelp(false).
		WithShowErrors(false)
	return m
}

func (m formModel) Init() tea.Cmd {
	return m.form.Init()
}

func (m formModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m formModel) State() huh.FormState {
	return m.form.State
}

func (m formModel) View() string {
	return m.form.View()
}
