package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	goblueprintbuilder "github.com/ikkyuuq/go-with-me/pkg/go-blueprint-builder"
)

type model struct {
	Builder goblueprintbuilder.Model
}

func initialModel() *goblueprintbuilder.Model {
	m := goblueprintbuilder.NewBuilder()
	return m
}

func (m model) Init() tea.Cmd {
	return m.Builder.Init()
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("there's been an error: %v", err)
		os.Exit(1)
	}
}
